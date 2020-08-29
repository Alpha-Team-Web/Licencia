package database

import (
	"back-src/model/existence"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

const (
	jsonsFolderPath = "model/database/jsons/"
)

var (
	dbc = dbConnection{
		username: "ashka",
		password: "a124578",
		/*		username: "postgres",
				password: "mbsoli1743399413",*/
	}

	options = &orm.CreateTableOptions{
		IfNotExists: true,
	}
)

type Metadata struct {
	IsFirstInit bool `json:"is_first_init"`
}

type Initializable interface {
	Initialize() error
}

type dbConnection struct {
	username string
	password string
}

type Database struct {
	db   *pg.DB
	meta *Metadata
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
		panic(err)
	}
	meta := &Metadata{}
	err = json.Unmarshal(bytes, &meta)

	return &Database{db, meta}
}

func (db *Database) Initialize() error {

	defer func() {
		db.meta.IsFirstInit = false
		if err := db.updateDBMetadata(); err != nil {
			panic(err)
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
