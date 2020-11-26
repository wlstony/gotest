package main

import "fmt"

//func main() {
//	defer println("in main")
//	go func() {
//		defer println("in goroutine")
//		panic("aaaa")
//	}()
//
//	time.Sleep(1 * time.Second)
//}


func main() {
	defer func() {
		fmt.Println("in main")
		if err := recover(); err != nil {
			fmt.Println("recover error:", err)
		}
		fmt.Println("end defer")
	}()
	panic("unknown erraaa")
	fmt.Println("end")

}

//func main() {
//	defer fmt.Println("in main")
//	defer func() {
//		defer func() {
//			panic("panic again and again")
//		}()
//		panic("panic again")
//	}()
//
//	go func() {
//		panic("panic once")
//	}()
//	fmt.Println("end")
//}
