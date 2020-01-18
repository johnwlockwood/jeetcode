package topological

// topological is an example of a topological ordering algorithm for a simple digraph

// Node is a simple node for a digraph
type Node struct {
	ID int
	//outbound edges
	Outbound []*Node
	Indegree int
}

// TopoOrder takes an adjacency list and returns
// nodes in topological order
// iterate over node list, finding all with indegree zero.
// * calculate the indegree while building the adjacency list from the edge list,
// 		every time you connect a tail, to a head, increment the indegree of the head.
// add them to a stack.
// While the stack is not empty
// pop a node of the stack and conceptually remove it from the graph.
// for every outbound node, decrement that node's indegree and if that node then has indegree zero, add it to the stack
// append the node to order.
func TopoOrder(g []*Node) []*Node {
	// Nodes are expected made as and adjacency list
	// with Outbound nodes set
	// Returns a slice of Nodes in topological order
	order := make([]*Node, 0, len(g))
	s := make([]*Node, 0)
	for _, sn := range g {
		if sn.Indegree == 0 {
			s = append(s, sn)
		}
	}
	for len(s) > 0 {
		n := s[len(s)-1]
		s = s[:len(s)-1]
		for _, v := range n.Outbound {
			v.Indegree--
			if v.Indegree == 0 {
				s = append(s, v)
			}
		}
		order = append(order, n)
	}
	return order
}

// NodesFromEdges make a Node based adjacency list from edges
// calculating the heads' indegree.
func NodesFromEdges(edges [][]int) []*Node {
	alMap := make(map[int]*Node, 0)
	for i := 0; i < len(edges); i++ {
		nv := edges[i][0]
		if _, ok := alMap[nv]; !ok {
			alMap[nv] = &Node{ID: nv}
		}
		v := edges[i][1]
		if _, ok := alMap[v]; !ok {
			alMap[v] = &Node{ID: v}
		}
	}
	al := make([]*Node, 0, len(alMap))

	for i := 0; i < len(edges); i++ {
		l := edges[i][0]
		r := edges[i][1]
		tail := alMap[l]
		head := alMap[r]
		// add head to tail.RequiredBy
		// increment head.RequiredCount
		tail.Outbound = append(tail.Outbound, head)
		head.Indegree++
	}

	for _, v := range alMap {
		al = append(al, v)
	}

	return al
}
