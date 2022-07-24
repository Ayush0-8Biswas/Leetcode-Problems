package main

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	var first, last = 0, num

	for first < last {
		var x = (first + last) / 2
		var t = num / x
		if x > t {
			last = x - 1
		} else if x < t {
			first = x + 1
		} else {
			return x*x == num
		}
	}

	return first*first == num
}
