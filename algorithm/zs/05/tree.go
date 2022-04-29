package main

import (
	"fmt"
	"math"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	head := GetHead()
	fmt.Print("pre expect 1 2 4 7 8 5 10 3 6 11, real ")
	prePrint(head)
	fmt.Println()
	prePrint1(head)
	fmt.Println()
	fmt.Print("middle expect 7 4 8 2 5 10 1 6 11 3, real ")
	middlePrint(head)
	fmt.Println()
	middlePrint1(head)
	fmt.Println()
	fmt.Print("post expect 7 8 4 10 5 2 11 6 3 1, real ")
	postPrint(head)
	fmt.Println()
	postPrint1(head)
	fmt.Println("width print")
	widthPrint(head)
	fmt.Println("width print")
	fmt.Println("maxWidth:", getMaxWidth(head))

}
func getMaxWidth(head *Node) int {
	if head == nil {
		return 0
	}
	maxWidth, width, level := 0, 0, 1
	nodesLevel := make(map[*Node]int)
	queue := NewQueue()
	nodesLevel[head] = 1
	queue.Push(head)
	for !queue.IsEmpty() {
		head = queue.Pop().(*Node)
		currentLevel := nodesLevel[head]
		//当前层
		if currentLevel == level {
			width++
		} else {
			level = currentLevel
			//换了一层
			width = 1
		}
		maxWidth = int(math.Max(float64(width), float64(maxWidth)))
		if head.Left != nil {
			nodesLevel[head.Left] =  currentLevel + 1
			queue.Push(head.Left)
		}
		if head.Right != nil {
			nodesLevel[head.Right] = currentLevel + 1
			queue.Push(head.Right)
		}
	}
	return maxWidth
}

func widthPrint(head *Node) {
	if head == nil {
		return
	}
	queue := NewQueue()
	queue.Push(head)
	for !queue.IsEmpty() {
		head = queue.Pop().(*Node)
		fmt.Print(head.Value, " ")
		if head.Left != nil {
			queue.Push(head.Left)
		}
		if head.Right != nil {
			queue.Push(head.Right)
		}
	}
}
func prePrint(head *Node) {
	if head == nil {
		return
	}
	node := head
	fmt.Print(node.Value, " ")
	prePrint(node.Left)
	prePrint(node.Right)
}

func prePrint1(head *Node) {
	if head == nil {
		return
	}
	node := head
	stack := NewStack()
	stack.Push(node)
	for node != nil {
		tmpNode := stack.Pop()
		if tmpNode == nil {
			break
		}
		node = tmpNode.(*Node)
		fmt.Print(node.Value, " ")
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
}
func middlePrint(head *Node) {
	if head == nil {
		return
	}
	node := head
	middlePrint(node.Left)
	fmt.Print(node.Value, " ")
	middlePrint(node.Right)
}

func middlePrint1(head *Node) {
	if head == nil {
		return
	}
	stack := NewStack()
	for !stack.IsEmpty() || head != nil {
		if head != nil {
			stack.Push(head)
			head = head.Left
		} else {
			head = stack.Pop().(*Node)
			fmt.Print(head.Value, " ")
			head = head.Right
		}

	}
}
func postPrint(head *Node) {
	if head == nil {
		return
	}
	node := head
	postPrint(node.Left)
	postPrint(node.Right)
	fmt.Print(node.Value, " ")
}

func postPrint1(head *Node) {
	if head == nil {
		return
	}
	stack := NewStack()
	collectStack := NewStack()
	stack.Push(head)
	for !stack.IsEmpty() {
		head = stack.Pop().(*Node)
		collectStack.Push(head)
		if head.Left != nil {
			stack.Push(head.Left)
		}
		if head.Right != nil {
			stack.Push(head.Right)
		}
	}
	for !collectStack.IsEmpty() {
		node := collectStack.Pop().(*Node)
		fmt.Print(node.Value, " ")
	}

}

func GetHead() *Node {
	head := &Node{
		Value: 1,
		Left: &Node{
			Value: 2,
			Left: &Node{
				Value: 4,
				Left: &Node{
					Value: 7,
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					Value: 8,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &Node{
				Value: 5,
				Left:  nil,
				Right: &Node{
					Value: 10,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &Node{
			Value: 3,
			Left: &Node{
				Value: 6,
				Left:  nil,
				Right: &Node{
					Value: 11,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
	}
	return head

}
