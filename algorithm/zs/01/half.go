package main

import "fmt"

func main() {
	input := []int{1,2,3,4,5,6,7,10,12}
	fmt.Println(exist(input, 8))
	fmt.Println(exist(input, 7))
	input = []int{1,2,2,3,3,4,5,6}
	//1
	fmt.Println(findIndex(input, 2, 0))
	//6
	fmt.Println(findIndex(input, 5, 0))

}

func exist(input []int, number int) bool  {
	length := len(input)
	if length <= 0{
		return false
	}
	if length == 1 {
		if number == input[0] {
			return true
		}
		return false
	}
	half := length / 2
	if input[half-1]  >= number {
		return exist(input[:half], number)
	} else {
		return exist(input[half:], number)
	}
	return false
}

func findIndex(input []int, number,baseIndex int) int  {
	if len(input) == 0 {
		return -1
	}
	if len(input) == 1  {
		if  input[0] == number {
			return baseIndex
		}
		return -1
	}

	half := len(input) / 2
	if input[half - 1] >= number {
		return findIndex(input[:half], number, baseIndex)
	} else {
		return findIndex(input[half:], number, half+baseIndex)
	}
	return -1
}