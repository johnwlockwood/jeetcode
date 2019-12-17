package permutations

import "fmt"

func permute(nums []int) [][]int {
	// some examples of permutations of [ 1, 2, 3]
	// try the first number in every position
	// [1, 2, 3], [2, 1, 3], [3, 2, 1]
	// [2, 1, 3], [1, 2, 3], [1, 3, 2]
	// I ended up finding the Heap algorithm
	// this is the non-recursive one
	n := len(nums)
	pCount := fac(n) // len(nums) factorial
	perms := make([][]int, pCount)
	for i := range perms {
		perms[i] = make([]int, n)
	}

	// c is the stack state
	c := make([]int, n)

	copy(perms[0][:], nums)
	pi := 1
	i := 0
	fmt.Printf("at i: %d, nums %v perms %v, c %v\n", i, nums, perms, c)
	for i < n {
		if c[i] < i {
			if i%2 == 0 { // if even swap
				nums[0], nums[i] = nums[i], nums[0]
			} else {
				nums[c[i]], nums[i] = nums[i], nums[c[i]]
			}
			copy(perms[pi], nums)
			pi++
			c[i]++
			i = 0 // reset i
		} else {
			c[i] = 0
			i++
		}
		fmt.Printf("at i: %d, nums %v perms %v, c %v\n", i, nums, perms, c)
	}

	// for every number, swap with another

	return perms
}

func fac(n int) int {
	count := 1
	for i := 1; i <= n; i++ {
		count = count * i
	}
	return count
}

func permute2(nums []int) [][]int {
	n := len(nums)
	perms := make([][]int, fac(n))
	for i := range perms {
		perms[i] = make([]int, n)
	}

	// c is the stack state
	c := make([]int, n)

	copy(perms[0], nums)
	pi := 1 // permutation starts at 1 because the original is at the 0 location
	i := 0
	for i < n {
		if c[i] < i {
			if i%2 == 0 { //is even
				// swap first and i
				nums[0], nums[i] = nums[i], nums[0]
			} else {
				// swap value from stack state i with i
				nums[c[i]], nums[i] = nums[i], nums[c[i]]
			}
			// output
			copy(perms[pi], nums)
			pi++
			c[i]++
			i = 0 //
		} else {
			c[i] = 0
			i++
		}
	}
	return perms
}

func permGenerate(k int, nums []int) [][]int {
	// Heap's algorithm recursive

	if k == 1 {
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		return [][]int{numsCopy}
	}
	perms := make([][]int, 0)
	subPerms := permGenerate(k-1, nums)
	for _, perm := range subPerms {
		perms = append(perms, perm)
	}

	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			nums[i], nums[k-1] = nums[k-1], nums[i]
		} else {
			nums[0], nums[k-1] = nums[k-1], nums[0]
		}
		subPerms = permGenerate(k-1, nums)
		for _, perm := range subPerms {
			perms = append(perms, perm)
		}
	}
	return perms
}

func permuteRecursive(nums []int) [][]int {
	return permGenerate(len(nums), nums)
}
