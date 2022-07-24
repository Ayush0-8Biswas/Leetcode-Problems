package main

func guess(num int) int {
	return 0
}

func guessNumber(n int) int {
	var first, last = 1, n

	for first < last {
		var mid = (first + last) / 2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == 1 {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}

	return first
}
