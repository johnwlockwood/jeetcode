package common

import "testing"

func TestFindLength(t *testing.T) {
	type test struct {
		name   string
		inputA []int
		inputB []int
		want   int
	}

	tests := []test{
		{
			name:   "AB",
			inputA: []int{1, 2, 3, 2, 1},
			inputB: []int{3, 2, 1, 4, 7},
			want:   3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := findLength(tc.inputA, tc.inputB); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
