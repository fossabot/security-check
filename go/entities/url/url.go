package url

import "regexp"

func IsValid(value string) bool {
	isValid, _ := regexp.MatchString("^https?://.+$", value)
	return isValid
}

func IsInvalid(value string) bool {
	return !IsValid(value)
}
