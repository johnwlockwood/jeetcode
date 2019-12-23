package phone

import (
	"reflect"
	"sort"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	type test struct {
		input string
		want  []string
	}
	tests := []test{
		{
			input: "23",
			want:  []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			input: "9",
			want:  []string{"w", "x", "y", "z"},
		},
	}
	for _, tc := range tests {
		if got := letterCombinations(tc.input); !sortNCompare(got, tc.want) {
			t.Errorf("digits %s, got %v, want %v", tc.input, got, tc.want)
		}
	}
}

func sortNCompare(left, right []string) bool {
	sort.Strings(left)
	sort.Strings(right)
	return reflect.DeepEqual(left, right)
}
