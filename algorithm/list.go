package main

import "fmt"

type SNode struct {
	Val  int
	Next *SNode
}
type DNode struct {
	Val      int
	Next     *DNode
	Previous *DNode
}

func main() {
	//input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//head := sliceToSNode(input)
	//node := head
	//outputSNode(node)
	//node = reverseSNode(head)
	//outputSNode(node)
	//-----------------
	//	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//	dnode := sliceToDNode(input)
	//outputDNode(dnode)
	//outputDNode(reverseDNode(dnode))
	//-----------------
	//input1 := sliceToSNode([]int{4, 6, 9, 11})
	//input2 := sliceToSNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	//outPutSame(input1, input2)
	//-----------------
	//bl, node := isPalindrome(sliceToSNode([]int{1, 2, 3, 2, 1}))
	//fmt.Println(bl)
	//outputSNode(node)
	//
	//bl, node = isPalindrome(sliceToSNode([]int{1, 2, 3, 3, 2, 1}))
	//fmt.Println(bl)
	//outputSNode(node)
	//
	//bl, node = isPalindrome(sliceToSNode([]int{1, 2, 3, 3, 5, 1}))
	//fmt.Println(bl)
	//outputSNode(node)
	//-----------------
}
func () *s {
	
}
func isPalindrome(node *SNode) (bool, *SNode) {
	slow, fast := node, node
	for {
		if fast == nil || fast.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	//后半部分逆序
	current := slow
	var last *SNode
	//1, 2, 3, 2, 1->1, 2, 1, 2, 3
	for {
		if current == nil {
			break
		}
		next := current.Next
		current.Next = last
		last = current
		current = next
	}
	first := node
	second := last
	bl := true
	//判断是否回文
	for {
		if first.Next == nil || second.Next == nil {
			break
		}
		if first.Next.Val != second.Next.Val {
			bl = false
			break
		}
		first = first.Next
		second = second.Next
	}
	//return bl, node
	//恢复顺序
	current = last
	var previous *SNode
	for {
		if current == nil {
			break
		}
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}

	return bl, node
}
func outPutSame(node1 *SNode, node2 *SNode) {
	for {
		if node1 == nil || node2 == nil {
			break
		}
		if node1.Val == node2.Val {
			fmt.Println(node1.Val)
			node1 = node1.Next
			node2 = node2.Next
		} else if node1.Val > node2.Val {
			node2 = node2.Next
		} else {
			node1 = node1.Next
		}
	}

}
func reverseSNode(head *SNode) *SNode {
	node := head
	var last, next *SNode
	for {
		if node == nil {
			break
		}
		next = node.Next
		node.Next = last
		last = node
		node = next
	}
	return last

}

func reverseDNode(head *DNode) *DNode {
	current := head
	var previous, next *DNode
	for {
		if current == nil {
			break
		}
		previous = current.Previous
		next = current.Next
		current.Previous = current.Next
		current.Next = previous
		current = next
	}
	return previous.Previous
}
func outputSNode(node *SNode) {
	for {
		if node == nil {
			break
		}
		fmt.Print(node.Val, ",")
		node = node.Next
	}
	fmt.Println()
}
func sliceToSNode(input []int) *SNode {
	if len(input) == 0 {
		return nil
	}
	head := &SNode{
		Val:  input[0],
		Next: nil,
	}
	last := head
	for i := 1; i < len(input); i++ {
		current := &SNode{
			Val:  input[i],
			Next: nil,
		}
		last.Next = current
		last = current
	}
	return head
}

func sliceToDNode(input []int) *DNode {
	if len(input) == 0 {
		return nil
	}
	head := &DNode{
		Val:      input[0],
		Next:     nil,
		Previous: nil,
	}
	last := head
	for i := 1; i < len(input); i++ {
		current := &DNode{
			Val:      input[i],
			Next:     nil,
			Previous: nil,
		}
		current.Previous = last
		last.Next = current
		last = current
	}
	return head
}
func outputDNode(node *DNode) {
	times := 0
	var tmp *DNode
	for {
		if node == nil {
			fmt.Println()
			times++
			if times == 1 {
				node = tmp
			} else if times == 2 {
				break
			}
		}
		tmp = node
		fmt.Print(node.Val, ",")
		if times == 0 {
			node = node.Next
		} else if times == 1 {
			node = node.Previous
		}
	}
	fmt.Println()
}
