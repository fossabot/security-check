package url

import "regexp"

func IsValid(value string) bool {
	isValid, _ := regexp.MatchString("^https?://.+$", value)
	return isValid
}

func IsInvalid(value string) bool {
	return !IsValid(value)
}

func ContainsValuePlaceholder(value string) bool {
	containsValuePlaceholder, _ := regexp.MatchString("{value}", value)
	return containsValuePlaceholder
}

func DoesNotContainValuePlaceholder(value string) bool {
	return !ContainsValuePlaceholder(value)
}
