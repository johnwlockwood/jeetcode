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
			inputN:       0,
			inputPrereqs: [][]int{},
			want:         true,
		},
		{
			inputN:       1,
			inputPrereqs: [][]int{},
			want:         true,
		},
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
		{
			inputN:       4,
			inputPrereqs: [][]int{{1, 0}, {2, 0}, {0, 3}, {0, 4}},
			want:         true,
		},
	}
	for _, tc := range tests {
		if got := canFinish(tc.inputN, tc.inputPrereqs); got != tc.want {
			t.Errorf("for n: %d, classes: %v: got %v, want %v", tc.inputN, tc.inputPrereqs, got, tc.want)
		}
	}
}
