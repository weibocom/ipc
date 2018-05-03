package util

import "encoding/json"

func ToJSON(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func FromJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
