package activities

import (
	"errors"
	"github.com/dbtedman/security-check/go/entities/queryselector"
	"github.com/dbtedman/security-check/go/entities/url"
	"github.com/dbtedman/security-check/go/gateways/http_gateway"
	"regexp"
)

const InvalidRequestURL = "InvalidRequestURL"
const InvalidQuerySelector = "InvalidQuerySelector"

type CheckReflectiveXSSRequest struct {
	RequestURL    string
	QuerySelector string
	// TODO: Use this library to interact with HTTP.
	HTTPGateway http_gateway.HTTPGateway
}

type CheckReflectiveXSSResponse struct {
	Error error
}

// CheckReflectiveXSS populates malicious data into a specific position in in
// the request URL in an attempt to have that data reflected back in the
// response which indicates potential vulnerability.
func CheckReflectiveXSS(request CheckReflectiveXSSRequest) CheckReflectiveXSSResponse {
	response := CheckReflectiveXSSResponse{}

	if url.IsInvalid(request.RequestURL) {
		response.Error = errors.New(InvalidRequestURL)
		return response
	}

	// TODO: We need to test that the URL contains a {value} placeholder.

	if queryselector.IsInvalid(request.QuerySelector) {
		response.Error = errors.New(InvalidQuerySelector)
		return response
	}

	return response
}

func GenerateTestURLs(requestURL string) []string {
	var testURLs []string

	testURLs = appendReplacedValue(
		testURLs,
		requestURL,
		"<script>window.alert('Hack!')</script>",
	)

	return testURLs
}

func appendReplacedValue(urls []string, source string, replace string) []string {
	return append(urls, replaceValue(
		source,
		replace,
	))
}

func replaceValue(source string, replace string) string {
	return regexp.MustCompile("{value}").ReplaceAllString(source, replace)
}
