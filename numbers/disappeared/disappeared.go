package disappeared

import "math"

func findDisappearedNumbers(nums []int) []int {
	// the index is a place to store information
	// if the value is negative, it is marked as seen at that index
	// so any positive values indicate its index has not been marked
	// for each number, mark that number by negating the number at it's index
	for _, v := range nums {
		vi := int(math.Abs(float64(v))) - 1
		if nums[vi] > 0 {
			nums[vi] *= -1
		}
	}
	// use the original array to record the values
	// the value at an index already looked at doesn't need to be
	// kept, so free to overwrite it.
	// if the number in i is positive.
	// record it at the next available place
	var outI int
	for i, v := range nums {
		if v > 0 {
			nums[outI] = i + 1
			outI++
		}
	}
	// return only the part of the array with the disappeared values
	return nums[:outI]
}

func findDisappearedNumbersBitmask(nums []int) []int {
	// find missing values with a bit mask. only works if nums < 64
	// memory still grows O(n) because 1 bit per n
	d := make([]int, 0)
	var bitMask uint
	for _, v := range nums {
		bitMask |= (1 << uint(v))
	}
	for i := 1; i <= len(nums); i++ {
		if (bitMask & (1 << uint(i))) == 0 {
			d = append(d, i)
		}
	}
	return d
}
