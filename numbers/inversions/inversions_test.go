package inversions

import "testing"

func TestCountInversions(t *testing.T) {
	type test struct {
		name  string
		input []int
		want  int
	}

	tests := []test{
		{
			name:  "even five inversions with left and right in order",
			input: []int{1, 5, 7, 2, 3, 6},
			want:  5,
		},
		{
			name:  "even two inversions",
			input: []int{1, 2, 4, 5, 3, 6},
			want:  2,
		},
		{
			name:  "even no inversions",
			input: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want:  0,
		},
		{
			name:  "odd no inversions",
			input: []int{1, 2, 3, 4, 5},
			want:  0,
		},
		{
			name:  "even one inversions",
			input: []int{1, 2, 4, 3, 5, 6},
			want:  1,
		},
		{
			name:  "odd one inversions",
			input: []int{1, 2, 3, 4, 6, 5, 7},
			want:  1,
		},
		{
			name:  "even 6 inversions",
			input: []int{2, 1, 6, 5, 3, 4},
			want:  6,
		},
		{
			name:  "TEST CASE - 1",
			input: []int{1, 3, 5, 2, 4, 6},
			want:  3,
		},
		{
			name:  "TEST CASE - 2",
			input: []int{1, 5, 3, 2, 4},
			want:  4,
		},
		{
			name:  "TEST CASE - 3",
			input: []int{5, 4, 3, 2, 1},
			want:  10,
		},
		{
			name:  "TEST CASE - 4",
			input: []int{1, 6, 3, 2, 4, 5},
			want:  5,
		},
		{
			name:  "Test Case - #1 - 15 numbers",
			input: []int{9, 12, 3, 1, 6, 8, 2, 5, 14, 13, 11, 7, 10, 4, 0},
			want:  56,
		},
		{
			name:  "Test Case - #2 - 50 numbers",
			input: []int{37, 7, 2, 14, 35, 47, 10, 24, 44, 17, 34, 11, 16, 48, 1, 39, 6, 33, 43, 26, 40, 4, 28, 5, 38, 41, 42, 12, 13, 21, 29, 18, 3, 19, 0, 32, 46, 27, 31, 25, 15, 36, 20, 8, 9, 49, 22, 23, 30, 45},
			want:  590,
		},
		{
			name:  "Test Case - #3 - 100 numbers",
			input: []int{4, 80, 70, 23, 9, 60, 68, 27, 66, 78, 12, 40, 52, 53, 44, 8, 49, 28, 18, 46, 21, 39, 51, 7, 87, 99, 69, 62, 84, 6, 79, 67, 14, 98, 83, 0, 96, 5, 82, 10, 26, 48, 3, 2, 15, 92, 11, 55, 63, 97, 43, 45, 81, 42, 95, 20, 25, 74, 24, 72, 91, 35, 86, 19, 75, 58, 71, 47, 76, 59, 64, 93, 17, 50, 56, 94, 90, 89, 32, 37, 34, 65, 1, 73, 41, 36, 57, 77, 30, 22, 13, 29, 38, 16, 88, 61, 31, 85, 33, 54},
			want:  2372,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountInversions(tc.input); got != tc.want {
				t.Errorf("for %v, got %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}

func BenchmarkCountInversions(b *testing.B) {

}
