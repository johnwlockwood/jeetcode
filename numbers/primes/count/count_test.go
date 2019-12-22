package count

import (
	"reflect"
	"testing"
)

func TestIsPrime(t *testing.T) {
	type test struct {
		input int
		want  bool
	}

	tests := []test{
		{
			input: 0,
			want:  false,
		},
		{
			input: 1,
			want:  false,
		},
		{
			input: 2,
			want:  true,
		},
		{
			input: 3,
			want:  true,
		},
		{
			input: 4,
			want:  false,
		},
		{
			input: 6,
			want:  false,
		},
		{
			input: 10,
			want:  false,
		},
		{
			input: 11,
			want:  true,
		},
		{
			input: 37,
			want:  true,
		},
	}
	for _, tc := range tests {
		if got := isPrime(tc.input); got != tc.want {
			t.Errorf("is %d prime. got %v, want %v\n", tc.input, got, tc.want)
		}
	}
}

func TestMakeSieve(t *testing.T) {
	type test struct {
		input int
		want  []int
	}
	tests := []test{
		{
			input: 1,
			want:  []int{},
		},
		{
			input: 2,
			want:  []int{},
		},
		{
			input: 3,
			want:  []int{2},
		},
		{
			input: 4,
			want:  []int{2, 3},
		},
		{
			input: 6,
			want:  []int{2, 3, 5},
		},
		{
			input: 10,
			want:  []int{2, 3, 5, 7},
		},
	}

	for _, tc := range tests {
		if got := makeSieve(tc.input); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("primes below %d. got %v, want %v\n", tc.input, got, tc.want)
		}
	}
}

func TestCountPrimes(t *testing.T) {
	type test struct {
		input int
		want  int
	}
	tests := []test{
		{
			input: 0,
			want:  0,
		},
		{
			input: 1,
			want:  0,
		},
		{
			input: 2,
			want:  0,
		},
		{
			input: 3,
			want:  1,
		},
		{
			input: 4,
			want:  2,
		},
		{
			input: 6,
			want:  3,
		},
		{
			input: 10,
			want:  4,
		},
		{
			input: 11,
			want:  4,
		},
		{
			input: 37,
			want:  11,
		},
		{
			input: 38,
			want:  12,
		},
	}
	for _, tc := range tests {
		if got := countPrimes(tc.input); got != tc.want {
			t.Errorf("primes below %d. got %v, want %v\n", tc.input, got, tc.want)
		}
	}
}
