package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}

	digitCount := int(math.Log10(float64(x))) + 1
	mask := int(math.Pow10(digitCount - 1))

	for i := 0; i < digitCount/2; i++ {
		if x%10 != x/mask {
			return false
		}
		x = x % mask
		x /= 10
		mask /= 100
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(100021))
}
