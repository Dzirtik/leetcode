package majorityElement

import "testing"

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{[]int{1, 2, 1, 2, 2}}, 2},
		{"test2", args{[]int{2, 2, 1, 1, 1}}, 1},
	}
	for _, tt := range tests {
		if got := majorityElement(tt.args.nums); got != tt.want {
			t.Errorf("%q. majorityElement() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
