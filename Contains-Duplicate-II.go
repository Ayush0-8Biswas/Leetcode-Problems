package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	var occurredPosition = make(map[int]int)
	var deletePosition int

	for k1, v := range nums {
		if k2, ok := occurredPosition[v]; ok {
			if k1-k2 <= k {
				return true
			}
		}
		if len(occurredPosition) > k {
			delete(occurredPosition, nums[deletePosition])
			deletePosition++
		}
		occurredPosition[v] = k1
	}

	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
