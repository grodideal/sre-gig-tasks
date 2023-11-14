package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocat [file]",
	Short: "A simple 'cat' command in Go",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		displayFile(filePath, numberLines, numberNonEmptyLines)
	},
}

var numberLines, numberNonEmptyLines bool

func init() {
	rootCmd.Flags().BoolVarP(&numberLines, "number", "n", false, "Number all output lines")
	rootCmd.Flags().BoolVarP(&numberNonEmptyLines, "number-nonempty", "b", false, "Number nonempty output lines")
}

func displayFile(filePath string, numberLines, numberNonEmptyLines bool) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()

		if numberLines || (numberNonEmptyLines && line != "") {
			fmt.Printf("%6d  ", lineNumber)
			lineNumber++
		}

		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
