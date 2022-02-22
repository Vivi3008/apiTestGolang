package http

import (
	"testing"

	"gotest.tools/v3/assert"
)

func AssertEqual(t *testing.T, want interface{}, got interface{}) {
	assert.Equal(t, want, got)
}
