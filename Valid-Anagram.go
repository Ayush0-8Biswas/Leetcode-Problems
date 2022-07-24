package main

func isAnagram(s string, t string) bool {
	var scounting, tcounting = make(map[rune]int), make(map[rune]int)

	for _, v := range s {
		scounting[v]++
	}
	for _, v := range t {
		tcounting[v]++
	}

	if len(s) != len(t) {
		return false
	}

	for k, v := range scounting {
		if v2, ok := tcounting[k]; ok {
			if v != v2 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
