package main

import (
	"fmt"
	"sync"
)

func main() {
	total := 0
	sum := 0
	var lo sync.Mutex
	for i := 1; i <= 10; i++ {
		sum += i
		fmt.Println("i:", i, "sum:", sum)
		lo.Lock()
		go func() {
			total += i
			fmt.Println("go i:", i, ", total:", total)
			lo.Unlock()
		}()
	}
	fmt.Printf("total:%d sum %d\n", total, sum)
}
