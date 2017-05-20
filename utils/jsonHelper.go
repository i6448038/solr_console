package utils

import "encoding/json"

func Decode(msg string)(map[string]interface{}, error){
	var m map[string]interface{}
	err := json.Unmarshal([]byte(msg), &m)
	return m,err
}
