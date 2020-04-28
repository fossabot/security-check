package activities

import (
	"gotest.tools/assert"
	"testing"
)

const ValidQuerySelector = ".query"
const ValidRequestURL = "https://example.com"

func TestCanPerformCheck(t *testing.T) {
	response := CheckReflectiveXSS(
		CheckReflectiveXSSRequest{
			RequestURL:    ValidRequestURL,
			QuerySelector: ValidQuerySelector,
		})
	assert.Equal(t, response.Error, nil)
}

func TestErrorsWithMissingRequestURL(t *testing.T) {
	response := CheckReflectiveXSS(
		CheckReflectiveXSSRequest{
			QuerySelector: ValidQuerySelector,
		})
	assert.Error(t, response.Error, InvalidRequestURL)
}

func TestErrorsWithInvalidRequestURL(t *testing.T) {

	invalidRequestURLs := []string{"", "example.com"}

	for _, invalidRequestURL := range invalidRequestURLs {
		t.Run("TestErrorsWithInvalidRequestURL["+invalidRequestURL+"]", func(t *testing.T) {
			response := CheckReflectiveXSS(
				CheckReflectiveXSSRequest{
					RequestURL:    invalidRequestURL,
					QuerySelector: ValidQuerySelector,
				})
			assert.Error(t, response.Error, InvalidRequestURL)
		})
	}
}

func TestErrorsWithMissingQuerySelector(t *testing.T) {
	response := CheckReflectiveXSS(
		CheckReflectiveXSSRequest{
			RequestURL: ValidRequestURL,
		})
	assert.Error(t, response.Error, InvalidQuerySelector)
}

func TestErrorsWithInvalidQuerySelector(t *testing.T) {

	invalidQuerySelectors := []string{""}

	for _, invalidQuerySelector := range invalidQuerySelectors {
		t.Run("TestErrorsWithInvalidQuerySelector["+invalidQuerySelector+"]", func(t *testing.T) {
			response := CheckReflectiveXSS(
				CheckReflectiveXSSRequest{
					RequestURL:    ValidRequestURL,
					QuerySelector: invalidQuerySelector,
				})
			assert.Error(t, response.Error, InvalidQuerySelector)
		})
	}
}
