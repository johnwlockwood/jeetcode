package permutations

import (
	"reflect"
	"sort"
	"testing"
)

type Perms [][3]int

func (m Perms) Len() int { return len(m) }
func (m Perms) Less(i, j int) bool {
	for x := range m[i] {
		if m[i][x] == m[j][x] {
			continue
		}
		return m[i][x] < m[j][x]
	}
	return false
}
func (m Perms) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func TestPermutations(t *testing.T) {
	type funcType = func(nums []int) [][]int
	type approach struct {
		f    funcType
		name string
	}
	approaches := []approach{
		{name: "permute", f: permute},
		{name: "permute2", f: permute2},
		{name: "permuteRecursive", f: permuteRecursive},
		{name: "permuteBacktrack", f: permuteBacktrack},
		{name: "permuteBacktrack2", f: permuteBacktrack2},
	}
	type test struct {
		input []int
		want  [][]int
	}
	tests := []test{
		{
			input: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3},
				{2, 1, 3},
				{3, 1, 2},
				{1, 3, 2},
				{2, 3, 1},
				{3, 2, 1},
			},
		},
		{
			input: []int{5, 6},
			want:  permute2([]int{5, 6}),
		},
		{
			input: []int{9, 8, 7, 6},
			want:  permute2([]int{9, 8, 7, 6}),
		},
		{
			input: []int{11, 45, 6, 7, 88},
			want:  permute2([]int{11, 45, 6, 7, 88}),
		},
	}
	for _, a := range approaches {

		t.Run(a.name, func(t *testing.T) {
			for _, tc := range tests {
				cleanInput := make([]int, len(tc.input))
				copy(cleanInput, tc.input)
				if got := a.f(cleanInput); !sortNCompare(got, tc.want) {
					t.Errorf("got %v, want %v\n", got, tc.want)
				}
			}
		})
	}
}

func sortNCompare(got, want [][]int) bool {
	sort.Slice(got, func(i, j int) bool {
		for x := range got[i] {
			if got[i][x] == got[j][x] {
				continue
			}
			return got[i][x] < got[j][x]
		}
		return false
	})
	sort.Slice(want, func(i, j int) bool {
		for x := range want[i] {
			if want[i][x] == want[j][x] {
				continue
			}
			return want[i][x] < want[j][x]
		}
		return false
	})
	return reflect.DeepEqual(got, want)
}
