package main

func intersect(nums1 []int, nums2 []int) []int {
	var occurred1 = make(map[int]int)

	for _, v := range nums1 {
		occurred1[v]++
	}

	var result = make([]int, 0, len(nums1))

	for _, v := range nums2 {
		if x, _ := occurred1[v]; x > 0 {
			result = append(result, v)
			occurred1[v]--
		}
	}

	return result
}
