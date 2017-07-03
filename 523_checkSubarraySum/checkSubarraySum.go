package checkSubarraySum

func checkSubarraySum(nums []int, k int) bool {
	hash := map[int]int{0: -1}
	var sum = 0
	for i, v := range nums {
		sum += v
		if k != 0 {
			sum = sum % k
		}

		if prev, ok := hash[sum]; ok {
			if i-prev > 1 {
				return true
			}
		} else {
			hash[sum] = i
		}

	}
	return false
}
