package mostwater

import "testing"

func TestMostWater(t *testing.T) {
	type test struct {
		input []int
		want  int
	}
	tests := []test{
		{
			input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:  49,
		},
	}
	for _, tc := range tests {
		if got := maxArea(tc.input); got != tc.want {
			t.Errorf("containers %v, got %v, want %v", tc.input, got, tc.want)
		}
	}
}
