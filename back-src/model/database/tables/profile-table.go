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

func (table *ProfileTable) UpdateProfileImage(profile existence.Profile) error {
	_, err := table.conn.Model(&profile).Where("id = ?", profile.Id).Where("type = ?", profile.Type).Update()
	return err
}

func (table *ProfileTable) GetProfileImage(profileType string, id string) (existence.Profile, error) {
	profile := existence.Profile{}
	err := table.conn.Model(&profile).Where("id = ?", id).Where("type = ?", profileType).Select()
	return profile, err
}

func (table *ProfileTable) HasProfile(id, profileType string) (bool, error) {
	profiles := []existence.Profile{}
	err := table.conn.Model(&profiles).Where("id = ?", id).Where("type = ?", profileType).Select()
	if err != nil {
		return false, err
	}
	return len(profiles) != 0, nil
}

func NewProfileTable(db *pg.DB) *ProfileTable {
	return &ProfileTable{db}
}
