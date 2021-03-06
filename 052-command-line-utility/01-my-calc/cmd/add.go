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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adding all numbers in the list",
	Long: `This command adds all the numbers in the list`,
	Run: func(cmd *cobra.Command, args []string) {
		flagStatus, _ := cmd.Flags().GetBool("float")
		if flagStatus {
			fmt.Println("The sum of the arguments: ", args, " is: ", addFloat(args))
		} else {
			fmt.Println("The sum of the arguments: ", args, " is: ", addInt(args))
		}
	},
}

func addInt(args []string) int {
	var sum int
	for _, value := range args {
		floatValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println(err)
		}
		sum += floatValue
	}
	return sum
}

func addFloat(args []string) float64 {
	var sum float64
	for _, value := range args {
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println(err)
		}

		sum += number
	}

	return sum
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().BoolP("float", "f", false, "Adds float numbers")
}
