package main

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	res := make([][]int, 2)
	res[0] = make([]int, rowIndex+1)
	res[1] = make([]int, rowIndex+1)

	res[0][0] = 1
	res[1][0], res[1][1] = 1, 1

	for i := 0; i <= rowIndex; i++ {
		nextRow(res[1-i%2], res[i%2], i+1)
	}

	return res[rowIndex%2]
}

func nextRow(previous []int, current []int, index int) {
	current[0] = 1
	for i := 0; i < index-2; i++ {
		current[i+1] = previous[i] + previous[i+1]
	}
	current[index-1] = 1
}

func main() {
	fmt.Println(getRow(22))
}
