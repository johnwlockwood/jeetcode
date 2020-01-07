package multiplication

import "testing"

func TestMultiplication(t *testing.T) {
	type test struct {
		inputA string
		inputB string
		want   string
	}

	tests := []test{
		{
			inputA: "1000",
			inputB: "5555",
			want:   "5555000",
		},
		{
			inputA: "5678",
			inputB: "1234",
			want:   "7006652",
		},
	}

	for _, tc := range tests {
		if got := multiply(tc.inputA, tc.inputA); got != tc.want {
			t.Errorf("for %s x %s, got %s, want %s", tc.inputA, tc.inputB, got, tc.want)
		}
	}
}
