package usecases

import (
	"errors"
	"github.com/dbtedman/security-check/go/entities/checkresult"
	"github.com/dbtedman/security-check/go/entities/queryselector"
	"github.com/dbtedman/security-check/go/entities/url"
	"github.com/dbtedman/security-check/go/gateways/http_gateway"
	"regexp"
)

const InvalidHTTPGateway = "InvalidHTTPGateway"
const InvalidQuerySelector = "InvalidQuerySelector"
const InvalidRequestURL = "InvalidRequestURL"

type CheckReflectiveXSSRequest struct {
	HTTPGateway   http_gateway.HTTPGateway
	QuerySelector string
	RequestURL    string
}

type CheckReflectiveXSSResponse struct {
	CheckResults []checkresult.CheckResult
	Error        error
}

// CheckReflectiveXSS populates malicious data into a specific position in in
// the request URL in an attempt to have that data reflected back in the
// response which indicates potential vulnerability.
func CheckReflectiveXSS(request CheckReflectiveXSSRequest) CheckReflectiveXSSResponse {
	response := CheckReflectiveXSSResponse{}

	if request.HTTPGateway == nil {
		response.Error = errors.New(InvalidHTTPGateway)
		return response
	}

	if url.IsInvalid(request.RequestURL) {
		response.Error = errors.New(InvalidRequestURL)
		return response
	}

	if url.DoesNotContainValuePlaceholder(request.RequestURL) {
		response.Error = errors.New(InvalidRequestURL)
		return response
	}

	if queryselector.IsInvalid(request.QuerySelector) {
		response.Error = errors.New(InvalidQuerySelector)
		return response
	}

	for _, testURL := range GenerateTestURLs(request.RequestURL) {
		request.HTTPGateway.Get(testURL)
		// TODO: Add status to the CheckResult here.
		checkResult := checkresult.CheckResult{}
		response.CheckResults = append(response.CheckResults, checkResult)
	}

	return response
}

// Based on provided requestURL, replace {value} with each attack scenario.
func GenerateTestURLs(requestURL string) []string {
	var testURLs []string

	testURLs = append(testURLs, replaceValue(
		requestURL,
		"<script>window.alert('Hack!')</script>",
	))

	return testURLs
}

func replaceValue(source string, replace string) string {
	return regexp.MustCompile("{value}").ReplaceAllString(source, replace)
}
