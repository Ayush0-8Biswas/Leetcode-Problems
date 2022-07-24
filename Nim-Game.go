package main

func canWinNim(n int) bool {
	if n < 4 {
		return true
	}
	return n%4 != 0
}
