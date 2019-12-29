package permutations

func backTrackPractice(first int, nums []int, perms *[][]int) {
	n := len(nums)
	if first == n {
		// append copy of nums to perms
		numsCopy := make([]int, n)
		copy(numsCopy, nums)
		*perms = append(*perms, numsCopy)
	} else {
		// iterate from first to n-1; swap first and i; backtrack from first+1; unswap
		for i := first; i < n; i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backTrackPractice(first+1, nums, perms)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
}

func PermuteBackTrackPractice(nums []int) [][]int {
	// init perms matrix; backtrack from 0;return perms
	perms := make([][]int, 0)
	backTrackPractice(0, nums, &perms)
	return perms
}
