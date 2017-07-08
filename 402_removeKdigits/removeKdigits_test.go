package removeKdigits

import "testing"

func Test_removeKdigits(t *testing.T) {
	type args struct {
		num string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"54321 3", args{"54321", 3}, "21"},
		{"1432219 3", args{"1432219", 3}, "1219"},
		{"10200 1", args{"10200", 1}, "200"},
		{"10 2", args{"10", 2}, "0"},
		{"10 1", args{"10", 1}, "0"},
		{"100 1", args{"100", 1}, "0"},
		{"9 1", args{"9", 1}, "0"},
		{"112 1", args{"112", 1}, "11"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeKdigits(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("removeKdigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
