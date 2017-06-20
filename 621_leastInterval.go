package main

import "fmt"

func leastInterval(tasks []byte, n int) int {
	hash := make(map[byte]int)
	for _, value := range tasks {
		hash[value]++
	}
	var max, sum, same int
	for _, value := range hash {
		sum += value
		if value > max {
			max = value
			same = 0
		} else if value == max {
			same++
		}
	}
	result := (max-1)*n + max + same
	if sum >= result {
		return sum
	}

	return result

}

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 50))
}
