package main

import "fmt"

func main() {
	input := []int{1, 3, 4, 2, 5, 5}
	fmt.Println(sumProcess(input, 0, len(input)-1))
}

func sumProcess(input []int, left, right int) int {
	if len(input) <= 1 || left == right {
		return 0
	}
	if left > right || left < 0 || right >= len(input) {
		panic(fmt.Sprintf("parameter error, left %d, right %d, length %d ", left, right, len(input)))
	}
	mid := left + (right-left)>>1

	return sumProcess(input, left, mid) + sumProcess(input, mid+1, right) + orderSum(input, left, mid, right)
}

func orderSum(input []int, left, mid, right int) int {
	help := make([]int, 0)
	i, j := left, mid+1
	sum := 0
	for {
		if i>mid || j>right {
			break
		}
		if input[i] < input[j] {
			sum += input[i] * (right - j +1)
			help = append(help, input[i])
			i++
		} else {
			help = append(help, input[j])
			j++
		}
	}
	for  {
		if i > mid {
			break
		}
		help = append(help, input[i])
		i++
	}
	for  {
		if j > right {
			break
		}
		help = append(help, input[i])
		j++
	}
	for _, v := range help {
		input[left] = v
		left++
	}
	return sum
}
