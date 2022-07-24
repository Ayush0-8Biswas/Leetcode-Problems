package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	var mapST, mapTS = make(map[rune]byte), make(map[rune]byte)

	for k, v1 := range s {
		if v2, ok := mapST[v1]; ok {
			if t[k] != v2 {
				return false
			}
		} else {
			mapST[v1] = t[k]
		}
	}

	for k, v1 := range t {
		if v2, ok := mapTS[v1]; ok {
			if s[k] != v2 {
				return false
			}
		} else {
			mapTS[v1] = s[k]
		}
	}

	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
}
