package rotate

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	type test struct {
		name  string
		input [][]int
		want  [][]int
	}

	tests := []test{
		{
			name:  "3n",
			input: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:  [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			name:  "2n",
			input: [][]int{{1, 2}, {3, 4}},
			want:  [][]int{{3, 1}, {4, 2}},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			inputStr := fmt.Sprintf("%v", tc.input)
			if rotate(tc.input); !reflect.DeepEqual(tc.input, tc.want) {
				t.Errorf("rotate %s, got %v, want %v", inputStr, tc.input, tc.want)
			} else {
				fmt.Printf("%v rotated to %v, expected %v\n", inputStr, tc.input, tc.want)
			}
		})

	}
}
