package merge

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type test struct {
		input [][]int
		want  [][]int
	}

	tests := []test{
		{
			input: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			input: [][]int{{1, 4}, {4, 5}},
			want:  [][]int{{1, 5}},
		},
		{
			input: [][]int{{1, 4}, {0, 4}},
			want:  [][]int{{0, 4}},
		},
	}
	for _, tc := range tests {
		if got := merge(tc.input); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%v got %v, want %v\n", tc.input, got, tc.want)
		}
	}
}
