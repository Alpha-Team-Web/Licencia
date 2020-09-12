package fields

import (
	"back-src/model/database"
	"back-src/model/existence"
)

type engine struct {
	skillWithField map[string]existence.Field
	fields         []existence.Field
}

func NewEngine(db *database.Database) engine {
	engine := engine{
		skillWithField: map[string]existence.Field{},
		fields:         []existence.Field{},
	}

	if fields, err := db.FieldTable.GetAllFieldsWithSkills(); err != nil {
		panic(err)
	} else {
		engine.fields = fields
		for _, field := range fields {
			for _, skill := range field.Skills {
				newField := existence.Field{
					Id:   field.Id,
					Name: field.Name,
				}
				engine.skillWithField[skill] = newField
			}
		}
	}
	return engine
}
