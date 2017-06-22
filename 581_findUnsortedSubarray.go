package main

import "fmt"
import "sort"

func findUnsortedSubarray(nums []int) int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)

	sort.Ints(tmp)
	fmt.Println(tmp)
	start, finish := -1, -1
	for i, value := range tmp {
		if value != nums[i] {
			start = i
			break
		}
	}

	if start == -1 {
		return 0
	}

	for i := len(tmp) - 1; i >= 0; i-- {
		if tmp[i] != nums[i] {
			finish = i
			break
		}
	}

	return finish - start + 1
}

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 1}))
}
