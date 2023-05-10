package common

import (
	"encoding/json"
)

func JSONStringify(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func JSONParse(s string, v interface{}) error {
	err := json.Unmarshal([]byte(s), v)
	return err
}

func ContainsString(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
