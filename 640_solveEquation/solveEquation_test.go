package solveEquation

import "testing"

func Test_solveEquation(t *testing.T) {
	type args struct {
		equation string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"x+5-3+x=6+x-2"}, "x=2"},
		{"test1", args{"x=x"}, "Infinite solutions"},
		{"test1", args{"2x=x"}, "x=0"},
		{"test1", args{"2x+3x-6x=x+2"}, "x=-1"},
		{"test1", args{"x=x+2"}, "No solution"},
		{"test1", args{"0x=0"}, "Infinite solutions"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveEquation(tt.args.equation); got != tt.want {
				t.Errorf("solveEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
