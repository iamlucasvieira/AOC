package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var ToTest bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&ToTest, "test", "t", false, "Run tests")
}

var rootCmd = &cobra.Command{
	Use:   "AOC [year] [day]",
	Short: "Runs the Advent of Code solutions for the specified year and day",
	Long:  `Runs the Advent of Code solutions for the specified year and day`,
	Args:  cobra.ExactArgs(2),
	// Set flag if it is to run tests

	Run: func(cmd *cobra.Command, args []string) {
		year := args[0]
		day := args[1]

		// Path is year/day (day starts with a zero)
		// Example: 2023/day01
		path := fmt.Sprintf("./%s/day%02s", year, day)

		// Get File Path
		filePath := fmt.Sprintf("%s/main.go", path)

		// Check if filePath and testFilePath exist
		for _, file := range []string{path, filePath} {
			if _, err := os.Stat(file); os.IsNotExist(err) {
				fmt.Printf("File %s does not exist\n", file)
				os.Exit(1)
			}
		}

		fmt.Printf("Running the Advent of Code solutions for the year %s and day %s\n", year, day)

		// Run the file
		fmt.Printf("> go run %s\n", filePath)
		executeCommand("go", "run", filePath)

		if ToTest {
			// Run the test file
			fmt.Printf("> go test -v %s\n", path)
			executeCommand("go", "test", "-v", path)
		}
	},
}

func executeCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing %s: %v\nOutput: %s", command, err, output)
	}
	fmt.Printf("%s\n", output)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
