package main

import "strings"

func wordPattern(pattern string, s string) bool {
	var patternMapping = make(map[rune]string)
	var words = strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	for k, v := range pattern {
		if prevWord, ok := patternMapping[v]; ok {
			if prevWord != words[k] {
				return false
			}
		} else {
			patternMapping[v] = words[k]
		}
	}

	var uniqueCheck = make(map[string]bool)

	for _, v := range patternMapping {
		if _, ok := uniqueCheck[v]; ok {
			return false
		} else {
			uniqueCheck[v] = true
		}
	}

	return true
}
