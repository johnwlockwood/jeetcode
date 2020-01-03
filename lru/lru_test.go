package lru

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	t.Run("swap neighbors", func(t *testing.T) {
		nodeA := &Node{Key: 1, Val: 1}
		nodeB := &Node{Key: 2, Val: 2}
		nodeC := &Node{Key: 3, Val: 3}
		nodeD := &Node{Key: 4, Val: 4}
		nodeE := &Node{Key: 5, Val: 5}
		connect(nodeA, nodeB)
		connect(nodeB, nodeC)
		connect(nodeC, nodeD)
		connect(nodeD, nodeE)
		printNodes(nodeA)
		swap(nodeB, nodeD)
		if nodeA.Next != nodeD {
			t.Errorf("A next should be D. got %v, want %v", nodeA.Next, nodeD)
		}
		if nodeD.Next != nodeC {
			t.Errorf("D next should be C. got %v, want %v", nodeD.Next, nodeC)
		}
		if nodeD.Prev != nodeA {
			t.Errorf("D Prev should be A. got %v, want %v", nodeD.Prev, nodeA)
		}
		if nodeC.Prev != nodeD {
			t.Errorf("C Prev should be D. got %v, want %v", nodeC.Prev, nodeD)
		}
		if nodeC.Next != nodeB {
			t.Errorf("C Next should be B. got %v, want %v", nodeC.Next, nodeB)
		}
		if nodeB.Prev != nodeC {
			t.Errorf("C Prev should be D. got %v, want %v", nodeB.Prev, nodeC)
		}
		if nodeB.Next != nodeE {
			t.Errorf("C Next should be B. got %v, want %v", nodeB.Next, nodeE)
		}

		printNodes(nodeA)
	})
}

func TestLRUCache(t *testing.T) {
	t.Run("Cap 2", func(t *testing.T) {
		cache := Constructor(2)
		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.PrintNodes()
		if got := cache.Get(1); got != 1 {
			t.Errorf("got %d, want %d", got, 1)
		}
		cache.PrintNodes()
		cache.Put(3, 3)
		if got := cache.Get(3); got != 3 {
			t.Errorf("got %d, want %d", got, 3)
		}
		cache.PrintNodes()
		if got := cache.Get(2); got != -1 {
			t.Errorf("got %d, want %d", got, -1)
		}
		cache.PrintNodes()
		cache.Put(3, 5)
		if got := cache.Get(3); got != 5 {
			t.Errorf("got %d, want %d", got, 5)
		}
		cache.PrintNodes()
	})

	t.Run("Cap 3", func(t *testing.T) {
		fmt.Println("cap 3")
		cache := Constructor(3)
		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.PrintNodes()
		if got := cache.Get(1); got != 1 {
			t.Errorf("got %d, want %d", got, 1)
		}
		cache.PrintNodes()
		cache.Put(3, 3)
		cache.PrintNodes()
		if got := cache.Get(3); got != 3 {
			t.Errorf("got %d, want %d", got, 3)
		}
		cache.PrintNodes()
		if got := cache.Get(2); got != 2 {
			t.Errorf("got %d, want %d", got, 2)
		}
		cache.PrintNodes()
		cache.Put(3, 5)
		cache.PrintNodes()
		if got := cache.Get(3); got != 5 {
			t.Errorf("got %d, want %d", got, 5)
		}
		cache.PrintNodes()
		fmt.Println("before adding 5")
		cache.Put(5, 5)
		cache.PrintNodes()
		if got := cache.Get(1); got != -1 {
			t.Errorf("got %d, want %d", got, -1)
		}
		cache.PrintNodes()
	})
}
