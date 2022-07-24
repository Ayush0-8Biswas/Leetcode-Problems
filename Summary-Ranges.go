package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var result = make([]string, 0)

	var first, last = nums[0], nums[0]
	for _, v := range nums {
		if v-last > 1 {
			if first != last {
				result = append(result, fmt.Sprintf("%v->%v", first, last))
			} else {
				result = append(result, strconv.Itoa(first))
			}
			first, last = v, v
		} else {
			last = v
		}
	}

	if first != last {
		result = append(result, fmt.Sprintf("%v->%v", first, last))
	} else {
		result = append(result, strconv.Itoa(first))
	}

	return result
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
