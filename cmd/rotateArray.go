/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rotateArrayCmd represents the rotateArray command
var rotateArrayCmd = &cobra.Command{
	Use:   "rotateArray",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		nums, shift, expected := []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}
		rotateWithCopy(nums, shift)
		fmt.Println("rotateArray called", nums, expected)
		nums, shift, expected = []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}
		nums, shift, expected = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
		nums, shift = []int{1, 2}, 1
		nums, shift = []int{1}, 0
		fmt.Println("constant attempt with ", nums)
		rotateConstantMem(nums, shift)

		fmt.Println("rotateConstantMem shift ", shift, " called", nums, expected)
	},
}

func rotateWithCopy(nums []int, k int) {
	// with memory O(n)
	alt := make([]int, len(nums))
	copy(alt, nums)
	for i := 0; i < len(nums); i++ {
		r := (i + k) % len(nums)
		nums[r] = alt[i]
	}
}

func rotateConstantMem(nums []int, k int) {
	// with memory O(1)
	// swapping
	// get k mod of array in case k > len(nums)
	if k <= 0 {
		return
	}
	k = k % len(nums)
	n := len(nums)
	fmt.Println("loop til", n-k)
	l := 0
	for l < n-k {
		for r := n - k; r < n || r == l; r++ {
			nums[l], nums[r] = nums[r], nums[l]
			fmt.Println("l: ", l, " r: ", r, " nums: ", nums)
			l++
		}
	}
	// if there is an overlap
	for k*2 > n {
		overlap := k*2 - n
		fmt.Println("overlap ", overlap)
		// shift by overlap
		k = overlap
		for l < n-k {
			for r := n - k; r < n || r == l; r++ {
				nums[l], nums[r] = nums[r], nums[l]
				fmt.Println("l: ", l, " r: ", r, " nums: ", nums)
				l++
			}
		}
	}
	// reverse last part

	// by 2
	// [1,2,3,4,5,6,7]
	//  ^ swap	  ^
	// [6,2,3,4,5,1,7]
	//    ^ swap    ^
	// [6,7,3,4,5,1,2]
	//      ^     ^
	// [6,7,1,4,5,3,2]
	//        ^     ^
	// [6,7,1,2,5,3,4]
	//          ^ ^
	// [6,7,1,2,3,5,4]
	//            ^ ^
	// [6,7,1,2,3,4,5]
	//              ^

	// by 4
	// [1,2,3,4,5,6,7]
	//  ^     ^
	// [4,2,3,1,5,6,7]
	//    ^     ^
	// [4,5,3,1,2,6,7]
	//      ^     ^
	// [4,5,6,1,2,3,7]
	//        ^     ^
	// now reverse the overlap
	// [4,5,6,7,2,3,1]
	//          ^   ^
	// [4,5,6,7,1,3,2]
	//            ^ ^
	// [4,5,6,7,1,2,3]

}

func init() {
	rootCmd.AddCommand(rotateArrayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rotateArrayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rotateArrayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
