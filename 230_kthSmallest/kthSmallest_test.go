package kthSmallest

import "testing"

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"one element", args{&TreeNode{5, nil, nil}, 1}, 5},
		{"three elements", args{&TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, nil}, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
