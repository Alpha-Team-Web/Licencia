package tables

import (
	"back-src/model/existence"
	"github.com/go-pg/pg"
)

type FieldTable struct {
	conn *pg.DB
}

func (table *FieldTable) GetFieldSkills(fieldId string) (skills []string, error error) {
	var field existence.Field
	error = table.conn.Model(&field).Column("skills").Where("id = ?", fieldId).Select()
	skills = field.Skills
	return
}

func (table *FieldTable) AddSkillToField(fieldId string, skill string) error {
	var skills []string
	field := existence.Field{Id: fieldId, Skills: skills}
	if err := table.conn.Model(&field).Column("skills").Where("id = ?", field.Id).Select(); err != nil {
		return err
	}
	field.Skills = append(field.Skills, skill)
	if _, err := table.conn.Model(&field).Column("skills").Where("id = ?", fieldId).Update(); err != nil {
		return err
	}
	return nil
}
