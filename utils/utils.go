package utils

import "fmt"

func GetBasicAuth(username, password string) string {

	userPassPreEncode := fmt.Sprintf("%s:%s", username, password)

	return fmt.Sprintf("Basic %s", Base64Encode(userPassPreEncode))
}
