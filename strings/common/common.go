package common

func findContiguousHistory(userA, userB []string) []string {
	// Longest Common Substring
	// Use Memoization, a component of Dynamic Programming
	// for each pair of strings, we need to store the count of contiguous prevous pairs matching
	// for each pair, if they match, look up the count of the previous pair and store that count +1
	// at this pair coordinates
	// keep track of the max count and the index of A of the pairs where that max was set
	// build the common history from the values of A from max index of A-result to the max index of A

	// create a matrix which counts contiguous common values
	lcs := make([][]int, len(userA)+1)
	for i := range lcs {
		lcs[i] = make([]int, len(userB)+1)
	}

	// userA = B A C E
	// userB = F A C
	//
	// 		    B A C E
	//		  0 0 0 0 0
	//		F 0 0 0 0 0
	//		A 0 0 1 0 0
	//		C 0 0 0 2 0
	for i := 1; i <= len(userA); i++ {
		for j := 1; j <= len(userB); j++ {
			if userA[i-1] == userB[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			}
		}
	}

	longestAIndex := -1
	result := -1
	for i := 0; i <= len(userA); i++ {
		for j := 0; j <= len(userB); j++ {
			if result < lcs[i][j] {
				result = lcs[i][j]
				longestAIndex = i
			}
		}
	}
	if result <= 0 {
		return []string{}
	}
	longest := make([]string, 0, result)
	for i := longestAIndex - result; i < longestAIndex; i++ {
		longest = append(longest, userA[i])
	}
	return longest
}
