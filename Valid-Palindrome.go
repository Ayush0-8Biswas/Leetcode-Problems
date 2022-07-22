package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	sByte := []byte(s)

	first, last := 0, len(s)-1
	for first < last {
		for first < last && !isAlphanumeric(sByte[first]) {
			first++
		}
		for first < last && !isAlphanumeric(sByte[last]) {
			last--
		}
		if first < last && sByte[first] != sByte[last] {
			return false
		}
		last--
		first++
	}
	return true
}

func isAlphanumeric(c byte) bool {
	return (c <= 'z' && c >= 'a') || (c >= '0' && c <= '9')
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
}
