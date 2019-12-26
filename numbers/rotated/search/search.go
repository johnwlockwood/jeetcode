package search

// find the rotation point
func findRotation(nums []int, left, right int) int {
	for left < right {
		mid := (left + 1 + right) / 2
		if nums[mid-1] > nums[mid] {
			return mid
		}
		if nums[mid] > nums[right] {
			// search right half
			left = mid + 1
		} else {
			// search left half
			right = mid - 1
		}
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

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Search in Rotated Sorted Array.
// Memory Usage: 2.6 MB, less than 50.00% of Go online submissions for Search in Rotated Sorted Array.
// https://leetcode.com/submissions/detail/288307651/
func search(nums []int, target int) int {
	// find rotation index with binary like search
	// find value with binary search
	if len(nums) <= 0 {
		return -1
	}
	pivot := 0
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + 1 + right) / 2
		if nums[mid-1] > nums[mid] {
			pivot = mid
			break
		}
		if nums[mid] > nums[right] {
			// search right half
			left = mid + 1
		} else {
			// search left half
			right = mid - 1
		}
	}
	if left >= right {
		pivot = left
	}
	if nums[pivot] == target {
		return pivot
	} else if pivot > 0 && nums[pivot-1] >= target && target > nums[len(nums)-1] {
		// search left of pivot
		return binarySearch(nums, target, 0, pivot)
	}
	// search right of pivot
	return binarySearch(nums, target, pivot, len(nums)-1)
}
