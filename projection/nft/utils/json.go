package utils

import (
	"encoding/json"
	"strings"
)

func GetValueFromJSONData(rawJSON []byte, accessor string) (interface{}, error) {
	var data map[string]interface{}
	parseDataErr := json.Unmarshal(rawJSON, &data)
	if parseDataErr != nil {
		return nil, parseDataErr
	}
	accessors := strings.Split(accessor, ".")
	lastIndex := len(accessors) - 1
	for index, accessor := range accessors {
		var keyExist bool
		if index == lastIndex {
			var result interface{}
			result, keyExist = data[accessor].(interface{})
			if !keyExist {
				return nil, nil
			}
			return result, nil
		} else {
			data, keyExist = data[accessor].(map[string]interface{})
			if !keyExist {
				return nil, nil
			}
		}
	}

	return nil, nil
}
