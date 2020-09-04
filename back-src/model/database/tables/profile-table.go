package tables

import (
	"back-src/model/existence"
	"github.com/go-pg/pg"
)

type ProfileTable struct {
	conn *pg.DB
}

func (table *ProfileTable) AddProfileImage(profile existence.Profile) error {
	_, err := table.conn.Model(&profile).Insert()
	return err
}

func NewProfileTable(db *pg.DB) *ProfileTable {
	return &ProfileTable{db}
}
