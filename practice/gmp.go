package main

func main()  {
	a := 1
	b := Add(a)
	println("b的值是：", b)
	c := Add2(a)
	println("c的值是：", c)
}
func Add(a int) int {
	a++
	defer func() {
		a++
	}()

	return a
}
func Add2(a int) (d int) {
	d = a
	d++
	defer func() {
		d++
	}()
	return
}