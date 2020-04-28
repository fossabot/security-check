package activities

import (
	"errors"
	"github.com/dbtedman/security-check/go/entities/queryselector"
	"github.com/dbtedman/security-check/go/entities/url"
)

const InvalidRequestURL = "InvalidRequestURL"
const InvalidQuerySelector = "InvalidQuerySelector"

type CheckReflectiveXSSRequest struct {
	RequestURL    string
	QuerySelector string
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

	if queryselector.IsInvalid(request.QuerySelector) {
		response.Error = errors.New(InvalidQuerySelector)
		return response
	}

	return response
}
