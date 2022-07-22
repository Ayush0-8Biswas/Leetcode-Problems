package main

import "fmt"

func titleToNumber(columnTitle string) int {
	var result int
	for _, v := range columnTitle {
		result = 26*result + int(v-'A'+1)
	}
	return result
}

func main() {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))
}
