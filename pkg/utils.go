package pkg

import (
	"encoding/base64"
	"encoding/json"

	"github.com/ochom/gutils/logs"
)

func GetFilters(filters string) []map[string]any {
	decodedFilters, err := base64.StdEncoding.DecodeString(filters)
	if err != nil {
		logs.Error("error decoding filters: %s", err)
		return []map[string]any{}
	}

	logs.Debug("decoded filters: %s", string(decodedFilters))

	var mapped []map[string]any
	if err := json.Unmarshal(decodedFilters, &mapped); err != nil {
		logs.Error("error unmarshalling filters: %s", err)
		return []map[string]any{}
	}

	return mapped
}

func GetFilterColumns(filters []map[string]any) []string {
	filterColumns := []string{}
	for _, filter := range filters {
		filterColumns = append(filterColumns, filter["column"].(string))
	}

	return filterColumns
}
