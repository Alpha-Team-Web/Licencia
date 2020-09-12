package fields

import (
	"back-src/model/database"
	"strings"
)

var Engine engine

func SearchSkillStartsWith(starter string) map[string]string {
	resultMap := map[string]string{}
	for skill, field := range Engine.SkillWithField {
		if strings.HasPrefix(strings.ToLower(skill), strings.ToLower(starter)) {
			resultMap[skill] = field.Name
		}
	}
	return resultMap
}

func GetSkillsByField(fieldId string) []string {
	for _, field := range Engine.Fields {
		if field.Id == fieldId {
			return field.Skills
		}
	}
	return []string{}
}

func AddSkillToField(fieldId string, skill string, db *database.Database) error {
	if _, ok := Engine.SkillWithField[skill]; ok {
		return nil
	} else {
		if field, err := db.FieldTable.GetField(fieldId); err != nil {
			return err
		} else {
			Engine.addNewSkill(field, skill)
			return db.FieldTable.AddSkillToField(fieldId, skill)
		}
	}
}
