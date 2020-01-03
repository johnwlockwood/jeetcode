package lru

type LRUCache struct {
	capacity int
	preHead  *Node
	postTail *Node
	cache    map[int]*Node
}

func Constructor(capacity int) LRUCache {
	preHead := &Node{}
	postTail := &Node{}
	connect(preHead, postTail)
	// preHead.Next = postTail
	// postTail.Prev = preHead

	return LRUCache{
		capacity: capacity,
		preHead:  preHead,
		postTail: postTail,
		cache:    make(map[int]*Node, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if ll, ok := this.cache[key]; ok {
		this.remove(ll)
		this.add(ll)
		return ll.Val
	}
	return -1
}

func (this *LRUCache) add(node *Node) {
	if node != this.preHead.Next {
		temp := this.preHead.Next
		node.Next = this.preHead.Next
		node.Prev = this.preHead
		this.preHead.Next = node
		temp.Prev = node
		this.cache[node.Key] = node
	}
}

func (this *LRUCache) remove(node *Node) {
	// remove an existing node
	temp := node.Prev
	temp.Next = node.Next
	node.Next.Prev = temp
}

func (this *LRUCache) removeTail() {
	// remove tail from cache
	temp := this.postTail.Prev.Prev
	delete(this.cache, this.postTail.Prev.Key)
	this.postTail.Prev = temp
	temp.Next = this.postTail
}

func (this *LRUCache) Put(key int, value int) {
	// if key exists, update the value and make sure it is at the head
	// if it doesn't exist, see if cache is at capacity and evict the
	//  tail and insert this node at head
	if ll, ok := this.cache[key]; ok {
		this.remove(ll)
		ll.Val = value
		this.add(ll)
	} else {
		this.add(&Node{Key: key, Val: value})
		if len(this.cache) > this.capacity {
			this.removeTail()
		}
	}
}

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

func connect(first, second *Node) {
	first.Next = second
	second.Prev = first
}

func insertNext(node *Node, key, val int) *Node {
	// insert new node in at node.Next
	var temp *Node
	temp = node.Next
	newNode := &Node{Key: key, Val: val}
	connect(node, newNode)
	if temp != nil {
		connect(newNode, temp)
	}
	return newNode
}

func removePrev(node *Node) {
	if node.Prev != nil {
		if node.Prev.Prev != nil {
			connect(node.Prev.Prev, node)
		} else {
			node.Prev = nil
		}
	}
}

func areTheyNeighbors(a, b *Node) bool {
	return a.Next == b.Prev && b.Prev == a || a.Prev == b && b.Next == a
}

func refreshOuterPointers(a *Node) {
	if a.Prev != nil {
		a.Prev.Next = a
	}

	if a.Next != nil {
		a.Next.Prev = a
	}
}

func swap(a, b *Node) {
	swapperVector := new([4]*Node)
	var temp *Node

	if a == b {
		return
	}

	if b.Next == a {
		temp = a
		a = b
		b = temp
	}

	swapperVector[0] = a.Prev
	swapperVector[1] = b.Prev
	swapperVector[2] = a.Next
	swapperVector[3] = b.Next

	if areTheyNeighbors(a, b) {
		a.Prev = swapperVector[2]
		b.Prev = swapperVector[0]
		a.Next = swapperVector[3]
		b.Next = swapperVector[1]
	} else {
		a.Prev = swapperVector[1]
		b.Prev = swapperVector[0]
		a.Next = swapperVector[3]
		b.Next = swapperVector[2]
	}

	refreshOuterPointers(a)
	refreshOuterPointers(b)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
