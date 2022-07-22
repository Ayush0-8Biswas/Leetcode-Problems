package main

func majorityElement(nums []int) int {
	var candidate, votes = -1, 0

	for _, v := range nums {
		if votes == 0 {
			candidate = v
			votes = 1
		} else {
			if v == candidate {
				votes++
			} else {
				votes--
			}
		}
	}

	return candidate
}
