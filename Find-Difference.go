package main

func findTheDifference(s string, t string) byte {
	var result byte
	for _, v := range []byte(s) {
		result ^= v
	}
	for _, v := range []byte(t) {
		result ^= v
	}

	return result
}