package invert

import (
	"fmt"
	"testing"

	"github.com/johnwlockwood/jeetcode/bintree"
)

func TestInvertTree(t *testing.T) {
	type test struct {
		input *bintree.TreeNode
		want  *bintree.TreeNode
	}
	tests := []test{
		{
			input: &bintree.TreeNode{
				Val: 4,
			},
			want: &bintree.TreeNode{
				Val: 4,
			},
		},
		{
			input: &bintree.TreeNode{
				Val: 1,
				Left: &bintree.TreeNode{
					Val: 5,
				},
				Right: &bintree.TreeNode{
					Val: 6,
				},
			},
			want: &bintree.TreeNode{
				Val: 1,
				Left: &bintree.TreeNode{
					Val: 6,
				},
				Right: &bintree.TreeNode{
					Val: 5,
				},
			},
		},
		{
			input: &bintree.TreeNode{
				Val: 1,
				Left: &bintree.TreeNode{
					Val: 5,
				},
				Right: &bintree.TreeNode{
					Val: 6,
				},
			},
			want: &bintree.TreeNode{
				Val: 1,
				Left: &bintree.TreeNode{
					Val: 6,
				},
				Right: &bintree.TreeNode{
					Val: 5,
				},
			},
		},
	}
	for _, tc := range tests {
		if got := invertTree(tc.input); fmt.Sprintf("%v", bintree.GetVals(got)) != fmt.Sprintf("%v", bintree.GetVals(tc.want)) {
			t.Errorf("got %v, want %v\n", bintree.GetVals(got), bintree.GetVals(tc.want))
		} else {
			fmt.Printf("success got %v\n", bintree.GetVals(got))
		}
	}
}
