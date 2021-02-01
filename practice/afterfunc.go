package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"
)

func main() {
	time.Sleep(10 * time.Second)
	debug.SetTraceback("system")
	if len(os.Args) == 1 {
		panic("before timers")
	}
	for i := 0; i < 10000; i++ {
		time.AfterFunc(time.Duration(5*time.Second), func() {
			fmt.Println("Hello!")
		})
	}

	panic("after timers")
}