package tsutils

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testcode/go.txt
var goCode []byte

func TestGetGoFuncs(t *testing.T) {
	expected := []string{
		"func saySomething()",
		"func saySomething(name string)",
		"func saySomething(name string) string",
		"func getGenericResult(fContent []byte, query string, language *ts.Language) ([]string, error)",
		"func Clone[S ~[]E, E any](s S) S",
	}
	resultChan := make(chan Result)
	go getGoFuncs(resultChan, goCode)

	got := <-resultChan

	require.NoError(t, got.Err)
	assert.Equal(t, expected, got.Results)
}

func TestGetGoTypes(t *testing.T) {
	expected := []string{
		"type MyInt int",
		`type Person struct {
    Name string
    Age  int
}`,
		`type Shape interface {
    Area() float64
    Perimeter() float64
}`,
	}

	resultChan := make(chan Result)
	go getGoTypes(resultChan, goCode)

	got := <-resultChan

	require.NoError(t, got.Err)
	assert.Equal(t, expected, got.Results)
}

func TestGetGoMethods(t *testing.T) {
	expected := []string{
		"func (p person) saySomething()",
		"func (p person) saySomething(say string)",
		"func (p person) saySomething(say string) string",
		"func (p person) saySomething(say string) (string, error)",
		"func (s *slice[E, V]) Map(doSomething func(E) V) ([]E, error)",
	}

	resultChan := make(chan Result)
	go getGoMethods(resultChan, goCode)

	got := <-resultChan

	require.NoError(t, got.Err)
	assert.Equal(t, expected, got.Results)
}
