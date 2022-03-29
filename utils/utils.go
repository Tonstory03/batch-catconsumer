package utils

import (
	"encoding/json"
	"fmt"
)

func GetBasicAuth(username, password string) string {

	userPassPreEncode := fmt.Sprintf("%s:%s", username, password)

	return fmt.Sprintf("Basic %s", Base64Encode(userPassPreEncode))
}

func ToJsonText(d *interface{}) *string {
	var result *string

	if d != nil {
		b, _ := json.Marshal(*d)
		text := string(b)
		result = &text
	}

	return result
}
