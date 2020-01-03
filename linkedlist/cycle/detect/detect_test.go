package detect

import "testing"

func TestDetectCycle(t *testing.T) {
	type test struct {
		input *ListNode
		pos   int
	}

	tests := []test{
		{
			input: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4}}}},
			pos:   1,
		},
		{
			input: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			pos:   -1,
		},
	}

	for _, tc := range tests {
		// connect list at pos
		var connectAt *ListNode
		if tc.pos >= 0 {
			connectAt = tc.input
			i := 0
			for connectAt != nil && connectAt.Next != nil && i < tc.pos {
				connectAt = connectAt.Next
				i++
			}
			node := tc.input
			for node != nil {
				if node.Next == nil {
					node.Next = connectAt
					node = nil
					break
				} else {
					node = node.Next
				}
			}
		}
		if got := detectCycle(tc.input); got != connectAt {
			t.Errorf("got %v, want %v", got, connectAt)
		}
	}
}
