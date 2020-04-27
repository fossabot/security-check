package check

import (
	"gotest.tools/assert"
	"testing"
)

func TestCheckConstruct(t *testing.T) {
	name := "Example"
	check := New(name)

	assert.Equal(t, check.Name(), name)
}
