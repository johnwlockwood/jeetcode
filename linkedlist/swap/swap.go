package swap

// ListNode is a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Recursive Approach
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	firstNode := head
	secondNode := head.Next

	firstNode.Next = swapPairs(secondNode.Next)
	secondNode.Next = firstNode

	return secondNode
}

// Iterative Approach
func swapPairsIterative(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prevNode := dummy
	for head != nil && head.Next != nil {
		// Nodes to be swapped
		f := head
		s := head.Next

		// Swapping
		prevNode.Next = s
		f.Next = s.Next
		s.Next = f

		// Reinitializing the head and prevNode for next swap
		prevNode = f
		head = f.Next
	}
	return dummy.Next
}
