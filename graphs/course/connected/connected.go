package connected

import "fmt"

// Given a list of edges of a directed graph, where the first item is the tail and the second the head,
// output the sizes of the 5 largest SCCs(strongly connected component)s contained within. output them in
// decreasing order of sizes, separated by commas(no spaces). If there are less than 5 SCCs, output all
// there are and 0 for the remaining terms. ie. if there are only two SCCs which are 10 and 5, output
// 10,5,0,0,0

func fiveLargestSCCs(edges [][]int) string {
	// kosaraju should return a list of ints, each int is the SCC identity of a vertex.
	// the size of each SCC is the number of ints with the identity.
	// make a map of each identity to it's count
	//  or make a list of counts where the index is the identity
	// make a list of the counts ordered by largest to smallest
	// take the first 5 and if that isn't 5, pad with zeros and join
	// them into a comma separated string

	g := makeAdjacencyList(edges)
	_ = kosaraju(g)

	return "0,0,0,0,0"
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
	finishing := make([]int, len(g))
	// initialize finishing time to the len of g because we want to make the second pass
	// from the lowest finishing time the the highest. the lecture labels them the opposite
	// way and then says to go from highest to lowest.
	t := len(g)
	//
	// mark i as explored
	// for each arc
	// make a stack and provide it with a starting node
	for s := range g {
		if explored[s] {
			continue
		}
		processNext := make([]int, 0)
		processNext = append(processNext, s)
		for len(processNext) > 0 {
			fmt.Printf("pn: %v\n", processNext)
			n := g[processNext[len(processNext)-1]]
			if ok := explored[n.ID]; !ok {
				explored[n.ID] = true
			}

			unexploredCount := 0
			fmt.Printf("adding inbound: ")

			for _, v := range n.Inbound {
				if ok := explored[v.ID]; !ok {
					fmt.Printf("%d ", v.ID)
					processNext = append(processNext, v.ID)
					unexploredCount++
				}
			}
			fmt.Println("")
			// no more unexplored inbound(rev of outbound)
			if unexploredCount == 0 {
				fmt.Printf("finish %d at position %d\n", n.ID, t)
				t--
				finishing[t] = n.ID
				processNext = processNext[:len(processNext)-1] //
			}

		}
	}

	// set leader(i) := node s
	// s is the leader of the DFS which first discovered that node. Not just it's parent, but the value from the outer loop.
	// need this for the first pass? NO, only the second pass, it labels the SCC

	// leaders := make([]int, len(g))
	// for _, s := range finishing {
	// 	processNext := make([]int, 0)
	// 	processNext = append(processNext, s)

	// 	if ok := explored[]
	// }
	headMax := 10
	if len(finishing) < headMax {
		headMax = len(finishing)
	}
	fmt.Printf("finishing head %v\n", finishing[:headMax])
	fmt.Printf("finishing tail %v\n", finishing[len(finishing)-headMax:])

	return []int{}
}

type node struct {
	ID int
	//outbound edges
	Outbound []*node
	Inbound  []*node
	// inbound count
	InboundCount int
	Explored     bool
}
