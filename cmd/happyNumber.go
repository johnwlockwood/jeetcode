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
	"math"
	"reflect"

	"github.com/spf13/cobra"
)

// happyNumberCmd represents the happyNumber command
var happyNumberCmd = &cobra.Command{
	Use:   "happyNumber",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("happyNumber called")
		if got, want := getNext(19), 82; got != want {
			fmt.Printf("got %v, want %v\n", got, want)
		} else {
			fmt.Printf("success: got %v, want %v\n", got, want)
		}
		fmt.Println("isHappyApproach1")
		if got, want := isHappyApproach1(7), true; got != want {
			fmt.Printf("got %v, want true\n", got)
		} else {
			fmt.Printf("success: got %v, want %v\n", got, want)
		}
		if got, want := isHappyApproach1(0), false; got != want {
			fmt.Printf("got %v, want true\n", got)
		} else {
			fmt.Printf("success: got %v, want %v\n", got, want)
		}
		fmt.Println("isHappy")
		theDigits := digits(934935)
		if got, want := theDigits, []int{9, 1}; !reflect.DeepEqual(got, want) {
			fmt.Printf("got %v, want %v\n", got, want)
		} else {
			fmt.Printf("success: got %v, want %v\n", got, want)
		}
		if got, want := isHappy(934935), false; got != want {
			fmt.Printf("got %v, want true\n", got)
		} else {
			fmt.Printf("success: got %v, want %v\n", got, want)
		}
	},
}

func isHappyApproach1(n int) bool {
	// a couple of examples to get started
	// n = 7; 7**2 = 49; 4**2 + 9**2; 16+81 = 97; 81+49=130; 1+9+0=10; 1
	// n = 116
	seen := make(map[int]struct{})
	for n != 1 {
		if _, ok := seen[n]; ok {
			break
		}
		seen[n] = struct{}{}
		n = getNext(n)
		fmt.Printf("next number: %d\n", n)

	}
	return n == 1
}

// given a number, what is the next number
func getNext(n int) int {
	var totalSum int
	for n > 0 {
		d := n % 10
		// fmt.Printf("next digit of %d is %d\n", n, d)
		n = n / 10
		totalSum = totalSum + d*d
	}
	return totalSum
}

func isHappy(n int) bool {
	// The nth digit is (the remainder of dividing by 10n) divided by 10n-1
	if n < 0 {
		return false
	}
	seen := make(map[int]struct{})
	for n != 1 {
		sum := 0
		theDigits := digits(n)
		fmt.Println("digits of ", n, " are ", theDigits)
		squares := make([]int, 0, len(theDigits))
		for _, v := range theDigits {
			sq := int(math.Pow(float64(v), float64(2)))
			squares = append(squares, sq)
			sum = sum + sq
		}
		fmt.Println("sum of squares of ", theDigits, " is ", sum, " squares: ", squares)
		if _, ok := seen[sum]; ok {
			break
		}
		seen[sum] = struct{}{}
		n = sum
	}
	return n == 1
}

func digits(n int) []int {
	values := make([]int, 0)
	if n < 0 {
		return values
	}

	for i := 0; int(math.Pow10(i)) <= n; i++ {
		dig := (n % int(math.Pow10(i+1))) / int(math.Pow10(i))
		values = append(values, dig)
	}
	return values
}

func init() {
	rootCmd.AddCommand(happyNumberCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// happyNumberCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// happyNumberCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
