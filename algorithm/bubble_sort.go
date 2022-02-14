package main

import "fmt"

func main() {
	data := []int{3, 4, 5}

	fmt.Println(bubbleSort1(data))
	fmt.Println(bubbleSort3(data))
	fmt.Println(bubbleSort2(data, len(data)-1))
}
func bubbleSort1(data []int) []int {
	if len(data) < 2 {
		return data
	}
	for j := 0; j < len(data)-1; j++ {
		for i := 0; i < len(data)-j-1; i++ {
			if data[i] > data[i+1] {
				tmp := data[i+1]
				data[i+1] = data[i]
				data[i] = tmp
			}
		}
	}
	return data
}

func bubbleSort2(data []int, length int) []int {
	if len(data) < 2 {
		return data
	}
	if length <= 0 {
		return data
	}
	for i := 0; i < len(data)-1; i++ {
		if data[i] > data[i+1] {
			tmp := data[i+1]
			data[i+1] = data[i]
			data[i] = tmp
		}
	}

	return bubbleSort2(data, length-1)
}

func bubbleSort3(data []int) []int {
	if len(data) < 2 {
		return data
	}
	bl := false
	for j := 0; j < len(data)-1; j++ {
		bl = false
		for i := 0; i < len(data)-j-1; i++ {
			if data[i] > data[i+1] {
				tmp := data[i+1]
				data[i+1] = data[i]
				data[i] = tmp
				bl = true
			}
		}
		//某次循环未发生交换，退出即可
		if !bl {
			break
		}
	}
	return data
}
