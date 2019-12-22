package diagonal

import "fmt"

// easy problem

// Given a matrix of numbers
// return a list of the diagonals
//  nums [
//  [ 1, 2, 3 ],
//  [ 4, 5, 6 ],
//  [ 7, 8, 9 ],
//  ]
// needs to return
// [
// [1],
// [2, 4],
// [3, 5, 7],
// [6, 8],
// [9]
// [
// nums[i][j], increment i, decrement j

func getDiagonals(nums [][]int) ([][]int, error) {
	m := len(nums)
	if m == 0 {
		return [][]int{}, nil
	}
	n := len(nums[0])
	for i := 1; i < m; i++ {
		if n != len(nums[i]) {
			return nil, fmt.Errorf("Matrix is not consistent")
		}
	}
	diagonalList := make([][]int, 0)
	// for each diagonal starting point i,j=0,0

	r := 0

	for c := 0; c < n; c++ {
		j := c
		dl := make([]int, 0)
		for i := r; i < m; i++ {
			dl = append(dl, nums[i][j])
			j--
			if j < 0 {
				break
			}
		}
		diagonalList = append(diagonalList, dl)
	}
	r++
	c := n - 1
	for r < m {
		j := c
		dl := make([]int, 0)
		for i := r; i < m; i++ {
			dl = append(dl, nums[i][j])
			j--
			if j < 0 {
				break
			}
		}
		diagonalList = append(diagonalList, dl)
		r++
	}
	return diagonalList, nil
}
