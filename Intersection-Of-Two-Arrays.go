package main

func intersection(nums1 []int, nums2 []int) []int {
	var occurred = make(map[int]int)
	for _, v := range nums1 {
		occurred[v] = 1
	}

	for _, v := range nums2 {
		if _, ok := occurred[v]; ok {
			occurred[v] = 2
		}
	}

	var result = make([]int, 0, len(occurred))
	for k, v := range occurred {
		if v == 2 {
			result = append(result, k)
		}
	}

	return result
}
