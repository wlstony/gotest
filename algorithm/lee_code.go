package main

import "fmt"

func main() {
	//input := []int{2, 7, 11, 15}
	//fmt.Println(twoSum(input, 9))
	//--------
	//l1 := sliceToNode(	[]int{9,9,9,9,9,9,9})
	//l2 := sliceToNode([]int{9,9,9,9})
	//fmt.Println(addTwoNumbers(l1, l2))
	//----------
	//fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	//----------
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 2}
	nums2 = []int{3, 4, 6, 8}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	length := len1 + len2
	pointer := 0
	var i, j int
	median := 0.0
	//奇数个
	if length%2 != 0 {
		pointer = length/2 + 1
		for {
			if i >= len1 || j >= len2 {
				break
			}

			if nums2[j] > nums1[i] {
				median = float64(nums1[i])
				i++
			} else {
				median = float64(nums2[j])
				j++
			}
			pointer--
			if pointer == 0 {
				return median
			}
		}
		for t := i; t < len1; t++ {
			median = float64(nums1[t])
			pointer--
			if pointer == 0 {
				return median
			}
		}
		for t := j; t < len2; t++ {
			median = float64(nums2[t])
			pointer--
			if pointer == 0 {
				return median
			}
		}
		//偶数个
	} else {
		numbers := make([]float64, 0)
		pointer = length/2
		for {
			if i >= len1 || j >= len2 {
				break
			}

			if nums2[j] > nums1[i] {
				median = float64(nums1[i])
				i++
			} else {
				median = float64(nums1[j])
				j++
			}
			pointer--
			if pointer == 0 || pointer == -1 {
				numbers = append(numbers, median)
			}
		}
		for t := i; t < len1; t++ {
			median = float64(nums1[t])
			pointer--
			if pointer == 0 || pointer == -1 {
				numbers = append(numbers, median)
			}
		}
		for t := j; t < len2; t++ {
			median = float64(nums2[t])
			pointer--
			if pointer == 0 || pointer == -1 {
				numbers = append(numbers, median)
			}
		}
		median = (numbers[0] + numbers[1]) / 2
	}
	return median
}

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	str1 := make(map[byte]byte)
	str2 := make(map[byte]byte)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			b := s[j]
			if _, has := str1[b]; !has {
				str1[b] = b
			} else {
				if len(str1) > len(str2) {
					str2 = str1
				} else if len(str1) == len(str2) {
					fmt.Sprintf("str1 is equal to str2 %s:%s", str1, str2)
				}
				str1 = make(map[byte]byte)
				break
			}
		}
	}
	if len(str2) >= len(str1) {
		return len(str2)
	}

	return len(str1)
}

//Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.
func twoSum(input []int, target int) []int {
	length := len(input)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if target == (input[i] + input[j]) {
				return []int{i, j}
			}
		}
	}
	return nil
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func sliceToNode(data []int) *ListNode {
	length := len(data)
	if length <= 0 {
		return nil
	}
	head := &ListNode{
		Val:  data[0],
		Next: nil,
	}
	current := head
	for i := 1; i < length; i++ {
		node := &ListNode{
			Val:  data[i],
			Next: nil,
		}
		current.Next = node
		current = node
	}
	return head
}

//You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node1 := l1
	node2 := l2
	num := 0
	var node, head *ListNode
	for {
		value := num
		if node1 == nil && node2 == nil {
			if value > 0 {
				if head == nil {
					node = &ListNode{
						Val:  value,
						Next: nil,
					}
					head = node
				} else {
					tNode := &ListNode{
						Val:  value,
						Next: nil,
					}
					node.Next = tNode
					node = tNode
				}

			}
			break
		}

		if node1 != nil {
			value += node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			value += +node2.Val
			node2 = node2.Next
		}
		if value < 10 {
			if head == nil {
				node = &ListNode{
					Val:  value,
					Next: nil,
				}
				head = node
			} else {
				tNode := &ListNode{
					Val:  value,
					Next: nil,
				}
				node.Next = tNode
				node = tNode
			}
			num = 0
		} else {
			mod := value % 10
			if head == nil {
				node = &ListNode{
					Val:  mod,
					Next: nil,
				}
				head = node
			} else {
				tNode := &ListNode{
					Val:  mod,
					Next: nil,
				}
				node.Next = tNode
				node = tNode
			}
			num = value / 10
		}
	}
	return head
}
