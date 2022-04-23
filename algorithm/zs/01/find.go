package main

import "fmt"

func main() {
	input := []int{2, 3, 2, 4, 4}
	fmt.Println(findOne(input))
	input = []int{2, 3, 2, 5, 4, 5, 5, 4}
	fmt.Println(findTwo(input))
}

//某个数出现奇数次，其它偶数次，找到这个数
func findOne(input []int) int {
	var tmp int
	for i := 0; i < len(input); i++ {
		tmp ^= input[i]
	}
	return tmp
}

//某两个不一样数出现奇数次，其它偶数次，找到这两个数
func findTwo(input []int) (int, int) {
	//or为两个数异或
	var or, n1, n2 int
	for i := 0; i < len(input); i++ {
		or ^= input[i]
	}
	//两个不一样的数必定某位为1，找到某个为1的位
	help := (^or + 1) & or
	//将数分为两半
	var s1, s2 []int

	for i := 0; i < len(input); i++ {
		if help&input[i] == 0 {
			s1 = append(s1, input[i])
		} else {
			s2 = append(s2, input[i])
		}
	}
	for i := 0; i < len(s1); i++ {
		n1 ^= s1[i]
	}
	for i := 0; i < len(s2); i++ {
		n2 ^= s2[i]
	}

	return n1, n2
}
