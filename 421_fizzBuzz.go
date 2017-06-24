package main

import (
	"fmt"
)

func fizzBuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		var str string
		if i%3 == 0 {
			str = "Fiz"
		}
		if i%5 == 0 {
			str += "Buzz"
		}
		if str != "" {
			result = append(result, str)
		} else {
			result = append(result, fmt.Sprintf("%v", i))
		}
	}

	return result
}

func main() {
	fmt.Println(fizzBuzz(100))
}
