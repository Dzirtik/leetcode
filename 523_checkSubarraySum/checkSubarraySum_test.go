package checkSubarraySum

import "testing"

func Test_checkSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"42 AND 4-6", args{[]int{23, 2, 4, 6, 7}, 6}, true},
		{"2-4", args{[]int{22, 2, 4, 6, 7}, 6}, true},
		{"42", args{[]int{23, 2, 5, 5, 7}, 6}, true},
		{"false", args{[]int{23, 2}, 6}, false},
		{"0, 0", args{[]int{0, 0}, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("checkSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
