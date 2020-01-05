package fraction

import "testing"

func TestFractionToDecimal(t *testing.T) {
	type test struct {
		numerator   int
		denominator int
		want        string
	}

	tests := []test{
		{
			numerator:   -50,
			denominator: 8,
			want:        "-6.25",
		},
		{
			numerator:   1,
			denominator: 2,
			want:        "0.5",
		},
		{
			numerator:   2,
			denominator: 1,
			want:        "2",
		},
		{
			numerator:   -2,
			denominator: 1,
			want:        "-2",
		},
		{
			numerator:   2,
			denominator: -1,
			want:        "-2",
		},
		{
			numerator:   -2,
			denominator: -1,
			want:        "2",
		},
		{
			numerator:   208,
			denominator: 104,
			want:        "2",
		},
		{
			numerator:   2,
			denominator: 3,
			want:        "0.(6)",
		},
		{
			numerator:   4,
			denominator: 333,
			want:        "0.(012)",
		},
	}

	for _, tc := range tests {
		if got := fractionToDecimal(tc.numerator, tc.denominator); got != tc.want {
			t.Errorf("for %d/%d got %v, want %v", tc.numerator, tc.denominator, got, tc.want)
		}
	}
}
