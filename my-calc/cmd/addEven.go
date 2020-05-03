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
	"strconv"
	"github.com/spf13/cobra"
)

// addEvenCmd represents the addEven command
var addEvenCmd = &cobra.Command{
	Use:   "addEven",
	Short: "Adding all even numbers in the list",
	Long: `This command identifies all the even numbers in the list and sums them up`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("The sum of even numbers in the arguments: ", args, " is: ", addEven(args))
	},
}

func addEven(args []string) int {
	var sum int
	for _, value := range args {
		number, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println(err)
		}

		if number % 2 == 0 {
			sum += number
		}
	}
	return sum
}


func init() {
	rootCmd.AddCommand(addEvenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addEvenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addEvenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
