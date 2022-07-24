package main

type NumArray struct {
	sumArray []int
}

func Constructor(nums []int) NumArray {
	var newArray NumArray
	newArray.sumArray = make([]int, len(nums)+1)

	var sum int
	for k, v := range nums {
		sum += v
		newArray.sumArray[k+1] = sum
	}

	return newArray
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sumArray[right+1] - this.sumArray[left]
}
