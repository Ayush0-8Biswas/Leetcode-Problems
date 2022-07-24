package main

func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}

	var count int
	for n&1 != 1 {
		n >>= 1
		count++
	}

	return (n == 1) && (count%2 == 0)
}
