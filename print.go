package hatty

import "encoding/json"

// Transform to JSON, output as string
func ToJSON(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), err
}

// Transform to JSON, output as string, if error occurred,
func MustToJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(data)
}
