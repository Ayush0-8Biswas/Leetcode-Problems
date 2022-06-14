package main

import "fmt"

func plusOne(digits []int) []int {
	carry, length := 0, len(digits)
	digits[length-1]++
	for length > 0 {
		length--
		digits[length] += carry
		carry = digits[length] / 10
		digits[length] %= 10
	}
	if carry > 0 {
		return append([]int{carry}, digits...)
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{9}))
}
