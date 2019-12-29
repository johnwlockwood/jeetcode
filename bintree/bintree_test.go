package bintree

import (
	"reflect"
	"testing"
)

func TestGetAll(t *testing.T) {
	type test struct {
		name   string
		inputS *TreeNode
		inputT *TreeNode
		want   []int
	}

	tests := []test{
		{
			name: "[4 1][4 1]",
			inputS: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
			},
			inputT: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
			},
			want: []int{1, 1, 2, 3, 4, 4, 7, 8},
		},
		{
			name: "[4 3 null 1][3 null 1]",
			inputS: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			inputT: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 1,
				},
			},
			want: []int{1, 1, 3, 3, 4},
		},
		{
			name: `[3,4,5,1,2,null,null,0][4,1,2]`,
			inputS: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 0,
						},
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			inputT: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			want: []int{0, 1, 1, 2, 2, 3, 4, 4, 5},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := getAllElements(tc.inputS, tc.inputT); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
