package main

import "fmt"

func romanToInt(s string) int {
	dictionary := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	sum, prev := 0, 0
	for _, v := range s {
		current := dictionary[v]

		if prev < current {
			sum -= 2 * prev
		}

		sum += current
		prev = current
	}

	return sum
}

func main() {
	fmt.Println(romanToInt("II"))
	fmt.Println(romanToInt("XII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
