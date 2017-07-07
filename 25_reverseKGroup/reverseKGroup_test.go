package reverseKGroup

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"second", args{makeList([]int{1, 2, 3, 4, 5}), 2}, makeList([]int{2, 1, 4, 3, 5})},
		{"third", args{makeList([]int{1, 2, 3, 4, 5}), 3}, makeList([]int{3, 2, 1, 4, 5})},
		{"empty", args{makeList([]int{}), 1}, makeList([]int{})},
		{"one", args{makeList([]int{1}), 1}, makeList([]int{1})},
		{"one-two", args{makeList([]int{1, 2}), 2}, makeList([]int{2, 1})},
		{"one-two three four", args{makeList([]int{1, 2, 3, 4}), 2}, makeList([]int{2, 1, 4, 3})},
		{"one-two and three", args{makeList([]int{1, 2}), 3}, makeList([]int{1, 2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makeList(array []int) *ListNode {
	var node *ListNode
	for i := len(array) - 1; i >= 0; i-- {
		node = &ListNode{array[i], node}
	}
	return node
}
