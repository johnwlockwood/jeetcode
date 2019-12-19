package anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	type test struct {
		inputS string
		inputT string
		want   bool
	}
	tests := []test{
		{
			inputS: "caat",
			inputT: "raat",
			want:   false,
		},
		{
			inputS: "rat",
			inputT: "cat",
			want:   false,
		},
		{
			inputS: "cat",
			inputT: "tac",
			want:   true,
		},
	}

	for _, tc := range tests {
		if got := isAnagram(tc.inputS, tc.inputT); got != tc.want {
			t.Errorf("is %s anagram of %s? got %v, want %v\n", tc.inputS, tc.inputT, got, tc.want)
		}
	}
}

func TestIsAnagramNaive(t *testing.T) {
	type test struct {
		inputS string
		inputT string
		want   bool
	}
	tests := []test{
		{
			inputS: "caat",
			inputT: "raat",
			want:   false,
		},
		{
			inputS: "rat",
			inputT: "cat",
			want:   false,
		},
		{
			inputS: "cat",
			inputT: "tac",
			want:   true,
		},
		{
			inputS: "ca🔛",
			inputT: "🔛ac",
			want:   true,
		},
	}

	for _, tc := range tests {
		if got := isAnagramNaive(tc.inputS, tc.inputT); got != tc.want {
			t.Errorf("is %s anagram of %s? got %v, want %v\n", tc.inputS, tc.inputT, got, tc.want)
		}
	}
}
