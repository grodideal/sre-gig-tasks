package main

import (
	"bytes"
	"os"
	"testing"
)

func TestDisplayFile(t *testing.T) {
	tests := []struct {
		fileContent      string
		numberLines      bool
		numberNonEmpty   bool
		expectedOutput   string
		expectedExitCode int
	}{
		{
			fileContent:      "line1\nline2\nline3",
			numberLines:      true,
			expectedOutput:   "     1  line1\n     2  line2\n     3  line3\n",
			expectedExitCode: 0,
		},
		{
			fileContent:      "line1\n\nline3",
			numberNonEmpty:   true,
			expectedOutput:   "     1  line1\n          \n     2  line3\n",
			expectedExitCode: 0,
		},
		{
			fileContent:      "invalid file",
			expectedExitCode: 1,
		},
	}

	for _, test := range tests {
		// Create a temporary file
		file, err := os.CreateTemp("", "testfile*.txt")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(file.Name())
		defer file.Close()

		// Write content to the file
		_, err = file.WriteString(test.fileContent)
		if err != nil {
			t.Fatal(err)
		}

		// Capture stdout
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Run the function to be tested
		displayFile(file.Name(), test.numberLines, test.numberNonEmpty)

		// Close the write end of the pipe and restore stdout
		w.Close()
		os.Stdout = oldStdout

		// Read captured output from the read end of the pipe
		var buf bytes.Buffer
		buf.ReadFrom(r)

		// Check if the output matches the expected output
		if buf.String() != test.expectedOutput {
			t.Errorf("Expected: %s\nGot: %s", test.expectedOutput, buf.String())
		}

		// Check if the exit code matches the expected exit code
		exitCode := 0
		if err != nil {
			exitCode = 1
		}
		if exitCode != test.expectedExitCode {
			t.Errorf("Expected exit code: %d\nGot exit code: %d", test.expectedExitCode, exitCode)
		}
	}
}
