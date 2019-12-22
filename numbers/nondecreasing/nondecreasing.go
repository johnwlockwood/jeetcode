package nondecreasing

import "fmt"

// easy problem

// naive attempt
func checkPossibility(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	fmt.Printf("checking %v\n", nums)
	count := 0
	for i := 1; i < len(nums); i++ {
		// fmt.Printf("%v > %v?\n", nums[i-1], nums[i])
		if nums[i-1] > nums[i] || i != len(nums)-1 && nums[i+1] < nums[i] {
			fmt.Printf("\t i-1: %v i: %v", nums[i-1], nums[i])
			if i != len(nums)-1 {
				fmt.Printf("\t i+1: %v", nums[i+1])
			}
			fmt.Println("")
			count++
			if count > 1 {
				return false
			}
			if i == 1 && nums[i-1] > nums[i] {
				fmt.Println("case 1")
				nums[i-1] = nums[i]
			} else if nums[i-1] > nums[i] && (i == len(nums)-1 || nums[i+1] > nums[i]) {
				fmt.Println("case 2")
				nums[i] = nums[i-1]
				if i != len(nums)-1 && nums[i] < nums[i+1] {
					nums[i] = nums[i+1]
				}
				if nums[i-1] > nums[i] || (i != len(nums)-1 && nums[i] > nums[i+1]) {
					fmt.Printf("\t%v > %v?\n", nums[i-1], nums[i])
					return false
				}
			} else if i != len(nums)-1 && nums[i+1] < nums[i] {
				fmt.Println("case 3")
				if nums[i-1] < nums[i+1] {
					fmt.Println("case 3.1")
					nums[i] = nums[i-1]
				} else {
					fmt.Println("case 3.2")
					nums[i+1] = nums[i]
				}
			} else {
				fmt.Println("case final")
			}
			i = 0
		}
	}
	return count <= 1
}

// Use the approach locate and analyze problem index
func checkPossibilityLocateAnalyze(nums []int) bool {
	// locate the problem index
	wasPSet := false
	p := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if p >= 0 {
				// if there is a second problem index, it cannot work.
				return false
			}
			wasPSet = true
			p = i
		}
	}
	// correcting in one change is possible if:
	// no problem index or
	// or p is the first index or
	// p is the next to last index
	// or the adjacent numbers before and after are non-decreasing
	// or the number is not more than two numbers ahead
	return wasPSet == false || p == 0 || p == len(nums)-2 || nums[p-1] <= nums[p+1] || nums[p] <= nums[p+2]
}
