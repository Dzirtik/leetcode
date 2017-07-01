package isUgly

func isUgly(num int) bool {
	if num == 1 {
		return true
	}
	if num < 1 {
		return false
	}
	for {
		if num%2 == 0 {
			num = num / 2
		} else if num%3 == 0 {
			num = num / 3
		} else if num%5 == 0 {
			num = num / 5
		} else if num == 1 {
			return true
		} else {
			return false
		}
	}
}
