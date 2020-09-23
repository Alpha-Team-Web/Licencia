package orm

import (
	"back-src/controller/utils/libs"
	"back-src/view/data"
	"strconv"
)

func HashFilter(filter data.Filter, projectIds []string) map[string]interface{} {
	return map[string]interface{}{
		"status":              filter.Status,
		"min-price":           filter.MinPrice,
		"max-price":           filter.MaxPrice,
		"must-include-skills": libs.Marshal(filter.MustIncludeSkills),
		"include-skills":      libs.Marshal(filter.IncludeSkills),
		"exclude-skills":      libs.Marshal(filter.ExcludeSkills),
		"is-filter-by-skill":  libs.Ternary(filter.IsFilterBySkill, "1", "0").(string),
		"projectIds":          libs.Marshal(projectIds),
	}
}

func UnHashFilter(hashed map[string]string) (data.Filter, []string) {
	filter := data.Filter{}
	strconv.ParseFloat()

}
