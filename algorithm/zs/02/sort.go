package main

import "fmt"

func main() {
	// -10, 1, 2,3,4,6,7,8,9
	input := []int{1, 4, 6, 9, 2, 3, 8, 7, -10}
	fmt.Println(process(input))

}
func process(input []int) []int  {
	if len(input) <= 1 {
		return input
	}
	half := len(input) / 2
	left := input[:half]
	right := input[half:]
	return mergeSort(process(left),process(right))
}

func mergeSort(in1 []int, in2 []int) []int {
	help := make([]int, 0)
	l1, l2 := len(in1), len(in2)
	i, j := 0, 0
	for {

		if i >= l1 || j >= l2 {
			break
		}
		if in1[i] <= in2[j] {
			help = append(help, in1[i])
			i++
		} else {
			help = append(help, in2[j])
			j++
		}

	}
	for {
		if i >= l1 {
			break
		}
		help = append(help, in1[i])
		i++
	}

	for {
		if j >= l2 {
			break
		}
		help = append(help, in2[j])
		j++
	}
	return help
}
