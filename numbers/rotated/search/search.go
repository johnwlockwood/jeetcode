package search

// find the rotation point
func findRotation(rotated []int, left, right int) int {
	if right < left {
		return -1
	}
	if right == left {
		return right
	}
	half := ((right + 1 - left) / 2) + left
	if rotated[left] > rotated[right] {
		if right-left == 2 {
			return right
		}
		if rotated[left] > rotated[half] {
			if half-left == 1 {
				return half
			}
			// search left
			return findRotation(rotated, left, half-1)
		}
		// search right
		return findRotation(rotated, half, right)

	}
	return left
}
func search(nums []int, target int) int {
	// find rotation index with binary like search
	// find value with binary search

	return -1
}
