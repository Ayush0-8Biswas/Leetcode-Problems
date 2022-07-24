package main

func isSubsequence(s string, t string) bool {
	var sCounter, tCounter int
	for tCounter < len(t) && sCounter < len(s) {
		if s[sCounter] == t[tCounter] {
			sCounter++
		}
		tCounter++
	}
	return sCounter == len(s)
}
