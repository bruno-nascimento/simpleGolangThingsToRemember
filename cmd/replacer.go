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
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// replacerCmd represents the replacer command
var replacerCmd = &cobra.Command{
	Use:   "replacer string (old new)... ",
	Short: "string.Replacer",
	Long: `https://golang.org/pkg/strings/#Replacer
sample command:
	replacer 'This is <b>HTML</b>!' '<' '&lt;' ">" "&gt;"
outputs:
	This is &lt;b&gt;HTML&lt;/b&gt;!
`,
	Run: func(cmd *cobra.Command, args []string) {
		if (len(args) - 1) % 2 != 0 {
			fmt.Println("Invalid command arguments")
			fmt.Println("replacer usage: replacer stringToBeReplaced oldPattern newPattern oldPattern newPattern ...")
			return
		}
		r := strings.NewReplacer(args[1:]...)
		fmt.Println(r.Replace(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(replacerCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// replacerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// replacerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
