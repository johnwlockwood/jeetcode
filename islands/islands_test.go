package islands

import "testing"

func TestNumIslands(t *testing.T) {
	type test struct {
		input [][]int
		want  int
	}

	tests := []test{
		{
			input: [][]int{
				{1, 1, 1, 1, 0},
				{1, 1, 0, 1, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			want: 1,
		},
	}

	for _, tc := range tests {
		if got := numIslands(tc.input); got != tc.want {
			t.Errorf("got %d, want %s", got, tc.want)
		}
	}
}
