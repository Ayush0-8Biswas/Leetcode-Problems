package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)

	fast, slow := 0, 0

	for fast < n {
		for fast < n && nums[fast] == nums[slow] {
			fast++
		}
		if fast < n {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

func main() {
	nums := []int{1, 1, 2}
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
	fmt.Println(removeDuplicates(nums2))
	fmt.Println(nums2)
}
