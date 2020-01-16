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
	"archive/zip"
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"

	"github.com/johnwlockwood/jeetcode/graphs/course/connected"
	"github.com/spf13/cobra"
)

// connectedCmd represents the connected command
var connectedCmd = &cobra.Command{
	Use:   "connected",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
		tests := map[string]string{
			"SCC.txt": "434821,968,459,313,211",
		}

		r, err := zip.OpenReader("graphs/course/connected/testdata/SCC.txt.zip")
		if err != nil {
			panic(err)
		}
		defer r.Close()
		for _, f := range r.File {
			want, ok := tests[f.Name]
			if !ok {
				continue
			}
			fmt.Printf("Contents of %s:\n", f.Name)
			rc, err := f.Open()
			defer rc.Close()
			if err != nil {
				panic(err)
			}

			edges := make([][]int, 0)

			scanner := bufio.NewScanner(rc)
			i := 0
			for scanner.Scan() {
				line := scanner.Text()
				edgeStr := strings.Split(line, " ")
				if len(edgeStr) < 2 {
					panic(fmt.Sprintf("bad data %s", line))
				}
				tail, err := strconv.Atoi(edgeStr[0])
				if err != nil {
					panic(err)
				}
				head, err := strconv.Atoi(edgeStr[1])
				if err != nil {
					panic(err)
				}

				edges = append(edges, []int{tail, head})
				i++
			}
			fmt.Printf("there are %d edges and i == %d\n", len(edges), i)
			got := connected.FiveLargestSCCs(edges)
			fmt.Println(got)
			if got != want {
				fmt.Printf("for %s: got %v , want %v\n", f.Name, got, want)
			}
		}

		fmt.Println("connected called")
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
	rootCmd.AddCommand(connectedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	connectedCmd.Flags().String("cpuprofile", "", "write cpu profile to file. View profile with go tool pprof -web <profile name>")
	connectedCmd.Flags().String("memprofile", "", "write memory profile to file. View profile with go tool pprof -web <profile name>")
}
