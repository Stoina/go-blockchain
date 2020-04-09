package encoding

import "encoding/json"

// ConvertToJSON exported
// ...
func ConvertToJSON(object interface{}) string {
	converted, err := json.Marshal(object)

	if err != nil {
		return ""
	}

	return string(converted)
}
