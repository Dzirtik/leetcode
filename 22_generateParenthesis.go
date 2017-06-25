package main

import (
	"fmt"
	"strings"
)

type hash struct {
	str string
	cnt int
}

func generateParenthesis(n int) []string {
	var graph []hash
	var result []string

	var i hash
	for {
		fmt.Println(i)
		if len(i.str)+i.cnt == n*2 {
			result = append(result, i.str+strings.Repeat(")", i.cnt))
			fmt.Println("result - ", i)
		} else if len(i.str)+i.cnt < n*2 {
			if i.cnt > 0 {
				graph = append(graph, hash{i.str + ")", i.cnt - 1})
				fmt.Println("add - ", graph[len(graph)-1])
			}
			i = hash{i.str + "(", i.cnt + 1}
			continue
		}
		if len(graph) > 0 {
			i, graph = graph[0], graph[1:]
		} else {
			break
		}
	}

	return result

}

func main() {
	fmt.Println(generateParenthesis(3))
}
