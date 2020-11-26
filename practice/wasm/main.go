package main

import "fmt"

//GOARCH=wasm GOOS=js go build -o lib.wasm main.go
//复制$(go env GOROOT)/misc/wasm/下
//GOARCH=wasm GOOS=js go build -o test.wasm main.go
func main() {
	fmt.Println("hello wasm")
}