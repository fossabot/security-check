package usecases

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/dbtedman/security-check/go/entities/checkresult"
	"github.com/dbtedman/security-check/go/entities/queryselector"
	"github.com/dbtedman/security-check/go/entities/url"
	"github.com/dbtedman/security-check/go/gateways/http_gateway"
	"regexp"
	"strings"
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

type GeneratedTestURL struct {
	RequestURL    string
	ReplacedValue string
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

	for _, generatedTestURL := range GenerateTestURLs(request.RequestURL) {
		getResponse := request.HTTPGateway.Get(generatedTestURL.RequestURL)

		isReflected := LocateReflectedContent(
			request.QuerySelector,
			generatedTestURL.ReplacedValue,
			getResponse.ResponseBody,
		)

		checkResult := checkresult.CheckResult{
			Successful: isReflected,
		}

		response.CheckResults = append(response.CheckResults, checkResult)
	}

	return response
}

// Try to located reflectedContent at position identified by querySelector in
// responseBody returned by a http request.
func LocateReflectedContent(querySelector string, reflectedContent string, responseBody string) bool {
	doc, queryError := goquery.NewDocumentFromReader(strings.NewReader(responseBody))

	if queryError != nil {
		return false
	}

	matchFound := false

	doc.Find(querySelector).Each(func(_ int, selection *goquery.Selection) {
		ret, selectionError := selection.Html()

		if selectionError == nil && strings.Trim(ret, " \n") == reflectedContent {
			matchFound = true
		}
	})

	return matchFound
}

// Based on provided requestURL, replace {value} with each attack scenario.
func GenerateTestURLs(requestURL string) []GeneratedTestURL {
	var testURLs []GeneratedTestURL

	for _, testValue := range []string{
		`<script>alert("Hack!")</script>`,
		`<script src="https://example.com/hack.js"></script>`,
		`<iframe src="https://example.com/hack"></iframe>`,
		`<div onload="alert('Hack!')"></div>`,
	} {
		testURLs = append(testURLs, GeneratedTestURL{
			RequestURL: replaceValue(
				requestURL,
				testValue,
			),
			ReplacedValue: testValue,
		})
	}

	return testURLs
}

func replaceValue(source string, replace string) string {
	return regexp.MustCompile("{value}").ReplaceAllString(source, replace)
}

// References:
//    https://developer.mozilla.org/en-US/docs/Web/HTML/Element/embed
//    https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe
//    https://developer.mozilla.org/en-US/docs/Web/HTML/Element/object
//    https://www.softwaretestinghelp.com/cross-site-scripting-xss-attack-test/
