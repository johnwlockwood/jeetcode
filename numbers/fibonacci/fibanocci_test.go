package fibonacci

import "testing"

func TestFibonacciRecursiveMem(t *testing.T) {
	type test struct {
		input int
		want  int
	}

	tests := []test{
		{
			input: 15,
			want:  610,
		},
		{
			input: 0,
			want:  0,
		},
		{
			input: 1,
			want:  1,
		},
		{
			input: 2,
			want:  1,
		},
		{
			input: 3,
			want:  2,
		},
		{
			input: 4,
			want:  3,
		},
		{
			input: 5,
			want:  5,
		},
		{
			input: 6,
			want:  8,
		},
		{
			input: 7,
			want:  13,
		},
	}

	t.Run("fibRecursiveMem", func(t *testing.T) {
		for _, tc := range tests {
			if got := fibRecursiveMem(tc.input); got != tc.want {
				t.Errorf("for %d got %d, want %d", tc.input, got, tc.want)
			}
		}
	})

	t.Run("fibBottomUp", func(t *testing.T) {
		for _, tc := range tests {
			if got := fibBottomUp(tc.input); got != tc.want {
				t.Errorf("for %d got %d, want %d", tc.input, got, tc.want)
			}
		}
	})
}
