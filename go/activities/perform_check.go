package activities

import "github.com/dbtedman/security-check/go/entities/check"

type PerformCheckRequest struct {
	Check check.Check
}

type PerformCheckResponse struct {
	Success bool
}

func PerformCheck(request PerformCheckRequest) PerformCheckResponse {
	return PerformCheckResponse{Success: true}
}
