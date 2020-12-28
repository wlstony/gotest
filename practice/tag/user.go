package main

import (
	"fmt"
	"reflect"
)

// Name of the struct tag used in examples
const tagName = "validate1"

type User struct {
	Id    int    `validate:"-"`
	Name  string `validate2:"presence,min=2,max=32"`
	Email string `validate1:"email,required"`
}

func main() {
	letters := []string{"a", "b", "c", "d"}
	fmt.Println(letters[0:0])
	user := User{
		Id:    1,
		Name:  "John Doe",
		Email: "john@example",
	}

	// TypeOf returns the reflection Type that represents the dynamic type of variable.
	// If variable is a nil interface value, TypeOf returns nil.
	t := reflect.TypeOf(user)

	// Get the type and kind of our user variable
	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())

	// Iterate over all available fields and read the tag value
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)

		// Get the field tag value
		tag := field.Tag.Get(tagName)

		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}
}
