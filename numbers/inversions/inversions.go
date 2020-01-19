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
			// if the number on the left is greater than the one on the right,
			// then number of inversions are the number of remaining items
			// on the left.  [1, 2, 4] 		[3, 5, 6]
			//						i=2		 j=0
			// so, 3-2 = 1
			// example #2
			// left: [14108 54044 79294], i: 1, right [25260 29649 60660], j: 1
			//				i=1							   	 j=1
			// 3-1 = 2
			//				i=1							   	 	   j=2
			// 3-1 = 2
			//					  i=2							   j=2
			// 3-2 = 1
			// totaling 5
			splitInv += len(left) - i
		}
	}
	return sortedNums, splitInv
}
