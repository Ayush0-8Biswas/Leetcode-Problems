package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	oneStepBehind, twoStepBehind, totalWays := 2, 1, 0
	for i := 2; i < n; i++ {
		totalWays = oneStepBehind + twoStepBehind
		twoStepBehind = oneStepBehind
		oneStepBehind = totalWays
	}

	return totalWays
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}
