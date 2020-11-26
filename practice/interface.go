package main

import "fmt"

func main() {
	var i interface{}
	fmt.Printf("%T-%+v\n", i)
}