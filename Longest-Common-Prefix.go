package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	res := ""
	minLength := len(strs[0])

	var x int
	for _, v := range strs[1:] {
		x = len(v)
		if x < minLength {
			minLength = x
		}
	}

LOOP:
	for i := 0; i < minLength; i++ {
		r := strs[0][i]
		for _, v := range strs[1:] {
			if v[i] != r {
				break LOOP
			}
		}
		res += string(r)
	}

	return res
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	strs2 := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
	fmt.Println(longestCommonPrefix(strs2))
}
