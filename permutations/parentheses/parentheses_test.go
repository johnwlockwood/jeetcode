package parentheses

import (
	"reflect"
	"sort"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	type test struct {
		input int
		want  []string
	}

	tests := []test{
		{
			input: 3,
			want: []string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
		{
			input: 2,
			want: []string{
				"(())",
				"()()",
			},
		},
		{
			input: 0,
			want:  []string{},
		},
	}

	for _, tc := range tests {
		if got := generateParenthesis(tc.input); !sortNCompare(got, tc.want) {
			t.Errorf("for %d, got %v, want %v", tc.input, got, tc.want)
		}
	}
}

func sortNCompare(left, right []string) bool {
	sort.Strings(left)
	sort.Strings(right)
	return reflect.DeepEqual(left, right)
}
