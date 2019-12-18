package diagonal

import (
	"reflect"
	"testing"
)

func TestGetDiagonal(t *testing.T) {
	type test struct {
		input [][]int
		want  [][]int
	}
	tests := []test{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{1},
				{2, 4},
				{3, 5, 7},
				{6, 8},
				{9},
			},
		},
		{
			input: [][]int{
				{1, 2, 3},
				{4, 55, 6},
				{7, 8, 9},
				{10, 11, 12},
			},
			want: [][]int{
				{1},
				{2, 4},
				{3, 55, 7},
				{6, 8, 10},
				{9, 11},
				{12},
			},
		},
	}

	for _, tc := range tests {
		if got, _ := getDiagonals(tc.input); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got %v, want %v\n", got, tc.want)
		}
	}
}
