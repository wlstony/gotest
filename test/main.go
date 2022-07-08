package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	m.Store("k1", 1)
	m.Store("k2", 2)
	m.Store("k3", 3)
	m.Range(func(key, value interface{}) bool {
		fmt.Println("key", key, ", value", value)
		return true
	})
}

func compare( value interface{}, ) bool  {
	return false
}