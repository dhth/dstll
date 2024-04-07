package tsutils

import (
	"reflect"
	"testing"
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

			if !reflect.DeepEqual(got.Results, tt.expectedResults) {
				t.Errorf("got: %#v, expected: %#v", got.Results, tt.expectedResults)
			}
			if got.Err != tt.err {
				t.Errorf("error mismatch; got: %v, expected: %v", got.Err, tt.err)
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

			if !reflect.DeepEqual(got.Results, tt.expectedResults) {
				t.Errorf("got: %#v, expected: %#v", got.Results, tt.expectedResults)
			}
			if got.Err != tt.err {
				t.Errorf("error mismatch; got: %v, expected: %v", got.Err, tt.err)
			}
		})
	}

}
