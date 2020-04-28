package queryselector

import "regexp"

func IsValid(value string) bool {
	isValid, _ := regexp.MatchString("^.+$", value)
	return isValid
}

func IsInvalid(value string) bool {
	return !IsValid(value)
}
