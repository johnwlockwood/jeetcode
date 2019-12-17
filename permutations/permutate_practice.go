package permutations

func backTrackPractice(first int, nums []int, perms *[][]int) {
	if first == len(nums) {
		// output copy of nums
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		*perms = append(*perms, numsCopy)
	} else {
		// iterate from first to n-1; swap nums first and i; backtrack with first+1; unswap
		for i := first; i < len(nums); i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backTrackPractice(first+1, nums, perms)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
}

func permuteBackTrackPractice(nums []int) [][]int {
	// initialize matrix; backtrack with 0 and return perms
	perms := make([][]int, 0)
	backTrackPractice(0, nums, &perms)
	return perms
}
