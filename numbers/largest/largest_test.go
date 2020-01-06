package largest

import (
	"testing"
)

func TestLargestNumber(t *testing.T) {
	type test struct {
		input []int
		want  string
	}
	tests := []test{
		{
			input: []int{999999998, 999999997, 999999999},
			want:  "999999999999999998999999997",
		},
		{
			input: []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want:  "9876543210",
		},
		{
			input: []int{0, 0},
			want:  "0",
		},
		{
			input: []int{10, 2},
			want:  "210",
		},
		{
			input: []int{3, 30, 34, 5, 9},
			want:  "9534330",
		},
	}

	for _, tc := range tests {
		if got := largestNumber(tc.input); got != tc.want {
			t.Errorf("for input %v: got %s, want %s", tc.input, got, tc.want)
		}
	}
}

func TestLargestNumberB(t *testing.T) {
	type test struct {
		input []int
		want  string
	}
	tests := []test{
		{
			input: []int{999999998, 999999997, 999999999},
			want:  "999999999999999998999999997",
		},
		{
			input: []int{0, 0},
			want:  "0",
		},
		{
			input: []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want:  "9876543210",
		},
		{
			input: []int{10, 2},
			want:  "210",
		},
		{
			input: []int{3, 30, 34, 5, 9},
			want:  "9534330",
		},
	}

	for _, tc := range tests {
		if got := largestNumberB(tc.input); got != tc.want {
			t.Errorf("for input %v: got %s, want %s", tc.input, got, tc.want)
		}
	}
}

// BenchmarkLargestNumber/3,_30,_34,_5,_9-16         	 1616294	       697 ns/op	     144 B/op	      15 allocs/op
// BenchmarkLargestNumber/3,_30,_34,_5,_9,_0,_0,_999999997,_999999999-16         	  767232	      1640 ns/op	     512 B/op	      29 allocs/op
// BenchmarkLargestNumber/range_3_-_66-16                                        	   53311	     23418 ns/op	    4384 B/op	     408 allocs/op
// PASS
func BenchmarkLargestNumberA(b *testing.B) {
	type bench struct {
		name  string
		input []int
	}

	a := new([50]int)
	for i := 0; i < 50; i++ {
		a[i] = i + 30
	}

	benches := []bench{
		{
			name:  "3, 30, 34, 5, 9",
			input: []int{3, 30, 34, 5, 9},
		},
		{
			name:  "3, 30, 34, 5, 9, 0, 0, 999999997, 999999999",
			input: []int{3, 30, 34, 5, 9, 999999997, 999999999},
		},
		{
			name:  "range 3 - 66",
			input: a[:],
		},
	}
	for _, bc := range benches {
		b.Run(bc.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = largestNumber(bc.input)
			}
		})
	}
}

func BenchmarkLargestNumbeB(b *testing.B) {
	type bench struct {
		name  string
		input []int
	}

	a := new([50]int)
	for i := 0; i < 50; i++ {
		a[i] = i + 30
	}

	benches := []bench{
		{
			name:  "3, 30, 34, 5, 9",
			input: []int{3, 30, 34, 5, 9},
		},
		{
			name:  "3, 30, 34, 5, 9, 0, 0, 999999997, 999999999",
			input: []int{3, 30, 34, 5, 9, 999999997, 999999999},
		},
		{
			name:  "range 3 - 66",
			input: a[:],
		},
	}
	for _, bc := range benches {
		b.Run(bc.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = largestNumberB(bc.input)
			}
		})
	}
}
