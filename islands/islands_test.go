package islands

import "testing"

func TestNumIslands(t *testing.T) {
	type test struct {
		input [][]byte
		want  int
	}

	tests := []test{
		{
			input: [][]byte{
				{'1'},
			},
			want: 1,
		},
		{
			input: [][]byte{
				{'0'},
			},
			want: 0,
		},
		{
			input: [][]byte{
				{},
			},
			want: 0,
		},
		{
			input: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			input: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'0', '0', '0', '0', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '1'},
				{'0', '0', '0', '0', '0'},
			},
			want: 4,
		},
	}

	for _, tc := range tests {
		if got := numIslands(tc.input); got != tc.want {
			t.Errorf("got %d, want %d", got, tc.want)
		}
	}
}
