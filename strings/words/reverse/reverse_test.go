package reverse

import "testing"

func TestReverseWords(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	tests := []test{
		{
			input: "hello world",
			want:  "olleh dlrow",
		},
		{
			input: "hello world die",
			want:  "olleh dlrow eid",
		},
		{
			input: "I love u",
			want:  "I evol u",
		},
	}

	for _, tc := range tests {
		if got := reverseWords(tc.input); got != tc.want {
			t.Errorf("reverse %s? got %v, want %v\n", tc.input, got, tc.want)
		}
	}
}
