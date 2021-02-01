package main

import lua "github.com/yuin/gopher-lua"

//func main() {
//	L := lua.NewState()
//	defer L.Close()
//	if err := L.DoString(`print("hello")`); err != nil {
//		panic(err)
//	}
//}

func Double(L *lua.LState) int {
	lv := L.ToInt(1)             /* get argument */
	L.Push(lua.LNumber(lv * 2)) /* push result */
	return 1                     /* number of results */
}

func main() {
	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("double", L.NewFunction(Double))

	if err := L.DoFile("/opt/case/gotest/lua/hello.lua"); err != nil {
		panic(err)
	}
}