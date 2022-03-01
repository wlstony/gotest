package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	for i := 0; i <= len(nums); i++ {
		num := nums[i]
		diff := target - num
	}
	return res
}
