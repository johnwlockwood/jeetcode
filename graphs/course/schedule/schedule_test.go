package schedule

import "testing"

func TestCanFinish(t *testing.T) {
	type test struct {
		inputN       int
		inputPrereqs [][]int
		want         bool
	}

	tests := []test{
		{
			inputN:       2,
			inputPrereqs: [][]int{{1, 0}},
			want:         true,
		},
		{
			inputN:       2,
			inputPrereqs: [][]int{{1, 0}, {0, 1}},
			want:         false,
		},
	}
	for _, tc := range tests {
		if got := canFinish(tc.inputN, tc.inputPrereqs); got != tc.want {
			t.Errorf("for n: %d, classes: %v: got %v, want %v", tc.inputN, tc.inputPrereqs, got, tc.want)
		}
	}
}
