package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (self *ListNode) String() string {
	str := fmt.Sprintf("%d", self.Val)
	for self = self.Next; self != nil; self = self.Next {
		str = str + fmt.Sprintf(" -> %d", self.Val)
	}
	return str
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var transfer int
	var l, current, prev *ListNode
	for {
		if l1 != nil || l2 != nil || transfer > 0 {
			val := 0
			if l1 != nil {
				val += l1.Val
			}
			if l2 != nil {
				val += l2.Val
			}
			val += transfer
			if val > 9 {
				transfer = 1
				val = val % 10
			} else {
				transfer = 0
			}
			current = &ListNode{Val: val, Next: nil}
			if l == nil {
				l = current
				prev = current
			} else {
				prev.Next = current
				prev = current
			}

			if l1 != nil {
				l1 = l1.Next
			}
			if l2 != nil {
				l2 = l2.Next
			}
		} else {
			return l
		}
	}
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 8, Next: nil}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
	fmt.Println(addTwoNumbers(l1, l2))
}
