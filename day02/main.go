package main

import (
	"fmt"
    "strings"
    "strconv"
	"aoc_2025/utils"
)

func main() {
	filePath := "input.txt"
	fmt.Printf("Reading input from: %s\n", filePath)
	dataLines := utils.ReadInputLines(filePath)

	resultP1 := solvePart1(dataLines)

	resultP2 := solvePart2(dataLines)

	fmt.Println("---")
	fmt.Printf("✨ Part 1 Result: %d\n", resultP1)
	fmt.Printf("🌟 Part 2 Result: %d\n", resultP2)
}

func solvePart1(lines []string) int {
    // exmple line input 11-22,95-115,
    parts := strings.Split(lines[0], ",")
    count := 0
    for _, part := range parts {
        rangeParts := strings.Split(part, "-")
        rangeStart, _ := strconv.Atoi(rangeParts[0])
        rangeEnd, _ := strconv.Atoi(rangeParts[1])

        for number := rangeStart; number <= rangeEnd; number++ {
            if (isRepeatedTwice(strconv.Itoa(number))) {
                count += number
            }
        }
    }
    return count
}

func solvePart2(lines []string) int {
    // exmple line input 11-22,95-115,
    parts := strings.Split(lines[0], ",")
    count := 0
    for _, part := range parts {
        rangeParts := strings.Split(part, "-")
        rangeStart, _ := strconv.Atoi(rangeParts[0])
        rangeEnd, _ := strconv.Atoi(rangeParts[1])

        for number := rangeStart; number <= rangeEnd; number++ {
            if (isRepeatedAtLeastTwice(strconv.Itoa(number))) {
                count += number
            }
        }
    }
    return count
}


func isRepeatedTwice(id string) bool {
    length := len(id)
    if (length % 2 == 1) {
        return false
    }
    halfLength := length / 2
    // fmt.Printf("id: %s, length: %d, %s, %s\n", id, length, id[0:halfLength], id[halfLength:])
    return id[0:halfLength] == id[halfLength:]
}


func isRepeatedAtLeastTwice(id string) bool {
    length := len(id)
    halfLength := length / 2

    for i := 1; i <= halfLength; i ++ {
        if (length % i != 0) {
            continue
        }

        chunks := splitStringIntoChunks(id, i)
        if(areAllChunksTheSame(chunks)) {
            return true
        }
    }

    return false
}

func splitStringIntoChunks(s string, chunkSize int) []string {
	runes := []rune(s)

	var chunks []string

	for i := 0; i < len(runes); i += chunkSize {
		end := i + chunkSize

		if end > len(runes) {
			end = len(runes)
		}

		chunk := runes[i:end]
		chunks = append(chunks, string(chunk))
	}

	return chunks
}

func areAllChunksTheSame(chunks []string) bool {
	ref := chunks[0]

	for i := 1; i < len(chunks); i++ {
		if chunks[i] != ref {
			return false
		}
	}

	return true
}
