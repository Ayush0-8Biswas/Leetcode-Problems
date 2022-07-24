package main

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	var x = n
	for x%2 == 0 {
		x /= 2
	}
	for x%3 == 0 {
		x /= 3
	}
	for x%5 == 0 {
		x /= 5
	}

	return x == 1
}
