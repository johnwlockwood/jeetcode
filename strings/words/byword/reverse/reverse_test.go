package reverse

import "testing"

func TestReverse(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{
			input: "the sky is blue",
			want:  "blue is sky the",
		},
		{
			input: "  hello world!  ",
			want:  "world! hello",
		},
		{
			input: "a good   example",
			want:  "example good a",
		},
	}

	for _, tc := range tests {
		if got := reverseWords(tc.input); got != tc.want {
			t.Errorf("%s reversed: got %s, want %s", tc.input, got, tc.want)
		}
	}
}
