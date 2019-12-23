package longest

import "testing"

func TestFindCenter(t *testing.T) {

	type test struct {
		input      string
		inputIndex int
		want       int
	}

	tests := []test{
		{
			input:      "babad",
			inputIndex: 1,
			want:       1,
		},
		{
			input:      "cbbd",
			inputIndex: 1,
			want:       2,
		},
		{
			input:      "cbbd",
			inputIndex: 2,
			want:       2,
		},
		{
			input:      "cbbd",
			inputIndex: 3,
			want:       3,
		},
		{
			input:      "abccda",
			inputIndex: 2,
			want:       3,
		},
		{
			input:      "abccda",
			inputIndex: 3,
			want:       3,
		},
	}
	for _, tc := range tests {
		if got := findCenter(tc.input, tc.inputIndex); got != tc.want {
			t.Errorf("from %s at %d, center got %d, want %v", tc.input, tc.inputIndex, got, tc.want)
		}
	}
}

func TestConstructPalindrome(t *testing.T) {
	type want struct {
		first int
		last  int
	}

	type test struct {
		input      string
		inputIndex int
		want       want
	}

	tests := []test{
		{
			input:      "babad",
			inputIndex: 1,
			want: want{
				first: 1,
				last:  2,
			},
		},
		{
			input:      "cbbd",
			inputIndex: 1,
			want: want{
				first: 1,
				last:  2,
			},
		},
		{
			input:      "cbbd",
			inputIndex: 2,
			want: want{
				first: 2,
				last:  2,
			},
		},
		{
			input:      "cbbd",
			inputIndex: 3,
			want: want{
				first: 3,
				last:  3,
			},
		},
		{
			input:      "abccda",
			inputIndex: 2,
			want: want{
				first: 0,
				last:  5,
			},
		},
		{
			input:      "abccda",
			inputIndex: 3,
			want: want{
				first: 3,
				last:  3,
			},
		},
	}
	for _, tc := range tests {
		if first, last := constructPalindrome(tc.input, tc.inputIndex); first != tc.want.first || last != tc.want.last {
			t.Errorf("from %s at %d, made palindrome , first %d, last %d, want %v", tc.input, tc.inputIndex, first, last, tc.want)
			//string(tc.input[start:end])
		}
	}
}

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
