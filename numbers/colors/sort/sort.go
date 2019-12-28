package sort

func sortColors(nums []int) {
	colorCounts := make([]int, 3)
	for i := 0; i < len(nums); i++ {
		colorCounts[nums[i]]++
	}
	p := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < colorCounts[i]; j++ {
			nums[p] = i
			p++
		}
	}
}

// based on the Dutch National Flag Problem
func sortColorsPointers(nums []int) {
	// pointer to the rightmost boundary of 0s
	p0 := 0
	// pointer to the leftmost boundary of 2s
	p2 := len(nums) - 1
	// index of current element
	curr := 0
	for curr <= p2 {
		if nums[curr] == 0 {
			nums[curr], nums[p0] = nums[p0], nums[curr]
			p0++
			curr++
		} else if nums[curr] == 2 {
			nums[curr], nums[p2] = nums[p2], nums[curr]
			p2--
		} else {
			curr++
		}
	}
}
