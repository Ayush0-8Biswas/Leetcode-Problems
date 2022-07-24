package main

func reverseVowels(s string) string {
	var result = []byte(s)
	var first, last = 0, len(result) - 1

	for first < last {
		for !isVowel(s[first]) {
			first++
		}
		for !isVowel(s[last]) {
			last--
		}
		if first < last {
			result[first], result[last] = result[last], result[first]
			first++
			last--
		}
	}
	return string(result)
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}
