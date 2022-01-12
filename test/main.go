package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var responseChannel = make(chan string, 15)

func main() {
	fmt.Println(time.Now())
	go response()
	wg := &sync.WaitGroup{}
	//并发10
	for i := 0; i < 10; i++ {
		//fmt.Println("i:%d", i)
		pc, file, line, ok := runtime.Caller(i)
		fmt.Println("pc:", pc, "file:", file, "line:", line, "ok:", ok)
		wg.Add(1)

		go httpGet(i, wg)
	}
	wg.Wait()
	time.Sleep(time.Second * 1)
	close(responseChannel)
	fmt.Println("close close")
	time.Sleep(time.Second * 1)
	responseChannel <- fmt.Sprintf("Hello Go %d", 1000000000000)
	time.Sleep(time.Second * 1)
	fmt.Println("all Done")
}

func httpGet(no int , wg *sync.WaitGroup) {
	defer wg.Done() //释放一个锁
	responseChannel <- fmt.Sprintf("Hello Go %d", no)
}
func response() {
	for rc := range responseChannel {
		fmt.Println("response:", rc)
	}
	fmt.Println("exit response")
}