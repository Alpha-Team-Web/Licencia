package database

import (
	"back-src/model/existence"
)

func (db *Database) GetFieldSkills(fieldId string) (skills []string, error error) {
	var field existence.Field
	error = db.db.Model(&field).Column("skills").Where("id = ?", fieldId).Select()
	skills = field.Skills
	return
}

func (db Database) AddSkillToField(fieldId string, skill string) error {
	var skills []string
	field := existence.Field{Id: fieldId, Skills: skills}
	if err := db.db.Model(&field).Column("skills").Where("id = ?", field.Id).Select(); err != nil {
		return err
	}
	field.Skills = append(field.Skills, skill)
	if _, err := db.db.Model(&field).Column("skills").Where("id = ?", fieldId).Update(); err != nil {
		return err
	}
	return nil
}
