package main

import "fmt"

func addBinary(a string, b string) string {
	var res string
	var C byte = 0
	aLen, bLen := len(a)-1, len(b)-1

	for aLen >= 0 && bLen >= 0 {
		A, B := a[aLen]-'0', b[bLen]-'0'
		res = string(A^B^C+'0') + res
		C = ((A ^ B) & C) | (A & B)
		aLen, bLen = aLen-1, bLen-1
	}

	for aLen >= 0 {
		A := a[aLen] - '0'
		res = string(A^C+'0') + res
		C = A & C
		aLen--
	}

	for bLen >= 0 {
		B := b[bLen] - '0'
		res = string(B^C+'0') + res
		C = B & C
		bLen--
	}
	if C > 0 {
		return "1" + res
	}
	return res
}

func main() {
	fmt.Println(addBinary("1", "0"))
}
