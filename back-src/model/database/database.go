package database

import (
	"back-src/model/database/tables"
	"back-src/model/existence"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/go-pg/pg"
)

const (
	jsonsFolderPath = "model/database/jsons/"
)

type Metadata struct {
	IsFirstInit bool `json:"is_first_init"`
}

type Initializable interface {
	Initialize() error
}

type Database struct {
	db                     *pg.DB
	meta                   *Metadata
	AuthTokenTable         *tables.AuthTokenTable
	EmployerTable          *tables.EmployerTable
	FieldTable             *tables.FieldTable
	FreelancerTable        *tables.FreelancerTable
	ProjectTable           *tables.ProjectTable
	ReviewTable            *tables.ReviewTable
	ProfileTable           *tables.ProfileTable
	ProjectAttachmentTable *tables.ProjectAttachmentTable
	MediaTable      *tables.MediaTable
}

func NewDb() *Database {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     dbc.username,
		Password: dbc.password,
		Database: "Licencia-First",
	})

	bytes, err := ioutil.ReadFile("model/database/jsons/db-metadata.json")
	if err != nil {
		//panic(err)
	}
	meta := &Metadata{}
	err = json.Unmarshal(bytes, &meta)

	return &Database{
		db:              db,
		meta:            meta,
		AuthTokenTable:  tables.NewAuthTokenTable(db),
		EmployerTable:   tables.NewEmployerTable(db),
		FieldTable:      tables.NewFieldsTable(db),
		FreelancerTable: tables.NewFreelancerTable(db),
		ProjectTable:    tables.NewProjectTable(db),
		ReviewTable:     tables.NewReviewTable(db),
		ProfileTable:    tables.NewProfileTable(db),
		ProjectAttachmentTable: tables.NewProjectAttachment(db),
		MediaTable:      tables.NewMediaTable(db),
	}
}

func (db *Database) Initialize() error {
	defer func() {
		db.meta.IsFirstInit = false
		if err := db.updateDBMetadata(); err != nil {
			//panic(err)
		}
	}()
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
	if db.meta.IsFirstInit {
		err = db.addDefaultFields()
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *Database) addDefaultFields() error {
	deafaultFieldsPath := jsonsFolderPath + "default-fields.json"
	bytes, err := ioutil.ReadFile(deafaultFieldsPath)
	if err != nil {
		return err
	}
	fields := []existence.Field{}
	err = json.Unmarshal(bytes, &fields)

	if err != nil {
		return err
	}

	_, err = db.db.Model(&fields).Insert()

	if err != nil {
		return err
	}

	return nil
}

func (db *Database) updateDBMetadata() error {
	metadataPath := jsonsFolderPath + "db-metadata.json"

	if err := os.Remove(metadataPath); err != nil {
		return nil
	}
	if bytes, err := json.Marshal(db.meta); err == nil {
		if file, err := os.Create(metadataPath); err == nil {
			if _, err := file.Write(bytes); err != nil {
				return err
			}
			if err := file.Close(); err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		return err
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
