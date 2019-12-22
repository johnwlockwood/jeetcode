package longest

import "testing"

func TestLongestPalindrome(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{
			input: "babad",
			want:  "bab",
		},
		{
			input: "cbbd",
			want:  "bb",
		},
	}
	for _, tc := range tests {
		if got := longestPalindrome(tc.input); got != tc.want {
			t.Errorf("from %s, got %s, want %s", tc.input, got, tc.want)
		}
	}
}
