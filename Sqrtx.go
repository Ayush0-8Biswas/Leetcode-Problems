package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	first, last := 0, x

	for {
		mid := (first + last) / 2
		if mid <= x/mid && (mid+1) > x/(mid+1) {
			return mid
		} else if mid < x/mid {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}

}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(9))
}
