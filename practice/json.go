package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var test interface{}
	str := `{"id":9223372036854775807, "name":"golang"}`
	err := json.Unmarshal([]byte(str), &test)
	fmt.Println("err:", err)
	fmt.Printf("test %#v\n", test)
}
