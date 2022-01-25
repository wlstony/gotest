package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	wg.Add(3)
	go outPutOne(wg, ch3, ch1)
	go outPutTwo(wg, ch2, ch1)
	go outPutThree(wg, ch3, ch2)
	wg.Wait()
}

func outPutOne(wg *sync.WaitGroup, ch3, ch1 chan string) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		if i > 0 {
			<-ch3
		}
		fmt.Println("one")
		ch1 <- "one"
	}
}

func outPutTwo(wg *sync.WaitGroup, ch2, ch1 chan string) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		<-ch1
		fmt.Println("two")
		ch2 <- "two"
	}
}
func outPutThree(wg *sync.WaitGroup, ch3, ch2 chan string) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		<-ch2
		fmt.Println("three")
		if i < 4 {
			ch3 <- "three"
		}
	}
}
