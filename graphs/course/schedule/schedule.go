package schedule

// There are a total of n courses you have to take, labeled from 0 to n-1.

// Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]

// Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?

// Example 1:

// Input: 2, [[1,0]]
// Output: true
// Explanation: There are a total of 2 courses to take.
//              To take course 1 you should have finished course 0. So it is possible.
// Example 2:

// Input: 2, [[1,0],[0,1]]
// Output: false
// Explanation: There are a total of 2 courses to take.
//              To take course 1 you should have finished course 0, and to take course 0 you should
//              also have finished course 1. So it is impossible.
// Note:

// The input prerequisites is a graph represented by a list of edges, not adjacency matrices. Read more about how a graph is represented.
// You may assume that there are no duplicate edges in the input prerequisites.

/*
Design:

This looks like a cycle in this directed graph would make every class
in the cycle and every class which has a prerequisite class in the a cycle impossible to take.

make a list of each class

Find each cycle and it's start point and add each of them to
a map of impossible nodes.

remove each impossible node from the class list
for every class in the class list, look at it's prereqs and if any are in the impossible map, eliminate it too.

compare class list length to number of required classes

Is it helpful to do a union find to group classes that are in the same graph? I don't think so.
If I iterate through the list of edges and...
with an edge's left, iterate through the rest of the edges looking for it's prereqs. * no, translate to a list of classes *
sort the list, by l then r. cause it is free and might help.

**
Where to start?

start at each root.

a root is a class with no prerequisite.

DFS from each root through the classes required by it.

record if a root has been explored, so when searching from another root, not to repeat

mark a node if it is impossible to take.

****

how to find each root?
make a map of classes to the list of immediate prerequisites. and then of those, find those which have none.

how to find which classes are required by it?

make a map of classes the ones which immediately require it.

	Do I add the list as children of the node? This may be better to traverse. since you don't need to refer to the map. I'll start with out that.

how do I get a list of classes?

iterate over each node and add each left an right to a set.

how to count the possible nodes?

iterate over each node and increment count if it is not marked as impossible.

****

design of the node struct

{
	Id int
	Possible bool // instead, it is possible by being in the order at the end
	Explored bool // maybe instead of this, remove it from the graph to the order queue
	Children/Required by
}

****

TIL about adjacency list representation of a graph which is good for exploring sparse graphs.
TIL about topological sort

Design using a topological sort

build the a graph then remove them to a topological sort
when processNext is empty, compare the number of nodes to the number of courses needed

*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	order := make([]*course, 0)
	processNext := make([]*course, 0)
	courseList := makeUnconnectedCourses(prerequisites)
	connectCourses(courseList, prerequisites)
	// prime process Next with all courses which have none other required.
	for _, c := range courseList {
		if c.RequiredCount == 0 {
			processNext = append(processNext, c)
		}
	}
	for len(processNext) > 0 {
		n := processNext[0]
		processNext = processNext[1:]
		for _, x := range n.RequiredBy {
			x.RequiredCount--
			if x.RequiredCount == 0 {
				processNext = append(processNext, x)
			}
		}
		order = append(order, n)
	}

	return len(order) >= numCourses
}

func makeUnconnectedCourses(prerequisites [][]int) map[int]*course {
	courseList := make(map[int]*course, 0)
	for i := 0; i < len(prerequisites); i++ {
		nv := prerequisites[i][0]
		if _, ok := courseList[nv]; !ok {
			courseList[nv] = &course{ID: nv}
		}
		v := prerequisites[i][1]
		if _, ok := courseList[v]; !ok {
			courseList[v] = &course{ID: v}
		}
	}
	return courseList
}

func connectCourses(courseList map[int]*course, prerequisites [][]int) {
	for i := 0; i < len(prerequisites); i++ {
		l := prerequisites[i][0]
		r := prerequisites[i][1]
		cl := courseList[l]
		cr := courseList[r]
		// add cr to cl.RequiredBy
		// increment cr.RequiredCount
		cl.RequiredBy = append(cl.RequiredBy, cr)
		cr.RequiredCount++
	}
}

type course struct {
	ID int
	//outbound edges
	RequiredBy []*course
	// inbound count
	RequiredCount int
}
