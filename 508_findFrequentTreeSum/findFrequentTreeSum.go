package findFrequentTreeSum

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Variables struct {
	result []int
	count  map[int]int
	max    int
}

func findFrequentTreeSum(root *TreeNode) []int {
	variables := &Variables{count: make(map[int]int)}
	sum(root, variables)

	return variables.result

}

func sum(root *TreeNode, variables *Variables) int {
	if root == nil {
		return 0
	}
	fmt.Println(root.Val)
	summa := root.Val + sum(root.Left, variables) + sum(root.Right, variables)

	variables.count[summa] = variables.count[summa] + 1
	if variables.count[summa] == variables.max {
		variables.result = append(variables.result, summa)
	} else if variables.count[summa] > variables.max {
		variables.max = variables.count[summa]
		variables.result = []int{summa}
	}

	return summa
}
