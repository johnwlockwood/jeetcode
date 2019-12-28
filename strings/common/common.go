package common

import "fmt"

func findContiguousHistory(userA, userB []string) []string {
	// create a matrix which counts contiguous common values
	lcs := make([][]int, len(userA)+1)
	for i := range lcs {
		lcs[i] = make([]int, len(userB)+1)
	}

	for i := 1; i <= len(userA); i++ {
		for j := 1; j <= len(userB); j++ {
			if userA[i-1] == userB[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			}
		}
	}

	longestI := -1
	longestJ := -1
	result := -1
	for i := 0; i <= len(userA); i++ {
		for j := 0; j <= len(userB); j++ {
			if result < lcs[i][j] {
				result = lcs[i][j]
				longestI = i
				longestJ = j
			}
		}
	}
	if result <= 0 {
		return []string{}
	}
	fmt.Printf("%v\n longest size %d\n", lcs, result)

	fmt.Printf("longest coord i:%d j:%d\n", longestI, longestJ)

	longest := make([]string, 0, result)
	for i := longestI - result; i < longestI; i++ {
		longest = append(longest, userA[i])
	}

	fmt.Printf("%v\n", longest)

	return longest
}
