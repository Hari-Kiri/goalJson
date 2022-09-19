package goalJson

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Encode data to json
func JsonEncode(content map[string]interface{}, indent bool) (string, error) {
	if indent {
		result, error := json.MarshalIndent(content, "", "\t")
		if error != nil {
			return "", fmt.Errorf("JSON Encode failed => %q", error)
		}
		return string(result), nil
	}
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	errorEncodingNoIndent := encoder.Encode(content)
	if errorEncodingNoIndent != nil {
		return "", fmt.Errorf("JSON Encode failed => %q", errorEncodingNoIndent)
	}
	return buffer.String(), nil
}

// Decode json data
func JsonDecode(content string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	error := json.Unmarshal([]byte(content), &result)
	if error != nil {
		return nil, fmt.Errorf("JSON Decode failed => %q", error)
	}
	return result, nil
}
