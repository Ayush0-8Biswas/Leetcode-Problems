package main

func moveZeroes(nums []int) {
	var first, second = 0, 0
	var n = len(nums)

	for second < n {
		if nums[second] != 0 {
			nums[second], nums[first] = nums[first], nums[second]
			first++
		}
		second++
	}
}
