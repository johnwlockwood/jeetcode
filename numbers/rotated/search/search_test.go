package search

import "testing"

func TestFindRotation(t *testing.T) {
	type test struct {
		name  string
		input []int
		want  int
	}
	tests := []test{
		{
			name:  "pivot in middle of even",
			input: []int{5, 6, 7, 8, 0, 1, 2, 4},
			want:  4,
		},
		{
			name:  "pivot right of odd",
			input: []int{3, 4, 5, 6, 0, 1, 2},
			want:  4,
		},
		{
			name:  "pivot right of even",
			input: []int{4, 5, 6, 7, 8, 0, 1, 2},
			want:  5,
		},
		{
			name:  "size 1",
			input: []int{4},
			want:  0,
		},
		{
			name:  "pivot at zero of odd",
			input: []int{0, 1, 2},
			want:  0,
		},
		{
			name:  "pivot at zero of even",
			input: []int{0, 1, 2, 4},
			want:  0,
		},
		{
			name:  "pivot left of even",
			input: []int{4, 0, 1, 2},
			want:  1,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := findRotation(tc.input, 0, len(tc.input)-1); got != tc.want {
				t.Errorf("finding pivot of %v, got %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}

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
