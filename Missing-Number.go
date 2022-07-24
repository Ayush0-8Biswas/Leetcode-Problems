package main

func missingNumber(nums []int) int {
	var result int
	var n = len(nums)

	for i := 0; i < n; i++ {
		result ^= nums[i] ^ i
	}
	result ^= n

	return result
}
