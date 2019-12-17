package permutations

func backtrackPractice(n, first int, nums []int, perms *[][]int) {
	if first == n {
		// output nums
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		*perms = append(*perms, numsCopy)
	} else {
		// iterate from first to n-1, swap first and i, backtrack and unswap
		for i := first; i < n; i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backtrackPractice(n, first+1, nums, perms)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
}

func permuteBacktrackPractice(nums []int) [][]int {
	// initialize perms matrix
	perms := make([][]int, 0)
	n := len(nums)
	nums1st := make([]int, n)
	copy(nums1st, nums)
	backtrackPractice(n, 0, nums1st, &perms)
	return perms
}
