package main

import (
	"strconv"
	"syscall/js"
)

func sum(args []js.Value) {
	var sum int
	for _, val := range args {
		sum += val.Int()
	}
	println(sum)
}

func registerCallbacks() {
	global := js.Global()
	document := global.Get("document")

	getElementById := func(id string) js.Value {
		return document.Call("getElementById", id)
	}
	aValue := getElementById("aValue")
	bValue := getElementById("bValue")
	cValue := getElementById("cValue")
	sumValue := getElementById("sum")

	//sumButton := getElementById("sumButton")
	//runButton := getElementById("runButton")

	//onRun := js.NewCallback(func(args []js.Value) {
	//	println("button on click")
	//})
	//onSum := js.NewCallback(func(args []js.Value) {
	//	a, _ := strconv.Atoi(aValue.Get("value").String())
	//	b, _ := strconv.Atoi(bValue.Get("value").String())
	//	c, _ := strconv.Atoi(cValue.Get("value").String())
	//	sumValue.Set("value", js.ValueOf(a+b+c))
	//})

	addSum := func(v js.Value, i []js.Value) interface{} {
		a, _ := strconv.Atoi(aValue.Get("value").String())
		b, _ := strconv.Atoi(bValue.Get("value").String())
		c, _ := strconv.Atoi(cValue.Get("value").String())
		sumValue.Set("value", js.ValueOf(a+b+c))
		return js.ValueOf(a + b + c)
	}
	global.Set("sum", js.FuncOf(addSum))
	//sumButton.Call("addEventListener", "click", addSum)
	//runButton.Call("addEventListener", "click", addSum)
}

func main() {
	c := make(chan struct{}, 0)
	println("hello wasm")
	registerCallbacks()
	<-c
}
