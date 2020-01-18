package topological

import "testing"

func TestTopoOrder(t *testing.T) {
	type test struct {
		name  string
		edges [][]int
		want  int
		isDAG bool
	}

	tests := []test{
		{
			name: "two disjointed DAGs",
			edges: [][]int{
				{1, 2},
				{2, 4},
				{5, 3},
			},
			want:  5,
			isDAG: true,
		},
		{
			name: "digraph with cycle",
			edges: [][]int{
				{1, 2},
				{2, 4},
				{5, 3},
				{4, 5},
				{5, 2},
			},
			want:  5,
			isDAG: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := TopoOrder(NodesFromEdges(tc.edges)); tc.isDAG && len(got) != tc.want || !tc.isDAG && len(got) == tc.want {
				t.Errorf("got %d, want %d", len(got), tc.want)
			}
		})

	}
}
