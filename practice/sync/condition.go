package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"
)

var status int64

func main() {
	c := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go listen(c, i)
	}
	time.Sleep(60 * time.Second)
	go broadcast(c)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("end")
}

func broadcast(c *sync.Cond) {
	fmt.Println("broadcast")
	c.L.Lock()
	fmt.Println("broadcast lock")

	atomic.StoreInt64(&status, 1)
	c.Broadcast()
	c.L.Unlock()
	fmt.Println("broadcast unlock")

}

func listen(c *sync.Cond, i int) {
	fmt.Println("start listen ", i)
	c.L.Lock()
	fmt.Println("listen lock ", i)

	for atomic.LoadInt64(&status) != 1 {
		fmt.Println("wait", i)
		c.Wait()
	}
	fmt.Println("listen:", i)
	c.L.Unlock()
	fmt.Println("listen unlock:", i)
}
