package main

func isBadVersion(n int) bool {
	return false
}

func firstBadVersion(n int) int {
	var first, last = 0, n
	for first != last {
		var mid = (first + last) / 2
		if isBadVersion(mid) {
			last = mid
		} else {
			first = mid + 1
		}
	}
	return first
}
