package fields

import (
	"back-src/model/database"
	"strings"
)

var Engine engine

func SearchSkillStartsWith(starter string) map[string]string {
	resultMap := map[string]string{}
	for skill, field := range Engine.skillWithField {
		if strings.HasPrefix(strings.ToLower(skill), strings.ToLower(starter)) {
			resultMap[skill] = field.Name
		}
	}
	return resultMap
}

func GetSkillsByField(fieldId string) []string {
	for _, field := range Engine.fields {
		if field.Id == fieldId {
			return field.Skills
		}
	}
	return []string{}
}

func AddSkillToField(fieldId string, skill string, db *database.Database) error {

}
