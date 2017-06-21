package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for key, value := range nums {
		x := target - value
		if key2, ok := hash[x]; ok {
			return []int{key2, key}
		}
		hash[value] = key
	}
	return []int{0, 0}
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
