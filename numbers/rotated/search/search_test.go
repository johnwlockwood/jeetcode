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
			name:  "size 0",
			input: []int{},
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
		name   string
		input  []int
		inputT int
		want   int
	}
	tests := []test{
		{
			name:   "[5,1,3] 3",
			input:  []int{5, 1, 3},
			inputT: 3,
			want:   2,
		},
		{
			name:   "0 in 1 pivot 0",
			input:  []int{1},
			inputT: 0,
			want:   -1,
		},
		{
			name:   "3 in 2 pivot 1",
			input:  []int{3, 1},
			inputT: 3,
			want:   0,
		},
		{
			name:   "4 in 7 pivot 4",
			input:  []int{4, 5, 6, 7, 0, 1, 2},
			inputT: 0,
			want:   4,
		},
		{
			name:   "3 not in 7 pivot 4",
			input:  []int{4, 5, 6, 7, 0, 1, 2},
			inputT: 3,
			want:   -1,
		},
		{
			name:   "3 not in 0, no pivot",
			input:  []int{},
			inputT: 3,
			want:   -1,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := search(tc.input, tc.inputT); got != tc.want {
				t.Errorf("searching %v for %d, got %d, want %d", tc.input, tc.inputT, got, tc.want)
			}
		})

	}
}
