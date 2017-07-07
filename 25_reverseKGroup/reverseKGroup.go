package reverseKGroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	curr := head
	var count int
	for curr != nil && count != k {
		curr = curr.Next
		count++
	}
	if count == k {
		curr = reverseKGroup(curr, k)
		for count > 0 {
			count--
			tmp := head.Next
			head.Next = curr
			curr = head
			head = tmp
		}
		head = curr
	}
	return head
}
