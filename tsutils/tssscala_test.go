package tsutils

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed testcode/scala/funcs.txt
	scalaCodeFuncs []byte
	//go:embed testcode/scala/classes.txt
	scalaCodeClasses []byte
	//go:embed testcode/scala/objects.txt
	scalaCodeObjects []byte
)

func TestGetScalaFunctions(t *testing.T) {
	expected := []string{
		"def method1(): Unit",
		"def method1(arg1: String, num: Int): Unit",
		"override def method1(arg1: String, num: Int): Unit",
		"private def method1(arg1: String, num: Int): Unit",
		"override protected def method1(arg1: String, num: Int): Unit",
		"def pair[A, B](first: A, second: B): (A, B)",
		"def max[T <: Ordered[T]](list: List[T]): T",
		"def mapContainer[F[_], A, B](container: F[A])(func: A => B)(implicit functor: Functor[F]): F[B]",
	}

	resultChan := make(chan Result)
	go getScalaFunctions(resultChan, scalaCodeFuncs)

	got := <-resultChan

	require.NoError(t, got.Err)
	assert.Equal(t, expected, got.Results)
}

func TestGetScalaClasses(t *testing.T) {
	expected := []string{
		"class MyClass",
		"sealed abstract class KafkaConsumer",
		"sealed abstract class Signature[+T]",
		"class OuterClass",
		"class InnerClass",
		"class MyClass(val name: String, val age: Int)",
		"class MyExtendedClass(name: String, age: Int, val occupation: String) extends MyClass(name, age)",
		`class MyHealthCheck extends HealthCheck[IO]("my-health-check")`,
	}

	resultChan := make(chan Result)
	go getScalaClasses(resultChan, scalaCodeClasses)

	got := <-resultChan
	require.NoError(t, got.Err)
	assert.Equal(t, expected, got.Results)
}

func TestGetScalaObjects(t *testing.T) {
	expected := []string{
		"object HelloWorld",
		"object Container",
	}

	resultChan := make(chan Result)
	go getScalaObjects(resultChan, scalaCodeObjects)

	got := <-resultChan
	require.NoError(t, got.Err)
	assert.Equal(t, expected, got.Results)
}
