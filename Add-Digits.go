package main

func addDigits(num int) int {
	if digitOrNot(num) {
		return num
	}
	if num%9 == 0 {
		return 9
	}
	return addDigits(sumDigits(num))
}

func sumDigits(n int) int {
	var result int
	for i := n; i > 0; i /= 10 {
		result += i % 10
	}
	return result
}

func digitOrNot(n int) bool {
	return n/10 == 0
}