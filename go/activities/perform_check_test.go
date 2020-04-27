package activities

import (
	"github.com/dbtedman/security-check/go/entities/check"
	"gotest.tools/assert"
	"testing"
)

func TestCanPerformCheck(t *testing.T) {
	result := PerformCheck(PerformCheckRequest{Check: check.New("Demo")})

	assert.Equal(t, result.Success, true)
}
