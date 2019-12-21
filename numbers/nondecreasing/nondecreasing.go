package nondecreasing

import "fmt"

func checkPossibility(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	fmt.Printf("checking %v\n", nums)
	// highWater := nums[0]
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			nums[i-1] = nums[i]
			count++
			if count > 1 {
				return false
			}
			i = 0
		}
	}
	return true
}
