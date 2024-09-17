package tsutils

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testcode/rust.txt
var rustCode []byte

func TestGetRustStructs(t *testing.T) {
	expected := []string{
		"struct Empty;",
		"struct Unit;",
		"struct Color(i32, i32, i32);",
		`struct Table {
    field1: i32,
    field2: String,
}`,
		`struct GenericStruct<T> {
    field: T,
}`,
		`pub struct Point {
    pub x: i32,
    y: i32,
}`,
		`pub(crate) struct Point {
    x: i32,
    y: i32,
}`,
		`pub(super) struct Point {
    x: i32,
    y: i32,
}`,
		`pub(in crate::some_module) struct Point {
    x: i32,
    y: i32,
}`,
	}

	resultChan := make(chan Result)
	go getRustStructs(resultChan, rustCode)

	got := <-resultChan

	require.NoError(t, got.Err)
	assert.Equal(t, expected, got.Results)
}

func TestGetRustEnums(t *testing.T) {
	expected := []string{
		`enum Direction {
    North,
    South,
    East,
    West,
}`,
		`pub enum Option<T> {
    Some(T),
    None,
}`,
		`pub(crate) enum Option<T> {
    Some(T),
    None,
}`,
		`pub(super) enum Option<T> {
    Some(T),
    None,
}`,
		`pub(in crate::some_module) enum Option<T> {
    Some(T),
    None,
}`,
		`enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor { r: u8, g: u8, b: u8 },
}`,
		`enum Result<T, E> {
    Ok(T),
    Err(E),
}`,
		`enum GenericEnum<T> {
    Value(T),
    Nothing,
}`,
		`enum RefEnum<'a, T> {
    Borrowed(&'a T),
    Owned(T),
}`,
	}

	resultChan := make(chan Result)
	go getRustEnums(resultChan, rustCode)

	got := <-resultChan

	assert.Equal(t, expected, got.Results)
	assert.NoError(t, got.Err)
}

func TestGetRustTraits(t *testing.T) {
	expected := []string{
		`trait Drawable {
    fn draw(&self);
}`,
		`trait Resizable<T> {
    fn resize(&mut self, value: T);
}`,
		`trait BasicTrait {
    fn required_method(&self);
}`,
		`trait TraitWithConstants {
    const CONSTANT: u32;
}`,
		`trait TraitWithTypes {
    type ItemType;
}`,
		`trait TraitWithProvidedMethods {
    fn provided_method(&self) {
        println!("This is a provided method.");
    }
}`,
		`trait TraitWithAssociatedFunctions {
    fn associated_function() -> Self;
}`,
		`trait ComprehensiveTrait {
    // Associated constant
    const CONSTANT: u32;

    // Associated type
    type ItemType;

    // Required method
    fn required_method2(&self);

    // Provided method
    fn provided_method2(&self) {
        println!("This is a provided method.");
    }

    // Associated function
    fn associated_function2() -> Self;
}`,
		`trait GenericTrait<T> {
    fn generic_method(&self, value: T);
}`,

		`trait LifetimeTrait<'a> {
    fn lifetime_method(&self, value: &'a str);
}`,
	}

	resultChan := make(chan Result)
	go getRustTraits(resultChan, rustCode)

	got := <-resultChan

	require.NoError(t, got.Err)
	assert.Equal(t, expected, got.Results)
}

func TestGetRustFuncs(t *testing.T) {
	expected := []string{
		"fn provided_method(&self)",
		"fn provided_method2(&self)",
		"fn function_name()",
		"fn generic_function<T>(value: T) -> T",
		"pub fn public_function() -> i32",
		"pub(crate) fn crate_function() -> i32",
		"pub(super) fn parent_function() -> i32",
		"pub(in crate::some_module) fn specific_module_function() -> i32",
		"fn new(field1: i32, field2: String) -> Self",
		"fn apply_to<'a>(&'a self, data: &'a mut Table) -> &'a mut Table",
		"fn update(&mut self, field1: i32)",
		"fn draw(&self)",
		"fn resize(&mut self, value: T)",
		"fn new(field: T) -> Self",
		"fn get_field(&self) -> &T",
	}

	resultChan := make(chan Result)
	go getRustFuncs(resultChan, rustCode)

	got := <-resultChan

	require.NoError(t, got.Err)
	assert.Equal(t, expected, got.Results)
}
