package connected

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Given a list of edges of a directed graph, where the first item is the tail and the second the head,
// output the sizes of the 5 largest SCCs(strongly connected component)s contained within. output them in
// decreasing order of sizes, separated by commas(no spaces). If there are less than 5 SCCs, output all
// there are and 0 for the remaining terms. ie. if there are only two SCCs which are 10 and 5, output
// 10,5,0,0,0

func FiveLargestSCCs(edges [][]int) string {
	// kosaraju should return a list of ints, each int is the SCC identity of a vertex.
	// the size of each SCC is the number of ints with the identity.
	// make a map of each identity to it's count
	//  or make a list of counts where the index is the identity
	// make a list of the counts ordered by largest to smallest
	// take the first 5 and if that isn't 5, pad with zeros and join
	// them into a comma separated string

	g := makeAdjacencyList(edges)
	sccCounts := kosaraju(g)

	sccCountsStrs := make([]string, 0, 5)
	for _, c := range sccCounts {
		sccCountsStrs = append(sccCountsStrs, strconv.Itoa(c))
	}

	return strings.Join(sccCountsStrs, ",")
}

func makeAdjacencyList(edges [][]int) map[int]*node {
	al := make(map[int]*node, 0)
	ledges := len(edges)

	fmt.Printf("edges %d, last edge: %v\n", ledges, edges[len(edges)-1])
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

var dEBUG = false

func debug(s string) {
	if dEBUG {
		fmt.Printf("%s", s)
	}
}

func kosaraju(g map[int]*node) []int {
	// reverse graph
	// call DFS from every vertex of the revese graph to compute the finishing position for each vertex
	// call DFS from every vertex of the graph processed from highest to lowest finishing position
	//  to compute the identity of each vertex's strongly connected component

	// track visited nodes
	explored := make(map[int]bool, len(g))
	finishing := make([]int, len(g))
	finishedMap := make(map[int]struct{}, len(g))
	// initialize finishing time to the len of g because we want to make the second pass
	// from the lowest finishing time the the highest. the lecture labels them the opposite
	// way and then says to go from highest to lowest.
	t := len(g)
	//
	// mark i as explored
	// for each arc
	// make a stack and provide it with a starting node
	x := 0
	pCount := 0
	for s := range g {
		x++
		if x%5000 == 0 {
			fmt.Printf("x %d\t\t%v\n", x, time.Now())
		}
		if explored[s] {
			continue
		}

		debug(fmt.Sprintf("next unexplored graph node: %d\n", s))
		processNext := make([]int, 0)
		processNext = append(processNext, s)

		for len(processNext) > 0 {
			pCount++
			debug(fmt.Sprintf("pn: %v\n", processNext))
			n := g[processNext[len(processNext)-1]]
			if ok := explored[n.ID]; !ok {
				explored[n.ID] = true
			}

			unexploredCount := 0
			debug(fmt.Sprintf("adding inbound of %d: ", n.ID))

			for _, v := range n.Inbound {
				if !explored[v.ID] {
					debug(fmt.Sprintf("%d ", v.ID))
					processNext = append(processNext, v.ID)
					unexploredCount++
				}
			}
			debug(fmt.Sprintf("\n"))
			// no more unexplored inbound(rev of outbound)
			if unexploredCount == 0 {
				debug(fmt.Sprintf("finish %d at position %d\n", n.ID, t))
				processNext = processNext[:len(processNext)-1]
				if len(processNext) > 0 {
					debug(fmt.Sprintf("backtracking to %d\n", processNext[len(processNext)-1]))
				}
				if _, ok := finishedMap[n.ID]; ok {
					continue
				}
				t--
				finishing[t] = n.ID
				finishedMap[n.ID] = struct{}{}
				if t%5000 == 0 {
					fmt.Printf("finishing t %d\t\tpCount %d\t\t%v\n", t, pCount, time.Now())
				}
			}

		}
	}
	headMax := 10
	if len(finishing) < headMax {
		headMax = len(finishing)
	}
	fmt.Printf("finishing head %v\n", finishing[:headMax])
	fmt.Printf("finishing tail %v\n", finishing[len(finishing)-headMax:])
	// set leader(i) := node s
	// s is the leader of the DFS which first discovered that node. Not just it's parent, but the value from the outer loop.
	// need this for the first pass? NO, only the second pass, it labels the SCC
	// considured  explored reversed, so if it is false, it's been explored
	// leaders := make([]int, len(g))
	leaderCount := make(map[int]int)
	finishedMap = make(map[int]struct{}, len(g))
	t = 0
	pCount = 0
	for _, s := range finishing {
		if !explored[s] {
			continue
		}
		sn := g[s]
		processNext := make([]int, 0)
		processNext = append(processNext, s)
		for len(processNext) > 0 {
			pCount++
			debug(fmt.Sprintf("pn: %v\n", processNext))
			n := g[processNext[len(processNext)-1]]
			if explored[n.ID] {
				explored[n.ID] = false
			}
			unexploredCount := 0
			debug(fmt.Sprintf("adding outbound: "))
			for _, v := range n.Outbound {
				if explored[v.ID] {
					debug(fmt.Sprintf("%d ", v.ID))
					processNext = append(processNext, v.ID)
					unexploredCount++
				}
			}
			debug(fmt.Sprintf("\n"))
			// no more unexplored inbound(rev of outbound)
			if unexploredCount == 0 {
				processNext = processNext[:len(processNext)-1]
				debug(fmt.Sprintf("leader of %d is %d\n", t, n.ID))
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
				t++
				if t%5000 == 0 {
					fmt.Printf("leader t %d\t\tpCount %d\t\t%v\n", t, pCount, time.Now())
				}
			}
		}
	}
	// no need to hold all the leaders, just the counts
	fmt.Printf("leaders finished\n")

	sccCounts := make([]int, 0, len(leaderCount))
	for _, v := range leaderCount {
		sccCounts = append(sccCounts, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sccCounts)))
	// pad with zeros if less than 5
	for len(sccCounts) < 5 {
		sccCounts = append(sccCounts, 0)
	}

	// sort.Sort(sort.IntSlice(leaders))
	// fmt.Printf("leaders %v\n", leaders)

	fmt.Printf("sccCounts %v\n", sccCounts[:5])

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
