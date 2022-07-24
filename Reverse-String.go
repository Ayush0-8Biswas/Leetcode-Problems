package main

func reverseString(s []byte) {
	var first, last = 0, len(s) - 1

	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}
