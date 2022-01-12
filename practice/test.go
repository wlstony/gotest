package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	x int
	y *[1<<23]byte
}

func main() {
	t := T{y: new([1<<23]byte)}
	p := uintptr(unsafe.Pointer(&t.y[0]))

	//... // use T.x and T.y

	// A smart compiler can detect that the value
	// t.y will never be used again and think the
	// memory block hosting t.y can be collected now.

	// Using *(*byte)(unsafe.Pointer(p))) is
	// dangerous here.

	// Continue using value t, but only use its x field.
	println(t.x)
	fmt.Println(p)
}