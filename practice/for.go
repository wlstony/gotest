package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	newArr := []*int{}
	for _, v := range arr {
		fmt.Println("v:", v)
		newArr = append(newArr, &v)
	}
	fmt.Println(newArr)
	for _, v := range newArr {
		fmt.Println(*v)
	}

	hash := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	for k, v := range hash {
		println(k, v)
	}
}
