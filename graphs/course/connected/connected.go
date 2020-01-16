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

	g := makeAdjacencyList(edges)

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

	// fmt.Println("starting kosaraju")
	// sccCounts := topSCCCounts(g, 5)

	sccCountsStrs := make([]string, 0, 5)
	for _, c := range counts {
		sccCountsStrs = append(sccCountsStrs, strconv.Itoa(c))
	}

	return strings.Join(sccCountsStrs, ",")
}

func makeAdjacencyList(edges [][]int) map[int]*Node {
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
		tail.Outbound = append(tail.Outbound, head)
		tail.OutboundCount++
		head.Inbound = append(head.Inbound, tail)
		head.InboundCount++
	}
	return al
}

// FindSCCLeaders returns leaders in all the SCCs of the given graph
func FindSCCLeaders(g []*Node) []*Leader {
	explored := make(map[int]bool, len(g))

	// compute finishing order
	finishing := make([]*Node, len(g))
	finishedMap := make(map[int]bool, len(g))
	t := len(g)
	for _, sn := range g {
		if explored[sn.ID] {
			continue
		}
		processNext := make([]*Node, 0)
		processNext = append(processNext, sn)

		for len(processNext) > 0 {
			n := processNext[len(processNext)-1]
			if ok := explored[n.ID]; !ok {
				explored[n.ID] = true
			}
			if n.InboundCount == 0 {
				processNext = processNext[:len(processNext)-1]
				if _, ok := finishedMap[n.ID]; ok {
					continue
				}
				t--
				finishing[t] = n
				finishedMap[n.ID] = true
			} else {
				for _, v := range n.Inbound {
					if !explored[v.ID] {
						processNext = append(processNext, v)
					}
					n.InboundCount--
				}
			}
		}
	}

	// compute leaders
	leaders := make(byLeaderCount, 0)
	for _, sn := range finishing {
		if !explored[sn.ID] {
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
			if explored[n.ID] {
				explored[n.ID] = false
			}
			if n.OutboundCount == 0 {
				processNext = processNext[:len(processNext)-1]
				// use finishedMap in reverse and just check the value
				// don't check if the key
				// exists, because the finish ordering filled them all in
				if ok := finishedMap[n.ID]; !ok {
					continue
				}
				finishedMap[n.ID] = false
				leader.Count++
			} else {
				for _, v := range n.Outbound {
					if explored[v.ID] {
						processNext = append(processNext, v)
					}
					n.OutboundCount--
				}
			}
		}
	}
	sort.Sort(sort.Reverse(leaders))
	return leaders
}

func topSCCCounts(g map[int]*Node, topCount int) []int {
	// kosaraju's algorithm for finding the strong components of a graph

	// reverse graph
	// call DFS from every vertex of the revese graph to compute the finishing position for each vertex
	// call DFS from every vertex of the graph processed from highest to lowest finishing position
	//  to compute the identity of each vertex's strongly connected component

	// track visited nodes
	explored := make(map[int]bool, len(g))
	finishing := make([]*Node, len(g))
	finishedMap := make(map[int]struct{}, len(g))
	// initialize finishing time to the len of g because we want to make the second pass
	// from the lowest finishing time the the highest. the lecture labels them the opposite
	// way and then says to go from highest to lowest.
	t := len(g)
	//
	// mark i as explored
	// for each arc
	// make a stack and provide it with a starting Node
	for s, sn := range g {
		if explored[s] {
			continue
		}

		processNext := make([]*Node, 0)
		processNext = append(processNext, sn)

		for len(processNext) > 0 {
			n := processNext[len(processNext)-1]
			if ok := explored[n.ID]; !ok {
				explored[n.ID] = true
			}
			if n.InboundCount == 0 {
				processNext = processNext[:len(processNext)-1]
				if _, ok := finishedMap[n.ID]; ok {
					continue
				}
				t--
				finishing[t] = n
				finishedMap[n.ID] = struct{}{}
			} else {
				for _, v := range n.Inbound {
					if !explored[v.ID] {
						processNext = append(processNext, v)
					}
					n.InboundCount--
				}
			}
		}
	}
	// set leader(i) := Node s
	// s is the leader of the DFS which first discovered that Node. Not just it's parent, but the value from the outer loop.
	// need this for the first pass? NO, only the second pass, it labels the SCC
	// considured  explored reversed, so if it is false, it's been explored
	// leaders := make([]int, len(g))
	leaderCount := make(map[int]int)
	finishedMap = make(map[int]struct{}, len(g))
	// t = 0
	for _, sn := range finishing {
		if !explored[sn.ID] {
			continue
		}
		processNext := make([]*Node, 0)
		processNext = append(processNext, sn)
		for len(processNext) > 0 {
			n := processNext[len(processNext)-1]
			if explored[n.ID] {
				explored[n.ID] = false
			}
			if n.OutboundCount == 0 {
				processNext = processNext[:len(processNext)-1]
				if _, ok := finishedMap[n.ID]; ok {
					continue
				}
				finishedMap[n.ID] = struct{}{}
				// leaders[t] = sn.ID
				if _, ok := leaderCount[sn.ID]; ok {
					leaderCount[sn.ID]++
				} else {
					leaderCount[sn.ID] = 1
				}
				// t++
			} else {
				for _, v := range n.Outbound {
					if explored[v.ID] {
						processNext = append(processNext, v)
					}
					n.OutboundCount--
				}
			}
		}
	}
	// no need to hold all the leaders, just the counts

	sccCounts := make([]int, 0, len(leaderCount))
	for _, v := range leaderCount {
		sccCounts = append(sccCounts, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sccCounts)))
	// pad with zeros if less than topCount
	for len(sccCounts) < topCount {
		sccCounts = append(sccCounts, 0)
	}
	return sccCounts[:topCount]
}

// Node is suited for representing and traversing a graph of ints
// forward and in reverse.
type Node struct {
	ID int
	//outbound edges
	Outbound []*Node
	Inbound  []*Node
	// inbound count
	InboundCount int
	// outbound count
	OutboundCount int
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
