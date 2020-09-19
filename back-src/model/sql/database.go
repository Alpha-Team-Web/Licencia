package sql

import (
	"back-src/model/existence"
	"back-src/model/sql/tables"
	"encoding/json"
	"github.com/go-pg/pg"
	"io/ioutil"
)

const (
	jsonsFolderPath = "model/sql/jsons/"
)

type Initializable interface {
	Initialize() error
}

type Database struct {
	db                     *pg.DB
	AuthTokenTable         *tables.AuthTokenTable
	EmployerTable          *tables.EmployerTable
	FieldTable             *tables.FieldTable
	FreelancerTable        *tables.FreelancerTable
	ProjectTable           *tables.ProjectTable
	ReviewTable            *tables.ReviewTable
	ProfileTable           *tables.ProfileTable
	ProjectAttachmentTable *tables.ProjectAttachmentTable
	MediaTable             *tables.MediaTable
}

func NewDb() *Database {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     dbc.username,
		Password: dbc.password,
		Database: "Licencia-First",
	})

	return &Database{
		db:                     db,
		AuthTokenTable:         tables.NewAuthTokenTable(db),
		EmployerTable:          tables.NewEmployerTable(db),
		FieldTable:             tables.NewFieldsTable(db),
		FreelancerTable:        tables.NewFreelancerTable(db),
		ProjectTable:           tables.NewProjectTable(db),
		ReviewTable:            tables.NewReviewTable(db),
		ProfileTable:           tables.NewProfileTable(db),
		ProjectAttachmentTable: tables.NewProjectAttachment(db),
		MediaTable:             tables.NewMediaTable(db),
	}
}

func (db *Database) Initialize() error {
	if err := db.initEmployerTable(); err != nil {
		return err
	}
	if err := db.initFieldTable(); err != nil {
		return err
	}
	if err := db.initProjectTable(); err != nil {
		return nil
	}
	if err := db.initFreelancerTable(); err != nil {
		return err
	}
	if err := db.initSessionTable(); err != nil {
		panic(err)
	}
	if err := db.initReviewTables(); err != nil {
		return err
	}
	if err := db.initProfileTable(); err != nil {
		return err
	}
	if err := db.initProjectAttachmentTable(); err != nil {
		return err
	}
	if err := db.initFollowTable(); err != nil {
		return err
	}
	if err := db.initEventTable(); err != nil {
		return err
	}

	return nil
}

func (db *Database) initEmployerTable() error {
	err := db.db.CreateTable(&existence.Employer{}, options)
	return err
}

func (db *Database) initFieldTable() error {
	err := db.db.CreateTable(&existence.Field{}, options)
	if err != nil {
		return err
	}
	err = db.addDefaultFields()
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) addDefaultFields() error {
	defaultFieldsPath := jsonsFolderPath + "default-fields.json"
	bytes, err := ioutil.ReadFile(defaultFieldsPath)
	if err != nil {
		return err
	}
	var fields []existence.Field

	if err := json.Unmarshal(bytes, &fields); err != nil {
		return err
	}

	for _, field := range fields {
		if isThere, _ := db.FieldTable.IsThereFieldWithId(field.Id); !isThere {
			if _, err := db.db.Model(&field).Insert(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (db *Database) initProjectTable() error {
	return db.db.CreateTable(&existence.Project{}, options)
}

func (db *Database) initFreelancerTable() error {
	return db.db.CreateTable(&existence.Freelancer{}, options)
}

func (db *Database) initSessionTable() error {
	return db.db.CreateTable(&existence.AuthToken{}, options)
}

func (db *Database) initReviewTables() error {
	if err := db.db.CreateTable(&existence.FreelancerEmployerReview{}, options); err != nil {
		return err
	}
	if err := db.db.CreateTable(&existence.EmployerFreelancerReview{}, options); err != nil {
		return err
	}
	return nil
}

func (db *Database) initProfileTable() error {
	return db.db.CreateTable(&existence.Profile{}, options)
}

func (db *Database) initProjectAttachmentTable() error {
	return db.db.CreateTable(&existence.ProjectAttachment{}, options)
}

func (db *Database) initFollowTable() error {
	return db.db.CreateTable(&existence.Follow{}, options)
}

func (db *Database) initEventTable() error {
	return db.db.CreateTable(&existence.Event{}, options)
}
