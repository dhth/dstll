package tsutils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetScalaData(t *testing.T) {
	cases := []struct {
		name         string
		fileContents []byte
		expected     []string
		err          error
	}{
		// SUCCESSES
		{
			name: "a simple class with 2 functions",
			fileContents: []byte(`
class MyClass {
  def method1(): Unit = {
    println("Function 1")
  }
  
  def method2(): Unit = {
    println("Function 2")
  }
}
`),
			expected: []string{"class MyClass\n\tdef method1(): Unit\n\tdef method2(): Unit"},
		},
		{
			name: "a simple class with class parameters",
			fileContents: []byte(`
class MyClass(val name: String, val age: Int) {
  def greet(): Unit = {
    println(s"Hello, my name is $name and I am $age years old.")
  }
}
`),
			expected: []string{"class MyClass(val name: String, val age: Int)\n\tdef greet(): Unit"},
		},
		{
			name: "a simple class with class parameters and an extends clause",
			fileContents: []byte(`
class MyExtendedClass(name: String, age: Int, val occupation: String) extends MyClass(name, age) {
  def introduce(): Unit = {
    println(s"I am a $occupation.")
  }
}
`),
			expected: []string{"class MyExtendedClass(name: String, age: Int, val occupation: String) extends MyClass(name, age)\n\tdef introduce(): Unit"},
		},
		{
			name: "a non scala class",
			fileContents: []byte(`
/#usr/bin/env bash

echo "hi"
`),
		},
		{
			name:         "an empty file",
			fileContents: []byte(""),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getScalaData(tt.fileContents)
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
