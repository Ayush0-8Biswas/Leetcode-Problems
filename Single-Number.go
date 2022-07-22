package main

import "fmt"

func singleNumber(nums []int) int {
	var result = 0

	for _, v := range nums {
		result ^= v
	}

	return result
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}