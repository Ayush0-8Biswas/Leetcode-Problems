package main

import "fmt"

func searchInsert(nums []int, target int) int {
	first, last := 0, len(nums)-1
	var mid, midElem int

	for first <= last {
		mid = (first + last) / 2
		midElem = nums[mid]
		if midElem == target {
			return mid
		} else if midElem > target {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}

	if nums[mid] > target {
		return mid
	} else {
		return mid + 1
	}
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 7
	fmt.Println(searchInsert(nums, target))
}
