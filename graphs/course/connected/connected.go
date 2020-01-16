package connected

import (
	"sort"
	"strconv"
	"strings"
)

// Given a list of edges of a directed graph, where the first item is the tail and the second the head,
// output the sizes of the 5 largest SCCs(strongly connected component)s contained within. output them in
// decreasing order of sizes, separated by commas(no spaces). If there are less than 5 SCCs, output all
// there are and 0 for the remaining terms. ie. if there are only two SCCs which are 10 and 5, output
// 10,5,0,0,0

// FiveLargestSCCs returns the counts of the five largest SCCs in the given graph
func FiveLargestSCCs(edges [][]int) string {
	// kosaraju should return a list of ints, each int is the SCC identity of a vertex.
	// the size of each SCC is the number of ints with the identity.
	// make a map of each identity to it's count
	//  or make a list of counts where the index is the identity
	// make a list of the counts ordered by largest to smallest
	// take the first 5 and if that isn't 5, pad with zeros and join
	// them into a comma separated string
	g := makeReverseAdjacencyList(edges)

	gl := make([]*Node, 0, len(g))
	for _, n := range g {
		gl = append(gl, n)
	}
	leaders := FindSCCLeaders(gl)
	least := 5
	if len(leaders) < least {
		least = len(leaders)
	}
	counts := make([]int, 5)
	for i := 0; i < least; i++ {
		counts[i] = leaders[i].Count
	}

	sccCountsStrs := make([]string, 0, 5)
	for _, c := range counts {
		sccCountsStrs = append(sccCountsStrs, strconv.Itoa(c))
	}

	return strings.Join(sccCountsStrs, ",")
}

func makeReverseAdjacencyList(edges [][]int) map[int]*Node {
	// makes a reverse adjacency list, each node
	// is given it's inbound nodes
	al := make(map[int]*Node, 0)

	for i := 0; i < len(edges); i++ {
		nv := edges[i][0]
		var tail, head *Node
		if l, ok := al[nv]; !ok {
			tail = &Node{ID: nv}
			// l = tail
			al[nv] = tail
		} else {
			tail = l
		}
		v := edges[i][1]
		if r, ok := al[v]; !ok {
			head = &Node{ID: v}
			// r = head
			al[v] = head
		} else {
			head = r
		}
		head.Inbound = append(head.Inbound, tail)
		head.InDegree++
	}
	return al
}

// FindSCCLeaders returns leaders in all the SCCs of the given graph using kosaraju's algorithm
// for finding the SCCs of a graph.
func FindSCCLeaders(g []*Node) []*Leader {
	finishing := computeFinishOrder(g)
	return computeLeaders(finishing[:])
}

// Node is suited for representing and traversing a graph of ints
// forward and in reverse.
type Node struct {
	ID int
	//outbound edges
	Outbound []*Node
	Inbound  []*Node
	// inbound count
	InDegree int
	// outbound count
	OutDegree int
}

// Leader represents a node which should be a leader of an SCC and the SCC's count
type Leader struct {
	Node  *Node
	Count int
}

type byLeaderCount []*Leader

func (a byLeaderCount) Len() int      { return len(a) }
func (a byLeaderCount) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byLeaderCount) Less(i, j int) bool {
	return a[i].Count < a[j].Count
}

func computeFinishOrder(g []*Node) []*Node {
	// Compute finishing order
	// call DFS from every vertex of the reverse graph to compute the finishing position for each vertex
	// call DFS from every vertex of the graph processed from highest to lowest finishing position
	//  to compute the identity of each vertex's strongly connected component

	// Track which nodes have been seen while traversing the graph
	seen := make(map[int]struct{}, len(g))
	// Track which nodes have been added to the finishing order.
	// Since this is iterative instead of recursive, it doesn't
	// remove a node from the stack when first seen, only
	// when all it's outbound(inbound in this case, since it is reversed)
	// have been explored and it's
	// been essentially backtracked to. A recursive process
	// would deal with the post process after it's DFS had returned.
	finished := make(map[int]struct{}, len(g))

	finishing := make([]*Node, len(g))
	t := len(g)
	for _, sn := range g {
		if _, ok := seen[sn.ID]; ok {
			continue
		}
		processNext := make([]*Node, 0)
		processNext = append(processNext, sn)

		for len(processNext) > 0 {
			n := processNext[len(processNext)-1]
			if _, ok := seen[n.ID]; !ok {
				seen[n.ID] = struct{}{}
			}
			if n.InDegree == 0 {
				processNext = processNext[:len(processNext)-1]
				if _, ok := finished[n.ID]; ok {
					continue
				}
				t--
				finishing[t] = n
				finished[n.ID] = struct{}{}
			} else {
				for _, v := range n.Inbound {
					if _, ok := seen[v.ID]; !ok {
						processNext = append(processNext, v)
					}
					// transpose graph
					v.Outbound = append(v.Outbound, n)
					v.OutDegree++
				}
				// These are not needed anymore.
				n.InDegree = 0
				n.Inbound = nil
			}
		}
	}
	return finishing
}

func computeLeaders(finishing []*Node) []*Leader {
	// Compute leaders
	// A leader of a node is the node which started the DFS and discovers it. A node will be taken
	// from the finishing order and a DFS started, it will find all the nodes in it's SCC.
	seen := make(map[int]struct{}, len(finishing))
	finished := make(map[int]struct{}, len(finishing))

	leaders := make(byLeaderCount, 0)

	for _, sn := range finishing {
		if _, ok := seen[sn.ID]; ok {
			continue
		}
		leader := &Leader{
			Node: sn,
		}
		leaders = append(leaders, leader)
		processNext := make([]*Node, 0)
		processNext = append(processNext, sn)
		for len(processNext) > 0 {
			n := processNext[len(processNext)-1]
			if _, ok := seen[n.ID]; !ok {
				seen[n.ID] = struct{}{}
			}
			if n.OutDegree == 0 {
				processNext = processNext[:len(processNext)-1]
				if _, ok := finished[n.ID]; ok {
					continue
				}
				finished[n.ID] = struct{}{}
				leader.Count++
			} else {
				for _, v := range n.Outbound {
					n.OutDegree--
					if _, ok := seen[v.ID]; !ok {
						processNext = append(processNext, v)
					}
				}
			}
		}
	}
	sort.Sort(sort.Reverse(leaders))
	return leaders
}
