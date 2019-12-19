package disappeared

import (
	"reflect"
	"testing"
)

func TestDisappearedNumbers(t *testing.T) {
	type test struct {
		input []int
		want  []int
	}
	tests := []test{
		{
			input: []int{3, 2, 3, 1},
			want:  []int{4},
		},
		{
			input: []int{4, 2, 1, 4},
			want:  []int{3},
		},
	}

	for _, tc := range tests {
		if got := findDisappearedNumbers(tc.input); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("finding disappeared in %v, got %v, want %v", tc.input, got, tc.want)
		}
	}
}
