package connected

import (
	"fmt"
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
	fmt.Println("starting kosaraju")
	sccCounts := kosaraju(g)

	sccCountsStrs := make([]string, 0, 5)
	for _, c := range sccCounts {
		sccCountsStrs = append(sccCountsStrs, strconv.Itoa(c))
	}

	return strings.Join(sccCountsStrs, ",")
}

func makeAdjacencyList(edges [][]int) map[int]*node {
	al := make(map[int]*node, 0)

	for i := 0; i < len(edges); i++ {
		nv := edges[i][0]
		var tail, head *node
		if l, ok := al[nv]; !ok {
			tail = &node{ID: nv}
			// l = tail
			al[nv] = tail
		} else {
			tail = l
		}
		v := edges[i][1]
		if r, ok := al[v]; !ok {
			head = &node{ID: v}
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

func kosaraju(g map[int]*node) []int {
	// reverse graph
	// call DFS from every vertex of the revese graph to compute the finishing position for each vertex
	// call DFS from every vertex of the graph processed from highest to lowest finishing position
	//  to compute the identity of each vertex's strongly connected component

	// track visited nodes
	explored := make(map[int]bool, len(g))
	finishing := make([]*node, len(g))
	finishedMap := make(map[int]struct{}, len(g))
	// initialize finishing time to the len of g because we want to make the second pass
	// from the lowest finishing time the the highest. the lecture labels them the opposite
	// way and then says to go from highest to lowest.
	t := len(g)
	//
	// mark i as explored
	// for each arc
	// make a stack and provide it with a starting node
	for s, sn := range g {
		if explored[s] {
			continue
		}

		processNext := make([]*node, 0)
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
	headMax := 10
	if len(finishing) < headMax {
		headMax = len(finishing)
	}
	// set leader(i) := node s
	// s is the leader of the DFS which first discovered that node. Not just it's parent, but the value from the outer loop.
	// need this for the first pass? NO, only the second pass, it labels the SCC
	// considured  explored reversed, so if it is false, it's been explored
	// leaders := make([]int, len(g))
	leaderCount := make(map[int]int)
	finishedMap = make(map[int]struct{}, len(g))
	t = 0
	for _, sn := range finishing {
		if !explored[sn.ID] {
			continue
		}
		processNext := make([]*node, 0)
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
				if _, ok := leaderCount[sn.ID]; ok {
					leaderCount[sn.ID]++
				} else {
					leaderCount[sn.ID] = 1
				}
				t++
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
	// pad with zeros if less than 5
	for len(sccCounts) < 5 {
		sccCounts = append(sccCounts, 0)
	}
	return sccCounts[:5]
}

type node struct {
	ID int
	//outbound edges
	Outbound []*node
	Inbound  []*node
	// inbound count
	InboundCount int
	// outbound count
	OutboundCount int
}
