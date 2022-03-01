package main

import (
	"fmt"
	"time"
)

func main() {
	guard := make(chan int, 3)
	seq := 0
	for {
		seq++
		guard <- seq // would block if guard channel is already filled
		go func() {
			fmt.Println("-----------")
			fmt.Println("seq:", seq)
			time.Sleep(time.Second * 1)
			defer func() {
				s := <- guard
				fmt.Println("defer ", s)
			}()
		}()
	}
}
