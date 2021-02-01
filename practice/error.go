package main

func main() {
	// []int, IntSlice and MySlice share
	// the same underlying type: []int
	type IntSlice []int
	type MySlice  []int

	var s  = []int{}
	var is = IntSlice{}
	var ms = MySlice{}
	var x struct{n int `foo`}
	var y struct{n int `bar`}

	// The two implicit conversions both doesn't work.
	/*
		is = ms // error
		ms = is // error
	*/

	// Must use explicit conversions here.
	is = IntSlice(ms)
	ms = MySlice(is)
	x = struct{n int `foo`}(y)
	y = struct{n int `bar`}(x)

	// Implicit conversions are okay here.
	s = is
	is = s
	s = ms
	ms = s
}
