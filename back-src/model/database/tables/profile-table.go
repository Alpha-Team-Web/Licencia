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

func (table *ProfileTable) GetProfileImage(profileType string, id string) (existence.Profile, error) {
	profile := existence.Profile{}
	err := table.conn.Model(&profile).Where("id = ?", id).Where("type", profileType).Select()
	return profile, err
}

func NewProfileTable(db *pg.DB) *ProfileTable {
	return &ProfileTable{db}
}
