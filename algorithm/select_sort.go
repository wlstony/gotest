package main

import (
	"fmt"
)

func main() {
	data := []int{3, 4, 5, 2, 1, 4, 5}
	fmt.Println(selectSort(data))
	fmt.Println(selectSort1(data, 0))
}

//从头至尾扫描序列，找出最小的一个元素
//和第一个元素交换
func selectSort(data []int) []int {
	length := len(data)
	if length <= 1 {
		return data
	}
	var minIndex, maxIndex int
	var min, max int

	for i := 0; i < length/2; i++ {
		minIndex = i
		maxIndex = i
		min = data[i]
		max = data[i]
		for j := i + 1; j < length-i; j++ {
			if min > data[j] {
				minIndex = j
				min = data[j]
			}
			if max < data[j] {
				maxIndex = j
				max = data[j]
			}
		}
		if minIndex == maxIndex {
			fmt.Println("minIndex, maxIndex", minIndex, maxIndex)
			continue
		}
		if minIndex != i {
			tmp := data[i]
			data[i] = min
			data[minIndex] = tmp
		}
		if maxIndex != i {
			tmp := data[length-i-1]
			data[length-i-1] = max
			data[maxIndex] = tmp
		}

	}
	return data
}

//递归
func selectSort1(data []int, start int) []int {
	length := len(data)
	if length <= 1 || start >= len(data) {
		return data
	}
	num := data[start]
	index := start
	for i := start + 1; i < length; i++ {
		if num > data[i] {
			index = i
			num = data[i]
		}
	}
	if index != start {
		tmp := data[index]
		data[index] = data[start]
		data[start] = tmp
	}
	return selectSort1(data, start+1)
}
