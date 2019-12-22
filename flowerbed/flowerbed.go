package flowerbed

import "fmt"

// easy problem

// w/o printing:
// Runtime: 12 ms, faster than 99.05% of Go online submissions for Can Place Flowers.
// Memory Usage: 5.9 MB, less than 100.00% of Go online submissions for Can Place Flowers.

func canPlaceFlowers(flowerbed []int, n int) bool {
	// mark existing flowers as -1
	// plant flowers in even spots which are >0
	// for every -1 spot, set the surrounding spots to 0 if they are in the flowerbed
	// plant the last plot if the final two are empty
	// count positive values
	fmt.Printf("original flowerbed \t\t\t%v \t%d free?\n", flowerbed, n)
	freePlots := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			flowerbed[i] = -1
		}
	}

	fmt.Printf("\tmarked flowerbed \t\t%v\n", flowerbed)
	for i := 0; i < len(flowerbed); i++ {
		if i%2 == 0 && flowerbed[i] >= 0 {
			flowerbed[i] = 1
		}
	}
	fmt.Printf("\tproposed dirty flowerbed \t%v\n", flowerbed)

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == -1 {
			if i < len(flowerbed)-1 && flowerbed[i+1] == 1 {
				flowerbed[i+1] = 0
			}
			if i > 0 && flowerbed[i-1] == 1 {
				flowerbed[i-1] = 0
			}
		}
	}
	fmt.Printf("\tcleaned flowerbed \t\t%v\n", flowerbed)

	// try to plant the last plot
	if len(flowerbed) > 1 {
		if flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
			flowerbed[len(flowerbed)-1] = 1
		}
	}
	fmt.Printf("\tlast plot planted flowerbed \t%v\n", flowerbed)

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] > 0 {
			freePlots++
		}
	}
	fmt.Printf("\t\t\t\t%d free\n", freePlots)
	return n <= freePlots
}

// Do it in a single pass
func canPlaceFlowersSinglePass(flowerbed []int, n int) bool {
	freePlots := 0
	lf := len(flowerbed)
	for i := 0; i < lf; i++ {
		prevNoBlock := (i == 0 || flowerbed[i-1] == 0)
		nextNoBlock := (i == lf-1 || flowerbed[i+1] == 0)
		currEmpty := flowerbed[i] == 0
		if currEmpty && prevNoBlock && nextNoBlock {
			flowerbed[i] = 1
			freePlots++
		}
	}
	return freePlots >= n
}

// Do it in a single pass, no changes to input
func canPlaceFlowersNoChanges(flowerbed []int, n int) bool {
	freePlots := 0
	lf := len(flowerbed)
	lp := 0
	for i := 0; i < lf; i++ {
		prevNoBlock := (i == 0 || (flowerbed[i-1] == 0 && lp != i-1) || lp < i-1)
		nextNoBlock := (i == lf-1 || flowerbed[i+1] == 0)
		currEmpty := flowerbed[i] == 0
		if currEmpty && prevNoBlock && nextNoBlock {
			lp = i
			freePlots++
		} else if !currEmpty {
			lp = i
		}

	}
	return freePlots >= n
}
