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

// linkedListCycleCmd represents the linkedListCycle command
var linkedListCycleCmd = &cobra.Command{
	Use:   "linkedListCycle",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("linkedListCycle called")
		wat := &ListNode{
			Val: 3,
		}
		head := &ListNode{
			Val:  1,
			Next: wat,
		}
		wat.Next = head
		fmt.Printf("%v\n", hasCycle(head))
	},
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// Runtime: 200 ms, faster than 11.68% of Go online submissions for Linked List Cycle.
// Memory Usage: 3.8 MB, less than 71.43% of Go online submissions for Linked List Cycle.

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	currentNode := head
	posCurrentNode := 0
	for currentNode.Next != nil {
		if posCurrentNode > 0 {
			isCycleNode := head
			for isCyclePos := 0; isCyclePos < posCurrentNode; isCyclePos++ {
				if isCycleNode == currentNode {
					return true
				}
				isCycleNode = isCycleNode.Next
			}
		}
		posCurrentNode++
		currentNode = currentNode.Next
	}
	return false
}

func init() {
	rootCmd.AddCommand(linkedListCycleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// linkedListCycleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// linkedListCycleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
