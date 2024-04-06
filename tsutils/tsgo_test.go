package tsutils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetGoData(t *testing.T) {
	cases := []struct {
		name         string
		fileContents []byte
		expected     []string
		err          error
	}{
		// SUCCESSES
		{
			name: "a function with no parameters",
			fileContents: []byte(`
func saySomething() {
    fmt.Println("hola")
}
`),
			expected: []string{"func saySomething()"},
		},
		{
			name: "a function with parameters",
			fileContents: []byte(`
func saySomething(name string) {
    fmt.Printf("hola, %s", name)
}
`),
			expected: []string{"func saySomething(name string)"},
		},
		{
			name: "a function with parameters and a return type",
			fileContents: []byte(`
func saySomething(name string) string {
    return fmt.Sprintf("hola, %s", name)
}
`),
			expected: []string{"func saySomething(name string) string"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getGoData(tt.fileContents)
			fmt.Println(got)
			fmt.Println(tt.expected)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got: %#v, expected: %#v", got, tt.expected)
			}
			if err != tt.err {
				t.Errorf("error mismatch; got: %v, expected: %v", err, tt.err)
			}
		})
	}

}
