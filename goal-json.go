/*
 * Go simple JSON library module.
 * Created: 21/06/2022
 * Write by:
 * - Hari Yulianto Pratama
 */

package json

import (
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
	result, error := json.Marshal(content)
	if error != nil {
		return "", fmt.Errorf("JSON Encode failed => %q", error)
	}
	return string(result), nil
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

