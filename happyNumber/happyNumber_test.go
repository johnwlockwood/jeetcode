package happyNumber

import (
	"testing"
)

func TestGetNext(t *testing.T) {
	type test struct {
		input int
		want  int
	}
	tests := []test{
		{input: 19, want: 82},
		{input: 0, want: 0},
		{input: 7, want: 49},
		{input: 9, want: 81},
	}
	for _, tc := range tests {
		if got := getNext(tc.input); got != tc.want {
			t.Errorf("got %v, want %v\n", got, tc.want)
		}
	}
}

func TestHappyNumber(t *testing.T) {
	type happyNumberType = func(n int) bool
	type approach struct {
		f    happyNumberType
		name string
	}
	approaches := []approach{
		{name: "IsHappyApproach1", f: IsHappyApproach1},
		{name: "IsHappyNaive", f: IsHappyNaive},
	}
	type test struct {
		input int
		want  bool
	}
	tests := []test{
		{input: 19, want: true},
		{input: 0, want: false},
		{input: 7, want: true},
		{input: 9, want: false},
	}
	for _, a := range approaches {
		t.Run(a.name, func(t *testing.T) {
			for _, tc := range tests {
				if got := a.f(tc.input); got != tc.want {
					t.Errorf("got %v, want %v\n", got, tc.want)
				}
			}
		})
	}
}
