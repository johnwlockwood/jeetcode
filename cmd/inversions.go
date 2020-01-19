/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"

	"github.com/johnwlockwood/jeetcode/numbers/inversions"
	"github.com/spf13/cobra"
)

// inversionsCmd represents the inversions command
var inversionsCmd = &cobra.Command{
	Use:   "inversions",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cpuprofile, err := cmd.Flags().GetString("cpuprofile")
		if err != nil {
			log.Fatal(err)
		}
		if cpuprofile != "" {
			f, err := os.Create(cpuprofile)
			if err != nil {
				log.Fatal(err)
			}
			pprof.StartCPUProfile(f)
			defer pprof.StopCPUProfile()
		}
		fmt.Println("inversions called")

		file, err := os.Open(args[0])
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}

		nums := make([]int, 0)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			n, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			nums = append(nums, n)
		}

		invCount := inversions.CountInversions(nums)
		fmt.Printf("Found %d inversions.\n", invCount)

		memprofile, err := cmd.Flags().GetString("memprofile")
		if memprofile != "" {
			f, err := os.Create(memprofile)
			if err != nil {
				log.Fatal(err)
			}
			pprof.WriteHeapProfile(f)
			f.Close()
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(inversionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inversionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inversionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	inversionsCmd.Flags().String("cpuprofile", "", "write cpu profile to file. View profile with go tool pprof -web <profile name>")
	inversionsCmd.Flags().String("memprofile", "", "write memory profile to file. View profile with go tool pprof -web <profile name>")
}
