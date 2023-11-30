package aoc

import (
	"aoc/pkg/setup"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:     "setup",
	Short:   "Setup Advent of Code challenges",
	Long:    "Takes input of the form `year/day` and creates directories and files for that challenge.\ne.g. `setup 2022/1` or `set `22/1`.\n\nSets up the entire year if only the year is provided\ne.g. `setup 2022` will download and setup workspace for all problems from 2022.",
	Aliases: []string{"set","stp"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		setup.SetupProject(args[0])
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
