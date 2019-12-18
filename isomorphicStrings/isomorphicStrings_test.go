package isomorphicStrings

import (
	"testing"
)

func TestIsomorphicStrings(t *testing.T) {
	type test struct {
		inputS string
		inputT string
		want   bool
	}
	tests := []test{
		{
			inputS: "egg",
			inputT: "add",
			want:   true,
		},
		{
			inputS: "peggy",
			inputT: "yeggy",
			want:   false,
		},
		{
			inputS: "yeggy",
			inputT: "peggy",
			want:   false,
		},
		{
			inputS: "paper",
			inputT: "title",
			want:   true,
		},
	}

	for _, tc := range tests {
		if got := isIsomorphic(tc.inputS, tc.inputT); got != tc.want {
			t.Errorf("got %v, want %v for s: %s t: %s\n", got, tc.want, tc.inputS, tc.inputT)
		}
	}
}
