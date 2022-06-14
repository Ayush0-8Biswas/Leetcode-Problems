package main

import "fmt"

func removeElement(nums []int, val int) int {
	n := len(nums)
	fast, slow := n-1, 0

	for slow <= fast {
		for slow < n && nums[slow] != val && slow <= fast {
			slow++
		}
		for fast >= 0 && nums[fast] == val {
			fast--
		}
		if slow < fast {
			nums[slow] = nums[fast]
			fast--
			slow++
		}
	}
	//fmt.Println(fast, slow)
	return slow
}

func main() {
	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums, 3))
	fmt.Println(nums)
}
