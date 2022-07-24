package main

func canConstruct(ransomNote string, magazine string) bool {
	var characterCount = make(map[rune]int)
	for _, v := range magazine {
		characterCount[v]++
	}

	for _, v := range ransomNote {
		if x, _ := characterCount[v]; x > 0 {
			characterCount[v]--
		} else {
			return false
		}
	}

	return true
}