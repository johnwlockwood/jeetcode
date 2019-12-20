package subtree

import (
	"testing"

	"github.com/johnwlockwood/jeetcode/bintree"
)

func TestIsSubTree(t *testing.T) {
	type test struct {
		name   string
		inputS *bintree.TreeNode
		inputT *bintree.TreeNode
		want   bool
	}
	tests := []test{
		{
			name: "[1 1][1]",
			inputS: &bintree.TreeNode{
				Val: 1,
				Left: &bintree.TreeNode{
					Val: 1,
				},
			},
			inputT: &bintree.TreeNode{
				Val: 1,
			},
			want: true,
		},
		{
			name: "[4 1][4 1]",
			inputS: &bintree.TreeNode{
				Val: 4,
				Left: &bintree.TreeNode{
					Val: 1,
				},
			},
			inputT: &bintree.TreeNode{
				Val: 4,
				Left: &bintree.TreeNode{
					Val: 1,
				},
			},
			want: true,
		},
		{
			name: "[4 3 null 1][3 null 1]",
			inputS: &bintree.TreeNode{
				Val: 4,
				Left: &bintree.TreeNode{
					Val: 3,
					Right: &bintree.TreeNode{
						Val: 1,
					},
				},
			},
			inputT: &bintree.TreeNode{
				Val: 3,
				Right: &bintree.TreeNode{
					Val: 1,
				},
			},
			want: true,
		},
		{
			name: `[3,4,5,1,2,null,null,0][4,1,2]`,
			inputS: &bintree.TreeNode{
				Val: 3,
				Left: &bintree.TreeNode{
					Val: 4,
					Left: &bintree.TreeNode{
						Val: 1,
						Left: &bintree.TreeNode{
							Val: 0,
						},
					},
					Right: &bintree.TreeNode{
						Val: 2,
					},
				},
				Right: &bintree.TreeNode{
					Val: 5,
				},
			},
			inputT: &bintree.TreeNode{
				Val: 4,
				Left: &bintree.TreeNode{
					Val: 1,
				},
				Right: &bintree.TreeNode{
					Val: 2,
				},
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := isSubtree(tc.inputS, tc.inputT); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})

	}
}

func TestIsSubTreeWithPreOrder(t *testing.T) {
	type test struct {
		name   string
		inputS *bintree.TreeNode
		inputT *bintree.TreeNode
		want   bool
	}
	tests := []test{
		{
			name: "[1 1][1]",
			inputS: &bintree.TreeNode{
				Val: 1,
				Left: &bintree.TreeNode{
					Val: 1,
				},
			},
			inputT: &bintree.TreeNode{
				Val: 1,
			},
			want: true,
		},
		{
			name: "[4 1][4 1]",
			inputS: &bintree.TreeNode{
				Val: 4,
				Left: &bintree.TreeNode{
					Val: 1,
				},
			},
			inputT: &bintree.TreeNode{
				Val: 4,
				Left: &bintree.TreeNode{
					Val: 1,
				},
			},
			want: true,
		},
		{
			name: "[4 3 null 1][3 null 1]",
			inputS: &bintree.TreeNode{
				Val: 4,
				Left: &bintree.TreeNode{
					Val: 3,
					Right: &bintree.TreeNode{
						Val: 1,
					},
				},
			},
			inputT: &bintree.TreeNode{
				Val: 3,
				Right: &bintree.TreeNode{
					Val: 1,
				},
			},
			want: true,
		},
		{
			name: `[3,4,5,1,2,null,null,0][4,1,2]`,
			inputS: &bintree.TreeNode{
				Val: 3,
				Left: &bintree.TreeNode{
					Val: 4,
					Left: &bintree.TreeNode{
						Val: 1,
						Left: &bintree.TreeNode{
							Val: 0,
						},
					},
					Right: &bintree.TreeNode{
						Val: 2,
					},
				},
				Right: &bintree.TreeNode{
					Val: 5,
				},
			},
			inputT: &bintree.TreeNode{
				Val: 4,
				Left: &bintree.TreeNode{
					Val: 1,
				},
				Right: &bintree.TreeNode{
					Val: 2,
				},
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := isSubtreeWithPreOrder(tc.inputS, tc.inputT); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})

	}
}
