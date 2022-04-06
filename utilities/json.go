package utilities

import "encoding/json"

func JsonEncode(object interface{}) string {
	jsonByte, _ := json.MarshalIndent(object, "", "  ")
	return string(jsonByte)
}