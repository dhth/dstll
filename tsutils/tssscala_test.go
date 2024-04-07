package tsutils

import (
	"reflect"
	"testing"
)

func TestGetScalaFunctions(t *testing.T) {
	cases := []struct {
		name         string
		fileContents []byte
		expected     []string
		err          error
	}{
		// SUCCESSES
		{
			name: "a simple function",
			fileContents: []byte(`
def method1(): Unit = {
    println("Function 1")
  }
`),
			expected: []string{"def method1(): Unit"},
		},
		{
			name: "two functions",
			fileContents: []byte(`
def method1(): Unit = {
    println("Function 1")
  }
def method2(): Unit = {
    println("Function 2")
  }
`),
			expected: []string{
				"def method1(): Unit",
				"def method2(): Unit",
			},
		},
		{
			name: "a function with arguments",
			fileContents: []byte(`
def method1(arg1: String, num: Int): Unit = {
  for (i <- 1 to num) {
    println(arg1)
  }
}
`),
			expected: []string{
				"def method1(arg1: String, num: Int): Unit",
			},
		},
		{
			name: "a function with override modifier",
			fileContents: []byte(`
override def method1(arg1: String, num: Int): Unit = {
  for (i <- 1 to num) {
    println(arg1)
  }
}
`),
			expected: []string{
				"override def method1(arg1: String, num: Int): Unit",
			},
		},
		{
			name: "a function with private modifier",
			fileContents: []byte(`
private def method1(arg1: String, num: Int): Unit = {
  for (i <- 1 to num) {
    println(arg1)
  }
}
`),
			expected: []string{
				"private def method1(arg1: String, num: Int): Unit",
			},
		},
		{
			name: "a function with two modifiers",
			fileContents: []byte(`
override protected def method1(arg1: String, num: Int): Unit = {
  for (i <- 1 to num) {
    println(arg1)
  }
}
`),
			expected: []string{
				"override protected def method1(arg1: String, num: Int): Unit",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			resultChan := make(chan Result)
			go getScalaFunctions(resultChan, tt.fileContents)

			got := <-resultChan

			if !reflect.DeepEqual(got.Results, tt.expected) {
				t.Errorf("got: %#v, expected: %#v", got.Results, tt.expected)
			}
			if got.Err != tt.err {
				t.Errorf("error mismatch; got: %v, expected: %v", got.Err, tt.err)
			}
		})
	}

}

func TestGetScalaClasses(t *testing.T) {
	cases := []struct {
		name            string
		fileContents    []byte
		expectedResults []string
		err             error
	}{
		// SUCCESSES
		{
			name: "a simple class with 2 functions",
			fileContents: []byte(`
class MyClass {
  def method1(): Unit = {
    println("Function 1")
  }
}
`),
			expectedResults: []string{"class MyClass"},
		},
		{
			name: "two classes",
			fileContents: []byte(`
class MyClass {
  def method1(): Unit = {
    println("Function 1")
  }
}
class MyClass2 {
  def method2(): Unit = {
    println("Function 2")
  }
}
`),
			expectedResults: []string{"class MyClass", "class MyClass2"},
		},
		{
			name: "a class with modifiers",
			fileContents: []byte(`
sealed abstract class KafkaConsumer {}
`),
			expectedResults: []string{"sealed abstract class KafkaConsumer"},
		},
		{
			name: "a class with modifiers and type params",
			fileContents: []byte(`
sealed abstract class Signature[+T] { self =>

  final def show: String = mergeShow(new StringBuilder(30)).toString

}
`),
			expectedResults: []string{"sealed abstract class Signature[+T]"},
		},
		{
			name: "class within a class",
			fileContents: []byte(`
class OuterClass {
  // Members and methods of the outer class
  
  class InnerClass {
    // Members and methods of the inner class
  }
}
`),
			// TODO: Find a way to describe this nested structure, as opposed to returning
			// classes in a flattened structure
			expectedResults: []string{"class OuterClass", "class InnerClass"},
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
			expectedResults: []string{"class MyClass(val name: String, val age: Int)"},
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
			expectedResults: []string{"class MyExtendedClass(name: String, age: Int, val occupation: String) extends MyClass(name, age)"},
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
			resultChan := make(chan Result)
			go getScalaClasses(resultChan, tt.fileContents)

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

func TestGetScalaObjects(t *testing.T) {
	cases := []struct {
		name         string
		fileContents []byte
		expected     Result
	}{
		// SUCCESSES
		{
			name: "a simple object",
			fileContents: []byte(`
object HelloWorld {
  def main(args: Array[String]): Unit = {
    println("Hello, world!")
  }
}
object HelloWorld2 {
  def main2(args: Array[String]): Unit = {
    println("Hello, world!")
  }
}
`),
			expected: Result{
				Results: []string{"object HelloWorld", "object HelloWorld2"},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			resultChan := make(chan Result)
			go getScalaObjects(resultChan, tt.fileContents)

			got := <-resultChan

			if !reflect.DeepEqual(got.Results, tt.expected.Results) {
				t.Errorf("got: %#v, expected: %#v", got, tt.expected)
			}
			if got.Err != tt.expected.Err {
				t.Errorf("error mismatch; got: %v, expected: %v", got.Err, tt.expected.Err)
			}
		})
	}

}
