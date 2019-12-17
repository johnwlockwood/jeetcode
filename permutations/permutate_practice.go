package permutations

func backtrack3(n, first int, nums []int, perms *[][]int) {
	if first == n {
		// copy nums to output perms
		numsCopy := make([]int, n)
		copy(numsCopy, nums)
		*perms = append(*perms, numsCopy)
	} else {
		// iterate from first to n-1, swapping nums first i, backtrack and then unswap
		for i := first; i < n; i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backtrack3(n, first+1, nums, perms)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
}

func permuteBacktrack3(nums []int) [][]int {
	// intialize the output matrix
	perms := make([][]int, 0)
	n := len(nums)
	nums1st := make([]int, n)
	copy(nums1st, nums)
	backtrack3(n, 0, nums, &perms)
	return perms
}
