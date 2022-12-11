/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"adventofcode/mod/2022/pkg/solutions"
	"fmt"

	"github.com/spf13/cobra"
)

type RunOptions struct {
	Day    int
	Puzzle int
}

var runOptions = RunOptions{}

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var solution solutions.Solution
		switch day := runOptions.Day; day {
		case 1:
			solution = solutions.Day1{Puzzle: runOptions.Puzzle}
		case 2:
			solution = solutions.Day2Init(runOptions.Puzzle)
		case 3:
			solution = solutions.Day3Init(runOptions.Puzzle)
		}
		fmt.Println(solution.Run())
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	runCmd.Flags().IntVarP(&runOptions.Day, "day", "d", 1, "Which day are you running")
	runCmd.Flags().IntVarP(&runOptions.Puzzle, "puzzle", "p", 1, "Which puzzle are you solving")
}
