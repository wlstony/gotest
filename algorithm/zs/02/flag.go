package main

import "fmt"

func main() {
	input := []int{2, 4, 5, 7, 9, 1, 2}
	//[2 1 2 7 9 4 5]
	fmt.Println(flagOrder(input, 2))

	input = []int{2, 4, 5, 7, 3, 3, 3, 9, 6}
	//[2 1 2 3 1 3 2 3 5 6 9 4 5 7]
	fmt.Println(flagOrder(input, 3))

}

//给定一个数组和一个数，请把小于等于num的数放在数组的左边，大于num的数放在数组的右边
func flagOrder(input []int, number int) []int {
	ltOeq := -1
	for i := 0; i < len(input); i++ {
		if input[i] <= number {
			ltOeq++
			if i-ltOeq >= 1 {
				//需要交换
				input[i] = input[i] ^ input[ltOeq]
				input[ltOeq] = input[i] ^ input[ltOeq]
				input[i] = input[i] ^ input[ltOeq]
			}
		}
	}
	return input
}

//给定一个数组arr和一个数组num，请把小于num的数放在数组的左边，等于num的数放在中间，
//大于num的数放在右边
func flagOrder2(input []int, number int) []int {
	lt, mt, i := -1, len(input), 0
	for {
		if lt == i || mt == i {
			break
		}
		if input[i] < number {
			lt++
			//需要交换数据
			if i-lt >= 1 {
				input[i] = input[i] ^ input[lt]
				input[lt] = input[i] ^ input[lt]
				input[i] = input[i] ^ input[lt]
			}

			i++
		} else if input[i] > number {
			mt--
			input[i] = input[i] ^ input[mt]
			input[mt] = input[i] ^ input[mt]
			input[i] = input[i] ^ input[mt]
		} else {
			i++
		}
	}
	return input
}
