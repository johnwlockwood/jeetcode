package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	// merge
	var loc int
	for i := 0; i < n; i++ {
		item := nums2[i]
		high := m + i - 1
		if m == 0 && high < 0 {
			loc = 0
		} else {
			loc = binarySearch(nums1, item, 0, high)
		}
		for j := high; j >= loc; j-- {
			nums1[j+1] = nums1[j]
		}
		nums1[loc] = nums2[i]
	}
}

func binarySearch(nums []int, item, low, high int) int {
	if high <= low {
		if item > nums[low] {
			return low + 1
		}
		return low
	}
	mid := (low + high) / 2

	if item == nums[mid] {
		return mid + 1
	}
	if item > nums[mid] {
		return binarySearch(nums, item, mid+1, high)
	}
	return binarySearch(nums, item, low, mid-1)
}

// merge solution 3
func mergeSolution3(nums1 []int, m int, nums2 []int, n int) {
	// 2 pointers
	var p1, p2, p int
	p1 = m - 1
	p2 = n - 1
	// set the destination pointer in nums1
	p = m + n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] < nums2[p2] {
			nums1[p] = nums2[p2]
			p2--
		} else {
			nums1[p] = nums1[p1]
			p1--
		}
		p--
	}
	copy(nums1[:p2+1], nums2[:p2+1])
}

func main() {
	fmt.Printf("%v\n", binarySearch([]int{3, 8, 9}, 7, 0, 3))
}
