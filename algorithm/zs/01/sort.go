package main

import (
	"fmt"
	"math"
)

func main() {
	// -10, 1, 2,3,4,6,7,8,9
	input := []int{1, 4, 6, 9, 2, 3, 8, 7, -10}
	fmt.Println(selectSort(input))
	input = []int{1, 4, 6, 9, 2, 3, 8, 7, -10}
	fmt.Println(insertSort(input))

}

func selectSort(input []int) []int {
	minValue, minIndex := math.MaxInt64, 0
	for i := 0; i < len(input); i++ {
		minValue = math.MaxInt64
		for j := i; j < len(input); j++ {
			if value := input[j]; value < minValue {
				minValue = value
				minIndex = j
			}
		}
		if i != minIndex {
			swap(&input, i, minIndex)
		}
	}
	return input
}
func swap(input *[]int, i int, j int) {
	(*input)[i] = (*input)[i] ^ (*input)[j]
	(*input)[j] = (*input)[i] ^ (*input)[j]
	(*input)[i] = (*input)[i] ^ (*input)[j]
}

func insertSort(input []int) []int {
	for i := 1; i < len(input); i++ {
		for j := i; j >= 1; j-- {
			if input[j] >= input[j-1] {
				break
			}
			swap(&input, j, j-1)
		}
	}

	return input
}
