package nondecreasing

import "testing"

func TestCheckPossibility(t *testing.T) {
	type test struct {
		input []int
		want  bool
	}

	tests := []test{
		{
			input: []int{4, 2, 3},
			want:  true,
		},
		{
			input: []int{4, 2, 1},
			want:  false,
		},
		{
			input: []int{3, 4, 2, 3},
			want:  false,
		},
	}

	for _, tc := range tests {
		if got := checkPossibility(tc.input); got != tc.want {
			t.Errorf("for %v: got %v, want %v", tc.input, got, tc.want)
		}
	}
}
