package main

func firstUniqChar(s string) int {
	var charOccurrence = make(map[rune]int)
	for k, v := range s {
		if _, ok := charOccurrence[v]; ok {
			charOccurrence[v] = -1
		} else {
			charOccurrence[v] = k
		}
	}

	var result int = len(s)
	for _, v := range charOccurrence {
		if v >= 0 && result > v {
			result = v
		}
	}
	if result == len(s) {
		return -1
	} else {
		return result
	}
}
