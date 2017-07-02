package findFrequentTreeSum

import (
	"reflect"
	"testing"
)

func Test_findFrequentTreeSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"many results", args{&TreeNode{5, &TreeNode{Val: 2}, &TreeNode{Val: -3}}}, []int{2, -3, 4}},
		{"one result", args{&TreeNode{5, &TreeNode{Val: 2}, &TreeNode{Val: -5}}}, []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFrequentTreeSum(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFrequentTreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
