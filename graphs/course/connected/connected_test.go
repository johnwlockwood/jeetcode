package connected

import (
	"archive/zip"
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestFiveLargestSCCSmallGraph(t *testing.T) {
	type test struct {
		name  string
		input [][]int
		want  string
	}
	tests := []test{
		{
			name: "nine nodes",
			input: [][]int{
				{1, 4},
				{4, 7},
				{7, 1},
				{9, 7},
				{9, 3},
				{3, 6},
				{6, 9},
				{8, 6},
				{2, 8},
				{5, 2},
				{8, 5},
			},
			want: "3,3,3,0,0",
		},
		{
			name: "two disjointed scc, with one out",
			input: [][]int{
				{1, 6},
				{6, 3},
				{3, 1},
				{9, 2},
				{2, 7},
				{7, 9},
				{8, 8},
			},
			want: "3,3,1,0,0",
		},
		{
			name: "four disjointed scc, one no outbound",
			input: [][]int{
				{1, 6},
				{6, 3},
				{3, 1},
				{99, 2},
				{2, 7},
				{7, 99},
				{8, 5},
			},
			want: "3,3,1,1,0",
		},
		{
			name: "multi in",
			input: [][]int{
				{1, 6},
				{6, 3},
				{3, 1},
				{99, 2},
				{2, 7},
				{7, 99},
				{8, 5},
				{11, 5},
				{5, 11},
			},
			want: "3,3,2,1,0",
		},
		{
			name: "full scc",
			input: [][]int{
				{1, 5},
				{1, 6},
				{1, 7},
				{5, 1},
				{5, 6},
				{5, 7},
				{6, 1},
				{6, 5},
				{6, 7},
				{7, 1},
				{7, 5},
				{7, 6},
			},
			want: "4,0,0,0,0",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := fiveLargestSCCs(tc.input); got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func TestFiveLargestSCC(t *testing.T) {

	tests := map[string]string{
		"SCC.txt": "1,2,3,4,5",
	}

	r, err := zip.OpenReader("testdata/SCC.txt.zip")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	for _, f := range r.File {
		want, ok := tests[f.Name]
		if !ok {
			continue
		}
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		defer rc.Close()
		if err != nil {
			t.Fatal(err)
		}

		edges := make([][]int, 0)

		scanner := bufio.NewScanner(rc)
		i := 0
		for scanner.Scan() {
			if i > 300000 {
				break
			}
			line := scanner.Text()
			edgeStr := strings.Split(line, " ")
			if len(edgeStr) < 2 {
				t.Fatalf("bad data %s", line)
				break
			}
			tail, err := strconv.Atoi(edgeStr[0])
			if err != nil {
				t.Fatal(err)
				break
			}
			head, err := strconv.Atoi(edgeStr[1])
			if err != nil {
				t.Fatal(err)
				break
			}

			edges = append(edges, []int{tail, head})
			i++
		}
		t.Logf("there are %d edges and i == %d", len(edges), i)
		fmt.Printf("there are %d edges and i == %d\n", len(edges), i)
		got := fiveLargestSCCs(edges)
		if got != want {
			t.Errorf("for %s: got %v, want %v", f.Name, got, want)
		}

	}
	t.Error("watt?")
}
