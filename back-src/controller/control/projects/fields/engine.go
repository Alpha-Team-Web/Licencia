package fields

import (
	"back-src/model/existence"
	"back-src/model/sql"
)

type engine struct {
	SkillWithField map[string]existence.Field
	Fields         []existence.Field
}

func (eng *engine) addNewSkill(field existence.Field, skill string) {
	eng.SkillWithField[skill] = field
	for _, e := range eng.Fields {
		if e.Id == field.Id {
			e.Skills = append(e.Skills, skill)
		}
	}
}

func NewEngine(db *sql.Database) engine {
	engine := engine{
		SkillWithField: map[string]existence.Field{},
		Fields:         []existence.Field{},
	}

	if fields, err := db.FieldTable.GetAllFieldsWithSkills(); err != nil {
		panic(err)
	} else {
		engine.Fields = fields
		for _, field := range fields {
			for _, skill := range field.Skills {
				newField := existence.Field{
					Id:   field.Id,
					Name: field.Name,
				}
				engine.SkillWithField[skill] = newField
			}
		}
	}
	return engine
}
