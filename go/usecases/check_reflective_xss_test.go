package usecases

import (
	"github.com/dbtedman/security-check/go/entities/url"
	"github.com/dbtedman/security-check/go/gateways/http_gateway"
	"gotest.tools/assert"
	"testing"
)

const ValidQuerySelector = ".query"
const ValidRequestURL = "https://danieltedman.com?q={value}"

func TestCanPerformCheck(t *testing.T) {
	httpGatewayFake := http_gateway.HTTPGatewayFake{}
	var httpGateway http_gateway.HTTPGateway = &httpGatewayFake
	response := CheckReflectiveXSS(
		CheckReflectiveXSSRequest{
			RequestURL:    ValidRequestURL,
			QuerySelector: ValidQuerySelector,
			HTTPGateway:   httpGateway,
		})
	assert.Equal(t, response.Error, nil)
}

func TestErrorsWithMissingRequestURL(t *testing.T) {
	httpGatewayFake := http_gateway.HTTPGatewayFake{}
	var httpGateway http_gateway.HTTPGateway = &httpGatewayFake
	response := CheckReflectiveXSS(
		CheckReflectiveXSSRequest{
			QuerySelector: ValidQuerySelector,
			HTTPGateway:   httpGateway,
		})
	assert.Error(t, response.Error, InvalidRequestURL)
}

func TestErrorsWithInvalidRequestURL(t *testing.T) {

	httpGatewayFake := http_gateway.HTTPGatewayFake{}
	var httpGateway http_gateway.HTTPGateway = &httpGatewayFake
	invalidRequestURLs := []string{"", "example.com", "https://example.com"}

	for _, invalidRequestURL := range invalidRequestURLs {
		t.Run("TestErrorsWithInvalidRequestURL["+invalidRequestURL+"]", func(t *testing.T) {
			response := CheckReflectiveXSS(
				CheckReflectiveXSSRequest{
					RequestURL:    invalidRequestURL,
					QuerySelector: ValidQuerySelector,
					HTTPGateway:   httpGateway,
				})
			assert.Error(t, response.Error, InvalidRequestURL)
		})
	}
}

func TestErrorsWithMissingQuerySelector(t *testing.T) {
	httpGatewayFake := http_gateway.HTTPGatewayFake{}
	var httpGateway http_gateway.HTTPGateway = &httpGatewayFake

	response := CheckReflectiveXSS(
		CheckReflectiveXSSRequest{
			RequestURL:  ValidRequestURL,
			HTTPGateway: httpGateway,
		})
	assert.Error(t, response.Error, InvalidQuerySelector)
}

func TestErrorsWithInvalidQuerySelector(t *testing.T) {

	invalidQuerySelectors := []string{""}

	httpGatewayFake := http_gateway.HTTPGatewayFake{}
	var httpGateway http_gateway.HTTPGateway = &httpGatewayFake

	for _, invalidQuerySelector := range invalidQuerySelectors {
		t.Run("TestErrorsWithInvalidQuerySelector["+invalidQuerySelector+"]", func(t *testing.T) {
			response := CheckReflectiveXSS(
				CheckReflectiveXSSRequest{
					RequestURL:    ValidRequestURL,
					QuerySelector: invalidQuerySelector,
					HTTPGateway:   httpGateway,
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

func TestErrorsWhenHTTPGatewayNotDefined(t *testing.T) {
	response := CheckReflectiveXSS(
		CheckReflectiveXSSRequest{
			RequestURL:    ValidRequestURL,
			QuerySelector: ValidQuerySelector,
		})
	assert.Error(t, response.Error, InvalidHTTPGateway)
}
