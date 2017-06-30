package minWindow

import (
	"fmt"
)

func minWindow(s string, t string) string {
	var result string
	var cnt, min int
	hash := make(map[rune]int)
	for _, c := range t {
		hash[c] = -1
	}

	for i, c := range s {
		if v, ok := hash[c]; ok {
			hash[c] = i

			if v < 0 {
				cnt++
			}
			fmt.Println(hash, min, v, cnt)
			if v == min || v < 0 {
				min = i
				for _, value := range hash {
					if value < min {
						min = value
					}
				}
				if cnt == len(t) {
					if len(result) == 0 || i+1-min < len(result) {
						result = s[min : i+1]
					}
				}
			}
		}
	}

	return result
}
