package permutations

func backtrackPractice(n, first int, nums []int, perms *[][]int) {
	if first == n {
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		*perms = append(*perms, numsCopy)
	} else {
		// iterate from first to n-1, swap nums first and i, backtrack with first +1, unswap nums
		for i := first; i < n; i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backtrackPractice(n, first+1, nums, perms)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
}

func permuteBacktrackPractice(nums []int) [][]int {
	// init perms matrix
	perms := make([][]int, 0)
	backtrackPractice(len(nums), 0, nums, &perms)
	return perms
}
