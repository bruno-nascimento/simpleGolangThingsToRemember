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
	"github.com/spf13/cobra"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// readInputCmd represents the readInput command
var readInputCmd = &cobra.Command{
	Use:   "readInput",
	Short: "Read input from stdin",
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().Unix())
		target := rand.Intn(100)+1
		fmt.Println("## Guess a number between 1 and 100 ##")
		reader := bufio.NewReader(os.Stdin)
		loop:
		for {
			guessStr, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input : " + err.Error())
				return
			}
			guess, err := strconv.Atoi(strings.TrimSpace(guessStr))
			if err != nil {
				fmt.Println("Error converting input : " + err.Error())
				return
			}
			switch res := target - guess; {
			case res == 0:
				fmt.Println("You guessed it")
				break loop
				return
			case res > 0:
				fmt.Println("Your guess was LOW")
				continue loop
			case res < 0:
				fmt.Println("Your guess was HIGH")
				continue loop
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(readInputCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readInputCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readInputCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
