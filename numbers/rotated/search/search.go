package search

// find the rotation point
func findRotation(rotated []int, left, right int) int {
	if right == left {
		return right
	}
	half := ((right + 1 - left) / 2) + left
	if rotated[half-1] > rotated[half] {
		// pivot is in the middle
		return half
	}
	if rotated[left] > rotated[right] {
		if rotated[left] > rotated[half] {
			// search left
			return findRotation(rotated, left, half-1)
		}
		// search right
		return findRotation(rotated, half, right)

	}
	return left
}

func binarySearch(nums []int, item, low, high int) int {
	if high <= low {
		if item == nums[low] {
			return low
		}
		return -1
	}
	mid := (low + high) / 2

	if item == nums[mid] {
		return mid
	}
	if item > nums[mid] {
		return binarySearch(nums, item, mid+1, high)
	}
	return binarySearch(nums, item, low, mid-1)
}

func search(nums []int, target int) int {
	// find rotation index with binary like search
	// find value with binary search
	pivot := findRotation(nums, 0, len(nums)-1)
	if nums[pivot] == target {
		return pivot
	} else if nums[pivot] > target && pivot > 0 {
		// search left of pivot
		return binarySearch(nums, target, 0, pivot-1)
	}
	// search right of pivot
	return binarySearch(nums, target, pivot, len(nums)-1)
}
