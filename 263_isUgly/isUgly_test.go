package isUgly

import "testing"

func Test_isUgly(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"ugly", args{6}, true},
		{"ugly power of 2", args{8}, true},
		{"not ugly", args{14}, false},
		{"prime", args{11}, false},
		{"one", args{1}, true},
		{"zero", args{0}, false},
		{"minus", args{-2147483648}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUgly(tt.args.num); got != tt.want {
				t.Errorf("isUgly() = %v, want %v", got, tt.want)
			}
		})
	}
}
