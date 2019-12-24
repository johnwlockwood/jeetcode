package search

import "testing"

func TestSearch(t *testing.T) {
	type test struct {
		input  []int
		inputT int
		want   int
	}
	tests := []test{
		{
			input:  []int{4, 5, 6, 7, 0, 1, 2},
			inputT: 0,
			want:   4,
		},
		{
			input:  []int{4, 5, 6, 7, 0, 1, 2},
			inputT: 3,
			want:   -1,
		},
	}
	for _, tc := range tests {
		if got := search(tc.input, tc.inputT); got != tc.want {
			t.Errorf("searching %v for %d, got %d, want %d", tc.input, tc.inputT, got, tc.want)
		}
	}
}
