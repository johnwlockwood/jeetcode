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
	"sort"

	"github.com/spf13/cobra"
)

// majorityElementCmd represents the majorityElement command
var majorityElementCmd = &cobra.Command{
	Use:   "majorityElement",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		nums := []int{3, 2, 3}
		el := majorityElementBrute(nums)
		fmt.Println("majorityElementBrute ", el)
		el = majorityElementBoyerMoore(nums)
		fmt.Println("majorityElementBoyerMoore ", el)
		el = majorityElementSort(nums)
		fmt.Println("majorityElementSort ", el)
	},
}

func majorityElementSort(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElementBoyerMoore(nums []int) int {
	var count, candidate int

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func majorityElementBrute(nums []int) int {
	counts := make(map[int]int, len(nums))
	for _, v := range nums {
		if c, ok := counts[v]; ok {
			counts[v] = c + 1
		} else {
			counts[v] = 1
		}
	}
	var majorityElement, majorityCount int
	for k, v := range counts {
		if majorityCount < v {
			majorityCount = v
			majorityElement = k
		}
	}
	if majorityCount > len(nums)/2 {
		return majorityElement
	}
	return -1
}

func init() {
	rootCmd.AddCommand(majorityElementCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// majorityElementCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// majorityElementCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
