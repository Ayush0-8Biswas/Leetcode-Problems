package main

func countBits(n int) []int {
	var result = make([]int, n+1)

	for i := 0; i <= n; i++ {
		result[i] = i&1 + result[i>>1]
	}

	return result
}