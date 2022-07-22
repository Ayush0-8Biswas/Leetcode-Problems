package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	mergedEnd, end1, end2 := m+n-1, m-1, n-1

	for end1 >= 0 && end2 >= 0 {
		if nums1[end1] > nums2[end2] {
			nums1[mergedEnd] = nums1[end1]
			end1--
		} else {
			nums1[mergedEnd] = nums2[end2]
			end2--
		}
		mergedEnd--
	}

	for end2 >= 0 {
		nums1[mergedEnd] = nums2[end2]
		mergedEnd--
		end2--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
