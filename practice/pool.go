package main
import (
	"fmt"
	"sync"
)

var pool *sync.Pool

type P struct {
	Name string
}

func initPool() {
	pool = &sync.Pool {
		New: func()interface{} {
			fmt.Println("Creating a new Person")
			return new(P)
		},
	}
}

func main() {
	defer func(f func()) {
		fmt.Println("aaaaaaa")
	}(func() {
		fmt.Println("bbbbbbbbbb")
	})
	initPool()

	p := pool.Get().(*P)
	fmt.Println("首次从 pool 里获取：", p)

	p.Name = "first"
	fmt.Printf("设置 p.Name = %s\n", p.Name)

	pool.Put(p)

	fmt.Println("Pool 里已有一个对象：&{first}，调用 Get: ", pool.Get().(*P))
	fmt.Println("Pool 没有对象了，调用 Get: ", pool.Get().(*P))
}