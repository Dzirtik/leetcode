package removeKdigits

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	var result string
	keep := len(num) - k
	for i := 0; i < len(num); i++ {
		for len(result) > 0 && result[len(result)-1] > num[i] && k > 0 {
			result = result[0 : len(result)-1]
			k--
		}
		result += string(num[i])
	}
	if len(result) > keep {
		result = result[0:keep]
	}
	for result[0] == '0' && len(result) > 1 {
		result = result[1:]
	}

	if result == "" {
		return "0"
	}
	return result
}
