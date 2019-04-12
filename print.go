package hatty

import (
	"encoding/json"
	"fmt"
)

// Transform to JSON, return string
func ToJSON(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), err
}

// Transform to JSON, return string. panic it if error occurred
func MustToJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// Print JSON to Stdout
func PrintJSON(v interface{}) error {
	data, err := json.Marshal(v)
	fmt.Println(string(data))
	return err
}

// Print JSON to Stdout, panic it if error occurred
func MustPrintJSON(v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
