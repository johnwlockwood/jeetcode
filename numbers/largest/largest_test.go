package largest

import (
	"testing"
)

func TestLargestNumber(t *testing.T) {
	type test struct {
		input []int
		want  string
	}
	tests := []test{
		{
			input: []int{999999998, 999999997, 999999999},
			want:  "999999999999999998999999997",
		},
		{
			input: []int{0, 0},
			want:  "0",
		},
		{
			input: []int{10, 2},
			want:  "210",
		},
		{
			input: []int{3, 30, 34, 5, 9},
			want:  "9534330",
		},
	}

	for _, tc := range tests {
		if got := largestNumber(tc.input); got != tc.want {
			t.Errorf("for input %v: got %s, want %s", tc.input, got, tc.want)
		}
	}
}
