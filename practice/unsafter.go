package main

import "fmt"

type LockMap struct {
	m map[interface{}]interface{}
}
func NewLockMap() *LockMap {
	return &LockMap{
		m: map[interface{}]interface{}{},
	}
}
func main() {
	t := NewLockMap()
	fmt.Println(t.m["aaa"])
}