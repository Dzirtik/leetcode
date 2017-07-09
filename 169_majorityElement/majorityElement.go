package majorityElement

func majorityElement(nums []int) int {
	var major, counts int
	for i := 0; i < len(nums); i++ {
		if counts == 0 {
			major = nums[i]
			counts = 1
		} else {
			if nums[i] == major {
				counts++
			} else {
				counts--
			}
		}
	}
	return major

}
