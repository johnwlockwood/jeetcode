package inversions

// CountInversions counts the number of inversions in the array
func CountInversions(nums []int) int {
	_, numInv := sortAndCountInv(nums)
	return numInv
}

func sortAndCountInv(nums []int) ([]int, int) {
	n := len(nums)
	if n == 0 || n == 1 {
		return nums, 0
	}
	left, leftInv := sortAndCountInv(nums[:n/2])
	right, rightInv := sortAndCountInv(nums[n/2:])
	sortedNums, splitInv := mergeAndCountSplitInv(left, right)
	return sortedNums, leftInv + rightInv + splitInv
}

func mergeAndCountSplitInv(left, right []int) ([]int, int) {
	i := 0
	j := 0
	splitInv := 0
	n := len(left) + len(right)
	sortedNums := make([]int, len(left)+len(right))
	for k := 0; k < n; k++ {
		// if i or j fall off the end of the left or right, copy the remaining elements over to sorted Nums
		switch {
		case i == len(left):
			sortedNums[k] = right[j]
			j++
		case j == len(right):
			sortedNums[k] = left[i]
			i++
		case left[i] < right[j]:
			sortedNums[k] = left[i]
			i++
		default:
			sortedNums[k] = right[j]
			j++
			splitInv += n/2 - i
		}
	}
	return sortedNums, splitInv
}
