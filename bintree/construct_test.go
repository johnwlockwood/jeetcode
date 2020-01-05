package bintree

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	type test struct {
		inorder   []int
		postorder []int
		want      *TreeNode
	}

	tests := []test{
		{
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		{
			inorder:   []int{2, 3, 1},
			postorder: []int{3, 2, 1},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
		},
		{
			inorder:   []int{3, 2, 1},
			postorder: []int{3, 2, 1},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
		},
	}

	t.Run("buildTree1", func(t *testing.T) {
		for _, tc := range tests {
			got := buildTree(tc.inorder, tc.postorder)
			gotV := GetVals(got)
			want := GetVals(tc.want)
			if !reflect.DeepEqual(gotV, want) {
				t.Errorf("got %v, want %v", gotV, want)
			}
		}
	})

	t.Run("buildTree2", func(t *testing.T) {
		for _, tc := range tests {
			got := buildTree2(tc.inorder, tc.postorder)
			gotV := GetVals(got)
			want := GetVals(tc.want)
			if !reflect.DeepEqual(gotV, want) {
				t.Errorf("got %v, want %v", gotV, want)
			}
		}
	})

}

// go test -v -bench "BenchmarkBuildTree/buildTreeA" -cpuprofile cpuA.out -memprofile memA.prof -run="Benchmark" ./bintree
// go tool pprof -web cpuA.out
// go tool pprof -web memA.prof
// go test -v -bench "BenchmarkBuildTree/buildTreeB/" -cpuprofile cpuB.out -memprofile memB.prof -run="Benchmark" ./bintree
// go tool pprof -web cpuB.out
// go tool pprof -web memB.prof
func BenchmarkBuildTree(b *testing.B) {
	type bench struct {
		name      string
		inorder   []int
		postorder []int
	}

	bigI := make([]int, 50)
	for i := 0; i < 50; i++ {
		bigI[i] = i
	}

	benches := []bench{
		{
			name:      "[9, 3, 15, 20, 7][9, 15, 7, 20, 3]",
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
		},
		{
			name:      "[2, 3, 1][3, 2, 1]",
			inorder:   []int{2, 3, 1},
			postorder: []int{3, 2, 1},
		},
		{
			name:      "[3, 2, 1][3, 2, 1]",
			inorder:   []int{3, 2, 1},
			postorder: []int{3, 2, 1},
		},
		{
			name:      "big",
			inorder:   bigI,
			postorder: bigI,
		},
	}

	for _, bc := range benches {
		b.Run("buildTreeA/"+bc.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = buildTree(bc.inorder, bc.postorder)
			}
		})
	}

	for _, bc := range benches {
		b.Run("buildTreeB/"+bc.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = buildTree2(bc.inorder, bc.postorder)
			}
		})
	}

}
