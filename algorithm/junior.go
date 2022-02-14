package main

import "fmt"

func main() {
	input := []int{1,3,4,2,6}
	fmt.Println(sortAsc(input))
	fmt.Println(findSmallNumNotExist(input))
}

func removeDuplicateELement(elements []int) int {
	length := 1
	for i := 1; i < len(elements); i++ {
		d := elements[i]
		if d != elements[i-1] {
			length++
		}
	}
	return length
}

func findSmallNumNotExist(A []int) int {
	A = sortAsc(A)
	// write your code in Go 1.4
	loopCount := len(A) + 1
	for i := 1; i < loopCount; i++ {
		if A[i - 1] != i {
			return i
		}
	}
	return 0
}

func sortAsc(A []int) []int  {
	for i := 0; i<len(A) ; i++  {
		if i == len(A) - 1 {
			continue
		}
		for j := 0; j< len(A) - i; j++ {
			if j == len(A) - i-1 {
				continue
			}
			if A[j] > A[j+1] {
				tmp := A[j+1]
				A[j+1] = A[j]
				A[j] = tmp
			}
		}
	}
	return A
}