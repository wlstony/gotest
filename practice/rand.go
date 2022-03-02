package main

import (
	"fmt"
	"math/rand"
	"time"
)

var staff = []string{
	"钱能武",
	"戴黎旻",
	"范飞",
	"秦帅",
	"吴林生",
	"夏达",
	"朱锋锦",
}
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	length := len(staff)
	fmt.Println(length)
	rand.Shuffle(length, func(i, j int) {
		staff[i], staff[j] = staff[j], staff[i]
	})
	for _, v := range staff {
		fmt.Println(v)
	}
}