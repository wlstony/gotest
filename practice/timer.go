package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(5 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired, ", time.Now())

	timer2 := time.NewTimer(3 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired, ", time.Now())
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped, ", time.Now())
	}

	time.Sleep(20 * time.Second)
}