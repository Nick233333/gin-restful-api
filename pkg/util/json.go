package util

import (
	"encoding/json"
)

func InterfaceToJson(data interface{}) string {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return string(jsonStr)
}
