package main

import "fmt"

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDistance(arrays [][]int) int {
	var dist int
	min := 10000
	max := -10000

	for _, array := range arrays {
		dist = maximum(dist, maximum(max-array[0], array[len(array)-1]-min))
		min = minimum(min, array[0])
		max = maximum(max, array[len(array)-1])
	}
	return dist
}

func main() {
	values := [][]int{}

	// These are the first two rows.
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5}

	// Append each row to the two-dimensional slice.
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row1)

	fmt.Println(maxDistance(values))
}
