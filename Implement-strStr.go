package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	needleHead := rune(needle[0])
	needleLength := len(needle)
	kLimit := len(haystack) - needleLength

	for k, v := range haystack {
		if k > kLimit {
			break
		}
		if v == needleHead {
			if haystack[k:k+needleLength] == needle {
				return k
			}
		}
	}

	return -1
}

func main() {
	haystack, needle := "aaaaa", "bba"

	fmt.Println(strStr(haystack, needle))
}
