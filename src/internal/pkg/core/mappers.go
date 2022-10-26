package core

import "encoding/json"

func MapToJson(toMap interface{}) string {
	result, err := json.Marshal(toMap)
	if err != nil {
		return "{}"
	}

	return string(result)
}

func MapFromJson(toMap string, target interface{}) error {
	err := json.Unmarshal([]byte(toMap), target)

	return err
}
