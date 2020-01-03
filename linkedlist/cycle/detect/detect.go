package detect

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://en.wikipedia.org/wiki/Cycle_detection#Tortoise_and_hare
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	tortoise := head.Next
	hare := head.Next.Next
	for tortoise != hare {
		if hare == nil || hare.Next == nil {
			return nil
		}
		tortoise = tortoise.Next
		hare = hare.Next.Next
	}

	// find the position of the first repitition.
	mu := 0
	tortoise = head
	for tortoise != hare {
		tortoise = tortoise.Next
		hare = hare.Next
		mu++
	}

	// lam := 1
	// hare = tortoise.Next
	// for tortoise != hare {
	// 	hare = hare.Next
	// 	lam++
	// }
	return hare
}
