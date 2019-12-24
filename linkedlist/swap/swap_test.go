package swap

import (
	"fmt"
	"strings"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	type test struct {
		input *ListNode
		want  *ListNode
	}

	tests := []test{
		{
			input: &ListNode{
				Val: 1,
			},
			want: &ListNode{
				Val: 1,
			},
		},
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
		},
	}
	for _, tc := range tests {
		if got := swapPairs(tc.input); getVal(got) != getVal(tc.want) {
			t.Errorf("got %v, want %v", getVal(got), getVal(tc.want))
		}
	}
}

func nextVal(b *strings.Builder, n *ListNode) {
	if n.Next != nil {
		nextVal(b, n.Next)
	}
	fmt.Fprintf(b, "%d ", n.Val)
}

func getVal(n *ListNode) string {
	var b strings.Builder
	nextVal(&b, n)
	return b.String()
}
