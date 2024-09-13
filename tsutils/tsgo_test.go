package tsutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGoFuncs(t *testing.T) {
	cases := []struct {
		name            string
		fileContents    []byte
		expectedResults []string
		err             error
	}{
		// SUCCESSES
		{
			name: "a function with no parameters",
			fileContents: []byte(`
func saySomething() {
    fmt.Println("hola")
}
`),
			expectedResults: []string{"func saySomething()"},
		},
		{
			name: "a function with parameters",
			fileContents: []byte(`
func saySomething(name string) {
    fmt.Printf("hola, %s", name)
}
`),
			expectedResults: []string{"func saySomething(name string)"},
		},
		{
			name: "a function with parameters and a return type",
			fileContents: []byte(`
func saySomething(name string) string {
    return fmt.Sprintf("hola, %s", name)
}
`),
			expectedResults: []string{"func saySomething(name string) string"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			resultChan := make(chan Result)
			go getGoFuncs(resultChan, tt.fileContents)

			got := <-resultChan

			if tt.err == nil {
				assert.Equal(t, tt.expectedResults, got.Results)
				assert.NoError(t, got.Err)
			} else {
				assert.Equal(t, tt.err, got.Err)
			}
		})
	}
}

func TestGetGoTypes(t *testing.T) {
	cases := []struct {
		name            string
		fileContents    []byte
		expectedResults []string
		err             error
	}{
		// SUCCESSES
		{
			name: "a simple type",
			fileContents: []byte(`
type MyInt int
`),
			expectedResults: []string{"type MyInt int"},
		},
		{
			name: "a struct",
			fileContents: []byte(`
type Person struct {
    Name string
    Age  int
}
`),
			expectedResults: []string{`type Person struct {
    Name string
    Age  int
}`},
		},
		{
			name: "an interface",
			fileContents: []byte(`
type Shape interface {
    Area() float64
    Perimeter() float64
}
`),
			expectedResults: []string{`type Shape interface {
    Area() float64
    Perimeter() float64
}`},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			resultChan := make(chan Result)
			go getGoTypes(resultChan, tt.fileContents)

			got := <-resultChan

			if tt.err == nil {
				assert.Equal(t, tt.expectedResults, got.Results)
				assert.NoError(t, got.Err)
			} else {
				assert.Equal(t, tt.err, got.Err)
			}
		})
	}
}

func TestGetGoMethods(t *testing.T) {
	cases := []struct {
		name            string
		fileContents    []byte
		expectedResults []string
		err             error
	}{
		// SUCCESSES
		{
			name: "a method with no parameters",
			fileContents: []byte(`
func (p person) saySomething() {
    fmt.Printf("%s says, hola", p.name)
}
`),
			expectedResults: []string{"func (p person) saySomething()"},
		},
		{
			name: "a method with parameters",
			fileContents: []byte(`
func (p person) saySomething(say string) {
    fmt.Printf("%s says, %s", p.name, say)
}
`),
			expectedResults: []string{"func (p person) saySomething(say string)"},
		},
		{
			name: "a method with parameters and a return type",
			fileContents: []byte(`
func (p person) saySomething(say string) string {
    return fmt.Sprintf("%s says %s", p.name, say)
}
`),
			expectedResults: []string{"func (p person) saySomething(say string) string"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			resultChan := make(chan Result)
			go getGoMethods(resultChan, tt.fileContents)

			got := <-resultChan

			if tt.err == nil {
				assert.Equal(t, tt.expectedResults, got.Results)
				assert.NoError(t, got.Err)
			} else {
				assert.Equal(t, tt.err, got.Err)
			}
		})
	}
}
