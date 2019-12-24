package swap

// ListNode is a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Naive solution to swap pairs
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nH := swapPairs(head.Next.Next)
	tmp := head
	head = head.Next
	head.Next = tmp
	head.Next.Next = nH
	return head
}
