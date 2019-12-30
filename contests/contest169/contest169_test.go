package contest169

import "testing"

func TestJumpGame(t *testing.T) {
	type test struct {
		name  string
		input []int
		start int
		want  bool
	}
	tests := []test{
		{
			name:  "arr = [4,2,3,0,3,1,2], start = 5",
			input: []int{4, 2, 3, 0, 3, 1, 2},
			start: 5,
			want:  true,
		},
		{
			name:  "arr = [4,2,3,0,3,1,2], start = 0",
			input: []int{4, 2, 3, 0, 3, 1, 2},
			start: 0,
			want:  true,
		},
		{
			name:  "arr = [3,0,2,1,2], start = 2",
			input: []int{3, 0, 2, 1, 2},
			start: 2,
			want:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := canReach(tc.input, tc.start); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestVerbalArithmeticPuzzle(t *testing.T) {
	type test struct {
		name        string
		inputWords  []string
		inputResult string
		want        bool
	}

	tests := []test{
		{
			name:        `words = ["SEND","MORE"], result = "MONEY"`,
			inputWords:  []string{"SEND", "MORE"},
			inputResult: "MONEY",
			want:        true,
		},
		{
			name:        `words = ["SIX","SEVEN","SEVEN"], result = "TWENTY"`,
			inputWords:  []string{"SIX", "SEVEN", "SEVEN"},
			inputResult: "TWENTY",
			want:        true,
		},
		{
			name:        `words = ["THIS","IS","TOO"], result = "FUNNY"`,
			inputWords:  []string{"THIS", "IS", "TOO"},
			inputResult: "FUNNY",
			want:        true,
		},
		{
			name:        `words = ["LEET","CODE"], result = "POINT"`,
			inputWords:  []string{"LEET", "CODE"},
			inputResult: "POINT",
			want:        false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := isSolvable(tc.inputWords, tc.inputResult); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
