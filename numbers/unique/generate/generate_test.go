package generate

import (
	"reflect"
	"testing"
)

func TestSumZero(t *testing.T) {
	type test struct {
		input int
		want  []int
	}

	tests := []test{
		{
			input: 4,
			want:  []int{-2, -1, 1, 2},
		},
		{
			input: 5,
			want:  []int{-2, -1, 0, 1, 2},
		},
	}

	for _, tc := range tests {
		if got := sumZero(tc.input); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}
