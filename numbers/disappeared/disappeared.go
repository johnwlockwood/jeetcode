package disappeared

func findDisappearedNumbers(nums []int) []int {
	// the index is a place to store information
	// if the value is negative, it is marked as seen at that index
	// so any positive values indicate its index has not been marked
	d := make([]int, 0)
	// for each number, mark that number by negating the number at it's index
	for _, v := range nums {
		if v < 0 {
			if nums[-1*v-1] > 0 {
				nums[-1*v-1] *= -1
			}
		} else {
			if nums[v-1] > 0 {
				nums[v-1] *= -1
			}
		}
	}
	// if the number in i is positive, add i+1 to d
	for i, v := range nums {
		if v > 0 {
			d = append(d, i+1)
		}
	}
	return d
}
