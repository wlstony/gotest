package main

import "fmt"

func main() {
	//44
	input := []int{2, 1, 5, 8, 9, 6, 3, 4}
	fmt.Println(smallSum(&input, 0, len(input)-1))

	//16
	input2 := []int{1, 3, 4, 2, 5}
	fmt.Println(smallSum(&input2, 0, len(input2)-1))
}

func smallSum(input *[]int, left, right int) int {
	if left == right {
		return 0
	}

	mid := left + ((right - left) >> 1)
	ls := smallSum(input, left, mid)
	rs := smallSum(input, mid+1, right)
	ms := mergeSum(input, left, mid, right)

	return ms + rs + ls
}


func mergeSum(input *[]int, left, mid, right int) int {

	sum := 0
	l := left
	r := mid + 1
	help := make([]int, 0)
	//if mid == 2 {
	//fmt.Println("aaa")
	//}
	for {
		if l > mid || r > right {
			break
		}
		if (*input)[l] < (*input)[r] {
			sum += (right - r+1) * (*input)[l]
			help = append(help, (*input)[l])
			l++
		} else {
			help = append(help, (*input)[r])
			r++
		}

	}
	for {
		if l > mid {
			break
		}
		help = append(help, (*input)[l])
		l++
	}

	for {
		if r > right {
			break
		}
		help = append(help, (*input)[r])
		r++
	}
	i := 0
	for  {
		if left > right{
			break
		}
		(*input)[left] = help[i]
		left++
		i++
	}
	return sum
}
