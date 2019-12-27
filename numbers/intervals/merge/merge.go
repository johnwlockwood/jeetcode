package merge

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	marked := make([]bool, len(intervals))
	for i := 0; i < len(intervals)-1; i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[j][0] <= intervals[i][1] && !marked[j] {
				intervals[i][1] = int(math.Max(float64(intervals[i][1]), float64(intervals[j][1])))
				marked[j] = true
			}
		}
	}
	merged := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		if !marked[i] {
			merged = append(merged, intervals[i])
		}
	}
	return merged
}
