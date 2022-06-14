package main

import "fmt"

func maxSubArray(nums []int) int {
	globalMax, endingPointer := nums[0], 0

	for _, v := range nums[0:] {
		endingPointer += v
		if globalMax < endingPointer {
			globalMax = endingPointer
		}
		if endingPointer < 0 {
			endingPointer = 0
		}
	}

	return globalMax
}

func main() {
	nums := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(nums))
}
