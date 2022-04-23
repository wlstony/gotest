package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	Value int
}

func main() {
	input := []int{1, 3, 4, 8, 1, 2, 4, 6, 5}
	//input := []int{1,3,4, 8}
	for i := 0; i < len(input); i++ {
		heapInsert(input, i)
	}
	fmt.Println("heapInsert:", input)
	//input = []int{8, 6, 4, 5, 1, 2, 3, 1, 4} //大根堆
	input = []int{4, 6, 4, 5, 1, 2, 3, 1, 8} //大根堆
	heapify(input, 0, len(input))
	fmt.Println("heapify:", input)
	input = []int{1, 3, 4, 8, 1, 2, 4, 6, 5}
	heapSort(input)
	fmt.Println("sort:", input)
}

func heapSort(input []int)  {
	if input == nil || len(input)<2 {
		return
	}
	for i:=0; i<len(input);i++ {
		heapInsert(input, i)
	}

	for i:=len(input) -1; i>=0; i-- {
		swap(input, 0, i)
		heapify(input, 0, i)
	}

}

func heapify(input []int, index, heapSize int) {
	left := index*2 + 1
	for {
		if left >= heapSize {
			break
		}
		largest := index
		if input[largest] < input[left] {
			largest = left
		}
		right := index + 1
		if right < heapSize && input[right] > input[largest] {
			largest = right
		}
		if index == largest {
			break
		}
		swap(input, largest, index)
		index = largest
		left = index*2 + 1
	}
}
func heapInsert(input []int, index int) {
	for {
		head := (index - 1) / 2
		if input[index] <= input[head] {
			break
		}
		swap(input, index, head)
		index = head
	}
}

func swap(input []int, i, j int) {
	if i == j {
		return
	}
	input[i] = input[i] ^ input[j]
	input[j] = input[i] ^ input[j]
	input[i] = input[i] ^ input[j]
}
