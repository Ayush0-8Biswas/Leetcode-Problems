package main

import "fmt"

func convertToTitle(columnNumber int) string {
	var result = make([]byte, 10)
	var pos = 9

	for columnNumber > 0 {
		columnNumber -= 1
		result[pos] = byte(columnNumber%26) + 'A'
		columnNumber = columnNumber / 26
		pos--
	}

	return string(result[pos+1:])
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func main() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
	fmt.Println(convertToTitle(2147483647))
}
