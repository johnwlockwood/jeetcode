package longest

import "testing"

func TestConstructPalindrome(t *testing.T) {
	type want struct {
		start int
		end   int
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
				start: 0,
				end:   3,
			},
		},
		{
			input:      "cbbd",
			inputIndex: 1,
			want: want{
				start: 1,
				end:   3,
			},
		},
		{
			input:      "cbbd",
			inputIndex: 2,
			want: want{
				start: 2,
				end:   3,
			},
		},
		{
			input:      "cbbd",
			inputIndex: 3,
			want: want{
				start: 3,
				end:   4,
			},
		},
		{
			input:      "abccda",
			inputIndex: 2,
			want: want{
				start: 0,
				end:   6,
			},
		},
		{
			input:      "abccda",
			inputIndex: 3,
			want: want{
				start: 3,
				end:   4,
			},
		},
	}
	for _, tc := range tests {
		if start, end := constructPalindrome(tc.input, tc.inputIndex); start != tc.want.start || end != tc.want.end {
			t.Errorf("from %s at %d, made palindrome , start %d, end %d, want %v", tc.input, tc.inputIndex, start, end, tc.want)
			//string(tc.input[start:end])
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	type want struct {
		is    bool
		start int
		end   int
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
				is:    true,
				start: 0,
				end:   3,
			},
		},
		{
			input:      "cbbd",
			inputIndex: 1,
			want: want{
				is:    true,
				start: 1,
				end:   3,
			},
		},
		{
			input:      "cbbd",
			inputIndex: 2,
			want: want{
				is:    true,
				start: 1,
				end:   3,
			},
		},
		{
			input:      "cbbd",
			inputIndex: 3,
			want: want{
				is:    false,
				start: 3,
				end:   4,
			},
		},
	}
	for _, tc := range tests {
		if got, start, end := isPalindrome(tc.input, tc.inputIndex); got != tc.want.is || start != tc.want.start || end != tc.want.end {
			t.Errorf("from %s at %d, got %v start %d, end %d, want %v", tc.input, tc.inputIndex, got, start, end, tc.want)
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
