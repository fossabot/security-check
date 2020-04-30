package activities

import (
	"github.com/dbtedman/security-check/go/entities/url"
	"gotest.tools/assert"
	"testing"
)

const ValidQuerySelector = ".query"
const ValidRequestURL = "https://example.com?q={value}"

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

func TestGenerateTestURLs(t *testing.T) {
	generatedURLs := GenerateTestURLs(ValidRequestURL)

	assert.Equal(t, len(generatedURLs) > 0, true)

	for _, generatedURL := range generatedURLs {
		assert.Equal(t, url.IsValid(generatedURL), true)
		assert.Equal(t, generatedURL != ValidRequestURL, true)
	}
}
