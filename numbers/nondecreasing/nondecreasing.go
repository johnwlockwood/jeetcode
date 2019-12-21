package nondecreasing

import "fmt"

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
				nums[i+1] = nums[i]
			} else {
				fmt.Println("case final")
			}
			i = 0
		}
	}
	return count <= 1
}
