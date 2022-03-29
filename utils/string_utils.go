package utils

func IsEmptyString(s *string) bool {
	if s == nil || *s == "" {
		return true
	}
	return false
}
