package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	res := generate(numRows - 1)
	newRow := make([]int, numRows)
	newRow[0], newRow[numRows-1] = 1, 1
	for i := 0; i < numRows-2; i++ {
		newRow[i+1] = res[numRows-2][i] + res[numRows-2][i+1]
	}
	res = append(res, newRow)
	return res
}

func main() {
	fmt.Println(generate(5))
}
