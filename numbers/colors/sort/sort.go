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
