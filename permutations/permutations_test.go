package permutations

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	type funcType = func(nums []int) [][]int
	type approach struct {
		f    funcType
		name string
	}
	approaches := []approach{
		{name: "permute", f: permute},
		{name: "permute2", f: permute2},
	}
	type test struct {
		input []int
		want  [][]int
	}
	tests := []test{
		{
			input: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3},
				{2, 1, 3},
				{3, 1, 2},
				{1, 3, 2},
				{2, 3, 1},
				{3, 2, 1},
			},
		},
	}
	for _, a := range approaches {

		t.Run(a.name, func(t *testing.T) {
			for _, tc := range tests {
				cleanInput := make([]int, len(tc.input))
				copy(cleanInput, tc.input)
				if got := a.f(cleanInput); !reflect.DeepEqual(got, tc.want) {
					t.Errorf("got %v, want %v\n", got, tc.want)
				}
			}
		})
	}
}
