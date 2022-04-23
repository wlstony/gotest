package main

import "fmt"

func main() {
	input := []int{3,2,4,5,5,0, 1}
	processDescOrder(input, 0, len(input) - 1)
}
func processDescOrder(input []int, left, right int)  {
	if left == right {
		return
	}
	mid := left + (right - left) >> 1
	processDescOrder(input, left, mid)
	processDescOrder(input, mid+1, right)
	outputDescOrder(input, left, mid, right)

}
func outputDescOrder(input []int, left, mid, right int)  {
	i, j := left, mid+1
	help := make([]int,0)
	for  {
		if i>mid || j>right  {
			break
		}
		if input[i] > input[j] {
			for t:=j; t<=right; t++ {
				fmt.Println(input[i], ",", input[t])
			}
			help = append(help, input[i])
			i++
		} else {
			help = append(help, input[j])
			j++
		}
	}
	for {
		if i > mid {
			break
		}
		help = append(help, input[i])
		i++
	}
	for {
		if j > right {
			break
		}
		help = append(help, input[j])
		j++
	}
	for x := 0; x< len(help); x++ {
		input[left+x] = help[x]
	}
}
