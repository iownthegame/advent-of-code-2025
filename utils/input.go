// aoc_2025/utils/input.go
package utils

import (
	"bufio"
	"fmt"
	"os"
)

// ReadInputLines opens the specified file and returns its content as a slice of strings.
// It uses a capitalized name to be exported from the 'utils' package.
func ReadInputLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file %s: %v", filePath, err))
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading file %s: %v", filePath, err))
	}

	return lines
}
