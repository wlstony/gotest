package main

import "fmt"

func main() {
	s1 := []int{3, 4, 38}
	s2 := []int{3, 4, 38}
	fmt.Println(mergeSort1(s1, s2))
	unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}

	fmt.Println(mergeSort2(unsorted))
}
func mergeSort1(s1 []int, s2 []int) []int {
	len1, len2 := len(s1), len(s2)
	var num1, num2 int
	data := make([]int, 0)
	i, j := 0, 0
	for {
		if i >= len1 && j >= len2 {
			break
		}
		if i < len1 {
			num1 = s1[i]
		} else {
			if j < len2 {
				data = append(data, s2[j:]...)
				break
			}
		}
		if j < len2 {
			num2 = s2[j]
		} else {
			if i < len1 {
				data = append(data, s1[i:]...)
				break
			}
		}
		if num1 < num2 {
			data = append(data, num1)
			i++
		} else {
			data = append(data, num2)
			j++
		}
	}
	return data
}

func mergeSort2(input []int) []int {
	length := len(input)
	if length < 2 {
		return input
	}
	//sorted := make([]int, 0)
	mid := length / 2
	l := mergeSort2(input[:mid])
	r := mergeSort2(input[mid:])
	return mergeSort1(l, r)
}
