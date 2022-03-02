package main

import (
	"algo/tools"
	"fmt"
)

type Tree struct {
	value int
	left *Tree
	right *Tree
}

func main() {
	head := makeTree()
	fmt.Println("preOutput")
	preOutput(head)
	fmt.Println()
	preOutput1(head)
	fmt.Println()
	fmt.Println("midOutput")
	midOutput(head)
	fmt.Println()
	midOutput1(head)
	fmt.Println()
	fmt.Println("postOutput")
	postOutput(head)

}
func preOutput(node *Tree)  {
	if node == nil {
		return
	}
	fmt.Print(node.value, ",")
	preOutput(node.left)
	preOutput(node.right)
}
//头->左->右
func preOutput1(node *Tree)  {
	stack := tools.NewStack()
	stack.Push(node)
	for {
		in := stack.Pop()
		if in == nil {
			break
		}
		node := in.(*Tree)
		fmt.Print(node.value, ",")
		if node.right != nil {
			stack.Push(node.right)

		}
		if node.left != nil {
			stack.Push(node.left)
		}
	}
}
func midOutput(node *Tree)  {
	if node == nil {
		return
	}
	midOutput(node.left)
	fmt.Print(node.value, ",")
	midOutput(node.right)

}
func midOutput1(node *Tree)  {
	stack := tools.NewStack()
	stack.Push(node)
	tmp := node.left
	for  {
		tmp = tmp.left
		if tmp == nil {
			break
		}
		stack.Push(tmp)
	}
}
func postOutput(node *Tree)  {
	if node == nil {
		return
	}
	postOutput(node.left)
	postOutput(node.right)
	fmt.Print(node.value, ",")
}

/*
		    1
	    /      \
	  2         3
		\      /
         4     5
		/ \   / \
       6   7 8  9
 */
func makeTree() *Tree  {
	node6 := Tree{
		value: 6,
		left:  nil,
		right: nil,
	}
	node7 := Tree{
		value: 7,
		left:  nil,
		right: nil,
	}
	node8 := Tree{
		value: 8,
		left:  nil,
		right: nil,
	}
	node9 := Tree{
		value: 9,
		left:  nil,
		right: nil,
	}
	node4 := Tree{
		value: 4,
		left: &node6,
		right: &node7,
	}

	node5:= Tree{
		value: 5,
		left:  &node8,
		right: &node9,
	}
	node3 :=Tree{
		value: 3,
		left:  &node5,
		right: nil,
	}
	node2 :=Tree{
		value: 2,
		left:  nil,
		right: &node4,
	}
	node1 :=Tree{
		value: 1,
		left:  &node2,
		right: &node3,
	}
	return &node1
		
}