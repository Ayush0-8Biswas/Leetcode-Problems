package main

import "fmt"

func isHappy(n int) bool {
	for !isDigit(n) {
		n = sumSquares(n)
	}
	fmt.Println(n)
	return n == 1 || n == 7
}

func sumSquares(n int) int {
	var result int
	for i := n; i > 0; i /= 10 {
		result += square(i % 10)
	}
	return result
}

func square(n int) int {
	return n * n
}

func isDigit(n int) bool {
	return n/10 == 0
}

func main() {
	fmt.Println(isHappy(19))
}
