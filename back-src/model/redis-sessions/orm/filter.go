package orm

import (
	"back-src/controller/utils/libs"
	"back-src/view/data"
	"encoding/json"
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
	filter.Status = hashed["status"]
	filter.MinPrice, _ = strconv.ParseFloat(hashed["min-price"], 64)
	filter.MaxPrice, _ = strconv.ParseFloat(hashed["max-price"], 64)
	json.Unmarshal([]byte(hashed["must-include-skills"]), &filter.MustIncludeSkills)
	json.Unmarshal([]byte(hashed["include-skills"]), &filter.IncludeSkills)
	json.Unmarshal([]byte(hashed["exclude-skills"]), &filter.ExcludeSkills)
	var projectIds []string
	json.Unmarshal([]byte(hashed["projectIds"]), &projectIds)
	filter.IsFilterBySkill = hashed["is-filter-by-skill"] == "1"
	return filter, projectIds
}
