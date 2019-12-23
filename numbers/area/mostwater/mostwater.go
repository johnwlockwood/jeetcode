package mostwater

import "math"

// Runtime: 1328 ms, faster than 5.11% of Go online submissions for Container With Most Water.
// Memory Usage: 5.6 MB, less than 46.67% of Go online submissions for Container With Most Water.

func maxArea(height []int) int {
	// for each value from the right to the left, calc area = min(ni, nj)*(j-i)
	// trac max h=min(ni, nj) // optimization for later
	// track max area
	// if hj is less than max h, skip // optimization for later
	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j := len(height) - 1; j > i; j-- {
			w := j - i

			h := int(math.Min(float64(height[i]), float64(height[j])))
			area := h * w
			if area > max {
				max = area
			}
		}
	}
	return max
}
