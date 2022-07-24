package main

func containsDuplicate(nums []int) bool {
	var occured = make(map[int]bool)

	for _, v := range nums {
		if _, ok := occured[v]; ok {
			return true
		}
		occured[v] = true
	}

	return false
}
