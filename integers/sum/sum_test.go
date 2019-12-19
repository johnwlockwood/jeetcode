package sum

import "testing"

type test struct {
	inputA int
	inputB int
	want   int
}

func TestGetSum(t *testing.T) {
	tests := []test{
		{
			inputA: 1,
			inputB: 500,
			want:   501,
		},
		{
			inputA: 0,
			inputB: 0,
			want:   0,
		},
		{
			inputA: 43,
			inputB: 15,
			want:   58,
		},
	}

	for _, tc := range tests {
		if got := getSum(tc.inputA, tc.inputB); got != tc.want {
			t.Errorf("sum of %d and %d. got %d, want %d", tc.inputA, tc.inputB, got, tc.want)
		}
	}
}
