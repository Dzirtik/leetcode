package main

import (
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	var queue, result []int
	for i, v := range nums {
		queue = append(queue, v)
		if len(queue) > k {
			queue = queue[1:len(queue)]
		}
		for j := 1; j < len(queue); {
			if queue[j] > queue[0] {
				queue = queue[1:len(queue)]
				j = 1
			} else {
				j++
			}
		}
		if i < k-1 {
			continue
		}
		result = append(result, queue[0])
	}
	return result
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))

}
