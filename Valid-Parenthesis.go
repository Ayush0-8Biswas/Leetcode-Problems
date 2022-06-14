package main

import "fmt"

func isValid(s string) bool {
	myStack := make([]rune, 0)
	stackSize := 0
	bracketPair := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, v := range s {
		if v == ')' || v == '}' || v == ']' {
			if stackSize <= 0 {
				return false
			}
			if myStack[stackSize-1] != bracketPair[v] {
				return false
			}
			myStack = myStack[:stackSize-1]
			stackSize--
		} else {
			myStack = append(myStack, v)
			stackSize++
		}
	}

	return len(myStack) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
}
