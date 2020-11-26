package main

import (
	"fmt"
)

type TestStruct struct {
	testFunc
}
type testFunc func() error

func (t testFunc) Print()  {
	fmt.Println("Print")
}

func (t *testFunc) Process() error {
	return nil
}

func newTestStruct() TestStruct {
	s := TestStruct{}
	s.testFunc = s.Process
	return s
}
func main() {
	fmt.Println("start")
	newTestStruct().Print()
	
}