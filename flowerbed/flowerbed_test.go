package flowerbed

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	type test struct {
		flowerbed []int
		flowers   int
		want      bool
	}
	tests := []test{
		{
			flowerbed: []int{0},
			flowers:   1,
			want:      true,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			flowers:   1,
			want:      true,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			flowers:   2,
			want:      false,
		},
		{
			flowerbed: []int{0, 0, 0, 0, 1},
			flowers:   2,
			want:      true,
		},
		{
			flowerbed: []int{0, 0, 1, 0, 0},
			flowers:   2,
			want:      true,
		},
		{
			flowerbed: []int{0, 0, 0, 0, 0, 0, 1},
			flowers:   3,
			want:      true,
		},
		{
			flowerbed: []int{0, 1, 0, 0, 0, 0, 1},
			flowers:   3,
			want:      false,
		},
		{
			flowerbed: []int{0, 1, 0, 1, 0, 1, 0},
			flowers:   1,
			want:      false,
		},
		{
			flowerbed: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			flowers:   1,
			want:      true,
		},
		{
			flowerbed: []int{0, 0, 0, 0, 0, 0, 0, 1, 0},
			flowers:   2,
			want:      true,
		},
		{
			flowerbed: []int{0, 1, 0, 1, 0, 1, 0, 0},
			flowers:   1,
			want:      true,
		},
		{
			flowerbed: []int{0, 0, 1, 0, 0, 1, 0, 0},
			flowers:   2,
			want:      true,
		},
	}

	for _, tc := range tests {
		if got := canPlaceFlowers(tc.flowerbed, tc.flowers); got != tc.want {
			t.Errorf("placing %d flowers in %v bed: got %v, want %v", tc.flowers, tc.flowerbed, got, tc.want)
		}
	}
}
