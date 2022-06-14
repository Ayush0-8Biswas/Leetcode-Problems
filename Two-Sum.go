package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for k, v := range nums {
		if k2, ok := seen[target-v]; ok {
			return []int{k2, k}
		}
		seen[v] = k
	}

	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
