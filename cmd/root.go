package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var rootCmd = &cobra.Command{
	Use:   "AOC [year] [day]",
	Short: "Runs the Advent of Code solutions for the specified year and day",
	Long:  `Runs the Advent of Code solutions for the specified year and day`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		year := args[0]
		day := args[1]

		fmt.Printf("Running the Advent of Code solutions for the year %s and day %s\n", year, day)

		// Path is year/day (day starts with a zero)
		// Example: 2023/day01
		path := fmt.Sprintf("%s/day%02s", year, day)

		// Run the main_test.go file in the specified path
		// Example: go test -v 2023/day01/main_test.go

		testFilePath := fmt.Sprintf("%s/main_test.go", path)

		// Check if the path exists
		if _, err := os.Stat(testFilePath); os.IsNotExist(err) {
			fmt.Printf("The path %s does not exist\n", testFilePath)
			os.Exit(1)
		}

		// Run the tests
		_ = exec.Command("go", "test", "-v", testFilePath)

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
