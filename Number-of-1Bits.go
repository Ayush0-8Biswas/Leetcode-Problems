package main

func hammingWeight(num uint32) int {
	var result int
	for num != 0 {
		result++
		num &= num - 1
	}
	return result
}
