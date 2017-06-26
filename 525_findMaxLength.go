package main

import (
	"fmt"
)

func findMaxLength(nums []int) int {
	var maxLen, count int
	hash := make(map[int]int)
	hash[0] = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count = count + 1
		} else {
			count = count - 1
		}
		if val, ok := hash[count]; ok {
			len := i - val
			if len > maxLen {
				maxLen = len
			}
		} else {
			hash[count] = i
		}
	}
	return maxLen

}

func main() {
	fmt.Println(findMaxLength([]int{0, 0, 0, 1, 1, 0, 0, 0, 0}))
}
