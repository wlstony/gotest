package main

import "fmt"

type Person struct {
	wk walk
}
type walk []string

func (w walk) run() {
	fmt.Println("hello run")
}
func main() {
	Person{wk:walk{}}.wk.run()
}