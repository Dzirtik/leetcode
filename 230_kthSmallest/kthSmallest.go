package kthSmallest

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var count []int
	count = helper(root, count)
	return count[k-1]

}
func helper(node *TreeNode, count []int) []int {
	if node == nil {
		return count
	}
	count = helper(node.Left, count)
	count = append(count, node.Val)
	count = helper(node.Right, count)

	return count
}
