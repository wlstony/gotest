package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("abcdcbefaba"))
	fmt.Println(longestPalindrome("abcddcbefbcb"))
}
func longestPalindrome(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		len1 := getPalindromeIndex(s, i, i)
		len2 := getPalindromeIndex(s, i, i+1)
		current := ""
		if len2 > len1 {
			current = s[i-len2/2+1 : i+len2/2+1]
		} else {
			current = s[i-len1/2 : i+len1/2+1]
		}
		if len(current) > len(longest) {
			longest = current
		}

	}
	return string(longest)
}
func getPalindromeIndex(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
