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

func TestInorderTraversal(t *testing.T) {
	type test struct {
		name  string
		input *TreeNode
		want  []int
	}

	tests := []test{
		{
			name: "[7, 8, 3, 4, 2, 1]",
			input: &TreeNode{
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
			want: []int{7, 8, 3, 4, 2, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := InorderTraversal(tc.input); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestInorderTraversalIterative(t *testing.T) {
	type test struct {
		name  string
		input *TreeNode
		want  []int
	}

	tests := []test{
		{
			name: "[7, 8, 3, 4, 2, 1]",
			input: &TreeNode{
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
			want: []int{7, 8, 3, 4, 2, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := InorderTraversalIterative(tc.input); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestLevelOrder(t *testing.T) {
	type test struct {
		name  string
		input *TreeNode
		want  [][]int
	}

	tests := []test{
		{
			name: "1",
			input: &TreeNode{
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
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "4 levels",
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 55,
					},
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
						Left: &TreeNode{
							Val: 98,
						},
						Right: &TreeNode{
							Val: 99,
						},
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			want: [][]int{{3}, {9, 20}, {55, 15, 7}, {98, 99}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := levelOrderValues(tc.input); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestHeapifyMin(t *testing.T) {
	type test struct {
		name  string
		input *TreeNode
		want  []int
	}

	tests := []test{
		{
			name: "[7, 8, 3, 4, 2, 1]",
			input: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			want: []int{1, 8, 7},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			HeapifyMin(tc.input)
			vals := make([]int, 0)
			getValsPreOrder(tc.input, &vals)
			if got := vals; !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
