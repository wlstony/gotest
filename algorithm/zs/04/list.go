package main

import "fmt"

type SingleList struct {
	Value int
	Next  *SingleList
}

type DoubleList struct {
	Value    int
	Next     *DoubleList
	Previous *DoubleList
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	head := sliceToSingleList(input)
	printSingleList(head)
	head = reverseSingleList(head)
	printSingleList(head)

	input = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	doubleHead := sliceToDoubleList(input)
	printDoubleList(doubleHead)
	doubleHead = reverseDouble(doubleHead)
	printDoubleList(doubleHead)

	list1 := sliceToSingleList([]int{1, 4, 5, 8, 11, 28, 29})
	list2 := sliceToSingleList([]int{8, 28, 29, 30})
	printSame(list1, list2)

	input = []int{1, 2, 3, 3, 2, 1}
	singleNode := sliceToSingleList(input)
	fmt.Println("expect true, output", isPalindrome(singleNode))

	input = []int{1, 2, 3, 4, 3, 2, 1}
	singleNode = sliceToSingleList(input)
	fmt.Println("expect true, output", isPalindrome(singleNode))

	input = []int{1, 2, 3, 4, 3, 2, 1, 1}
	singleNode = sliceToSingleList(input)
	fmt.Println("expect false, output", isPalindrome(singleNode))

	input = []int{1, 2, 3, 4, 3, 2, 1, 1, 9}
	singleNode = sliceToSingleList(input)
	printSingleList(sort(singleNode, 3))

}

func sort(head *SingleList, value int) *SingleList {
	pointer := head
	var sh, st, eh, et, bh, bt *SingleList
	for pointer != nil {
		if pointer.Value < value {
			if sh == nil || st == nil {
				sh, st = pointer, pointer
			} else {
				st.Next = pointer
				st = pointer
			}
		} else if pointer.Value == value {
			if eh == nil || et == nil {
				eh, et = pointer, pointer
			} else {
				et.Next = pointer
				et = pointer
			}
		} else {
			if bh == nil || bt == nil {
				bh, bt = pointer, pointer
			} else {
				bt.Next = pointer
				bt = pointer
			}
		}
		pointer = pointer.Next
	}
	if st != nil {
		st.Next = eh
	} else {
		if et == nil {
			return bh
		}
	}
	if et != nil {
		et.Next = bh
	} else {
		if sh == nil {
			return bh
		}
	}
	return sh
}

func isPalindrome(head *SingleList) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}
	slow, fast, beforeSlow := head, head, head
	for slow != nil && fast != nil {
		beforeSlow = slow
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	//逆序
	var last *SingleList
	current := slow
	for current != nil {
		next := current.Next
		current.Next = last
		last = current
		current = next
	}
	beforeSlow.Next = last
	//判断
	b, h := beforeSlow.Next, head
	for b != nil && h != nil {
		if b.Value != h.Value {
			return false
		}
		b = b.Next
		h = h.Next
	}
	//再恢复顺序
	var from *SingleList
	current = last
	for current != nil {
		next := current.Next
		current.Next = from
		from = current
		current = next
	}
	beforeSlow.Next = from
	return true
}

func printSame(list1 *SingleList, list2 *SingleList) {
	fmt.Print("printSame ")
	for list1 != nil && list2 != nil {
		if list1.Value == list2.Value {
			fmt.Print(list1.Value, " ")
			list2 = list2.Next
			list1 = list1.Next
		} else if list1.Value > list2.Value {
			list2 = list2.Next
		} else {
			list1 = list1.Next
		}
	}
	fmt.Println()
}

func sliceToDoubleList(input []int) *DoubleList {
	var head *DoubleList
	if input == nil || len(input) == 0 {
		return head
	}
	head = &DoubleList{
		Value:    input[0],
		Next:     nil,
		Previous: nil,
	}
	last := head
	for i := 1; i < len(input); i++ {
		current := &DoubleList{
			Value:    input[i],
			Next:     nil,
			Previous: nil,
		}
		current.Previous = last
		last.Next = current

		last = current
	}
	return head
}

func sliceToSingleList(input []int) *SingleList {
	if input == nil || len(input) == 0 {
		return nil
	}
	head := &SingleList{
		Value: input[0],
		Next:  nil,
	}
	last := head
	for i := 1; i < len(input); i++ {
		current := &SingleList{
			Value: input[i],
			Next:  nil,
		}
		last.Next = current
		last = current
	}
	return head
}

func reverseSingleList(head *SingleList) *SingleList {
	var last *SingleList
	for node := head; node != nil; {
		next := node.Next
		node.Next = last
		last = node
		node = next
	}
	return last
}

func printSingleList(head *SingleList) {
	for node := head; node != nil; node = node.Next {
		fmt.Print(" ", node.Value)
	}
	fmt.Println()
}

func printDoubleList(head *DoubleList) {
	fmt.Print("printDoubleDoubleList next ")
	current, last := head, head
	for current != nil {
		fmt.Print(current.Value, " ")
		last = current
		current = current.Next
	}
	fmt.Print(", previous ")
	for last != nil {
		fmt.Print(last.Value, " ")
		last = last.Previous
	}
	fmt.Println()
}

func reverseDouble(head *DoubleList) *DoubleList {
	current, tail := head, head
	for current != nil {
		next, previous := current.Next, current.Previous
		current.Next, current.Previous = previous, next
		tail = current
		current = next
	}
	return tail
}
