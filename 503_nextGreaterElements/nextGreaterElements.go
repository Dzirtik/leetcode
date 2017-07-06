package nextGreaterElements

func nextGreaterElements(nums []int) []int {
	var stack []int
	result := make([]int, len(nums))
	for i := 2*len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[0]] <= nums[i%len(nums)] {
			stack = stack[1:len(stack)]
		}
		if len(stack) > 0 {
			result[i%len(nums)] = nums[stack[0]]
		} else {
			result[i%len(nums)] = -1
		}
		stack = append([]int{i % len(nums)}, stack...)
	}
	return result
}
