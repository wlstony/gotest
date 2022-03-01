package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i<10; i++ {
		go func(i int) {
			fmt.Println("i is ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("finished")
}