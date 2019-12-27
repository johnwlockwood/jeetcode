package group

import (
	"reflect"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	type test struct {
		input []string
		want  [][]string
	}

	tests := []test{
		{
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"ate", "eat", "tea"},
				{"nat", "tan"},
				{"bat"},
			},
		},
	}

	for _, tc := range tests {
		if got := groupAnagrams(tc.input); reflect.DeepEqual(got, tc.want) {
			t.Errorf("%v got %v, want %v\n", tc.input, got, tc.want)
		}
	}
}
