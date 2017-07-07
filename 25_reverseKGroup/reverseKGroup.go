package reverseKGroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func revert(node *ListNode, k int) (*ListNode, *ListNode, *ListNode, bool) {
	if k > 1 {
		if node.Next != nil {
			head, prev, last, finished := revert(node.Next, k-1)
			if finished {
				prev.Next = node
				node.Next = nil
			} else {
				head = node
			}
			return head, node, last, finished
		}
		return node, node, nil, false
	}
	return node, node, node.Next, true
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var first, mid, last, prev, result *ListNode
	var finished bool
	for {
		first, mid, last, finished = revert(head, k)
		if result == nil {
			result = first
		} else {
			if finished {
				prev.Next = first
			} else {
				prev.Next = head
			}

		}
		if last == nil {
			break
		}
		prev = mid
		head = last
	}

	return result
}
