package multiplication

import (
	"fmt"
	"testing"
)

func TestMultiplication(t *testing.T) {
	type test struct {
		inputA string
		inputB string
		want   string
	}

	tests := []test{
		{
			inputA: "1",
			inputB: "2",
			want:   "2",
		},
		{
			inputA: "10",
			inputB: "20",
			want:   "200",
		},
		{
			inputA: "10",
			inputB: "55",
			want:   "550",
		},
		{
			inputA: "20",
			inputB: "15",
			want:   "300",
		},
		{
			inputA: "5555",
			inputB: "1000",
			want:   "5555000",
		},
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
		{
			inputA: "3141592653589793238462643383279502884197169399375105820974944592",
			inputB: "2718281828459045235360287471352662497757247093699959574966967627",
			want:   "8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184",
		},
	}

	for _, tc := range tests {
		fmt.Println("")
		if got := multiplyRecursive(tc.inputA, tc.inputB, 0); got != tc.want {
			t.Errorf("for %s x %s, got %s, want %s", tc.inputA, tc.inputB, got, tc.want)
		}
	}
}

func TestKaratsubaMultiplication(t *testing.T) {
	type test struct {
		inputA string
		inputB string
		want   string
	}

	tests := []test{
		{
			inputA: "1",
			inputB: "2",
			want:   "2",
		},
		{
			inputA: "10",
			inputB: "20",
			want:   "200",
		},
		{
			inputA: "10",
			inputB: "55",
			want:   "550",
		},
		{
			inputA: "20",
			inputB: "15",
			want:   "300",
		},
		{
			inputA: "5555",
			inputB: "1000",
			want:   "5555000",
		},
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
		fmt.Println("")
		if got := multiply(tc.inputA, tc.inputB, 0); got != tc.want {
			t.Errorf("for %s x %s, got %s, want %s", tc.inputA, tc.inputB, got, tc.want)
		}
	}
}

func TestAddition(t *testing.T) {
	type test struct {
		inputA string
		inputB string
		want   string
	}

	tests := []test{
		{
			inputA: "10",
			inputB: "-5",
			want:   "5",
		},
		{
			inputA: "1",
			inputB: "1",
			want:   "2",
		},
		{
			inputA: "10",
			inputB: "1",
			want:   "11",
		},
		{
			inputA: "5",
			inputB: "6",
			want:   "11",
		},
		{
			inputA: "26",
			inputB: "16",
			want:   "42",
		},
		{
			inputA: "26",
			inputB: "-16",
			want:   "10",
		},
		// {
		// 	inputA: "26",
		// 	inputB: "-36",
		// 	want:   "-10",
		// },
		// {
		// 	inputA: "-26",
		// 	inputB: "16",
		// 	want:   "-10",
		// },
		{
			inputA: "1000",
			inputB: "5555",
			want:   "6555",
		},
	}

	for _, tc := range tests {
		if got := add(tc.inputA, tc.inputB); got != tc.want {
			t.Errorf("for %s x %s, got %s, want %s", tc.inputA, tc.inputB, got, tc.want)
		}
	}
}
