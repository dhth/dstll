package tsutils

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testcode/python.txt
var pyCode []byte

func TestGetPyData(t *testing.T) {
	expected := []string{
		"def say_something()",
		"def say_something(x)",
		"def say_something(x: str, y: str)",
		"def say_something() -> str",
		"def say_something(x: str, y: str) -> str",
	}

	got, err := getPyData(pyCode)

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}
