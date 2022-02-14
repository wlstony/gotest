package main

import "fmt"

func main() {
	data := []int{3, 4, 5, 2, 1, 4, 5}

	fmt.Println(insertSort(data))
	data = []int{1, 2, 6, 7, 5}
	fmt.Println(insertSort(data))
}

func insertSort(data []int) []int {
	length := len(data)
	if length < 2 {
		return data
	}
	index := 0
	for i := 1; i < length; i++ {
		if data[i] > data[i-1] {
			continue
		}
		tmp := data[i]
		for j := i-1; j >= 0; j-- {
			if tmp <= data[j] {
				index = j
				data[j+1] = data[j]
			}
		}
		data[index] = tmp



			//-----------//
			//if j != 0 && data[i] < data[j] && data[i]>=data[j-1] {
			//	tmp := data[i]
			//	for k := i-1; k>=j ; k--   {
			//		data[k+1] = data[k]
			//	}
			//	data[j] = tmp
			//	break
			//}
			//if j == 0 && data[i] < data[0] {
			//	tmp := data[i]
			//	for k := i-1; k>=0 ; k--   {
			//		data[k+1] =data[k]
			//	}
			//	data[0] = tmp
			//	break
			//}

		//}
	}

	return data
}
