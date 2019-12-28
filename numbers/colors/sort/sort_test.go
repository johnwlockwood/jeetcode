package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type test struct {
		input []int
		want  []int
	}

	tests := []test{
		{
			input: []int{2, 0, 2, 1, 1, 0},
			want:  []int{0, 0, 1, 1, 2, 2},
		},
		{
			input: []int{2, 0, 1},
			want:  []int{0, 1, 2},
		},
	}
	for _, tc := range tests {
		inputStr := fmt.Sprintf("%v", tc.input)
		if sortColors(tc.input); !reflect.DeepEqual(tc.input, tc.want) {
			t.Errorf("%v got %v, want %v\n", inputStr, tc.input, tc.want)
		}
	}
}

func TestSortPointers(t *testing.T) {
	type test struct {
		input []int
		want  []int
	}

	tests := []test{
		{
			input: []int{2, 0, 2, 1, 1, 0},
			want:  []int{0, 0, 1, 1, 2, 2},
		},
		{
			input: []int{2, 0, 1},
			want:  []int{0, 1, 2},
		},
	}
	for _, tc := range tests {
		inputStr := fmt.Sprintf("%v", tc.input)
		if sortColorsPointers(tc.input); !reflect.DeepEqual(tc.input, tc.want) {
			t.Errorf("%v got %v, want %v\n", inputStr, tc.input, tc.want)
		}
	}
}
