package helper

import (
	"encoding/json"
	"fmt"
)

func StringDefault(val *string, def string) string {
	if val == nil {
		return def
	}
	return *val
}

func StringToMap(body string) map[string]interface{} {
	var data map[string]interface{}

	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}

	return data
}
