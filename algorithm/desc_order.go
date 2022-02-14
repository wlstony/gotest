package main

import "fmt"

func main() {

	input := []int{2, 1, 5, 8, 9, 6, 3, 4}
	descOrder(&input, 0, len(input)-1)
}

func descOrder(input *[]int, left, right int) {
	if left == right {
		return
	}
	mid := left + (right-left)>>2
	descOrder(input, left, mid)
	descOrder(input, mid+1, right)
	output(input, left, mid, right)
}

func output(input *[]int, left, mid, right int) {
	//if left == 5{
	//	fmt.Println(input, "left:", left, ",mid:", mid, ",right:", right)
	//}
	l, r := left, mid+1
	help := make([]int, 0)
	for {
		if l > mid || r > right {
			break
		}
		if (*input)[l] > (*input)[r] {
			help = append(help, (*input)[l])
			for i := r; i <= right; i++ {
				fmt.Println((*input)[l], ",", (*input)[i])
			}
			l++
		} else {
			help = append(help, (*input)[r])
			r++
		}

	}
	for i := l;i<=mid;i++ {
		help = append(help, (*input)[i])
	}
	for i := r;i<=right;i++ {
		help = append(help,(*input)[i])
	}
	j := 0
	for i := left; i <= right; i++{
		//if i == 5{
		//	fmt.Println(input, "left:", left, ",mid:", mid, ",right:", right)
		//}


		(*input)[i] = help[j]
		j++
	}
}
