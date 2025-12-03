package main

import (
	"fmt"
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
    // exmple line input L51, R3
    startPos := 50
    countZero := 0
    for i, line := range lines {
        rotation := line[0]
        distance, _ := strconv.Atoi(line[1:])
        fmt.Printf("Line %d: rotation: %c, distance: %d\n", i, rotation, distance)

        switch rotation {
        case 'R':
            startPos += distance
        case 'L':
            startPos -= distance
        }
        startPos = (startPos + 100) % 100
        if (startPos == 0) {
            countZero += 1
        }
    }
	return countZero
}

func solvePart2(lines []string) int {
    // exmple line input L51, R3
    startPos := 50
    countZero := 0
    for i, line := range lines {
        rotation := line[0]
        distance, _ := strconv.Atoi(line[1:])
        fmt.Printf("Line %d: rotation: %c, distance: %d, startPos: %d\n", i, rotation, distance, startPos)

        countZero += distance / 100
        distance %= 100

        newPos := 0
        switch rotation {
        case 'R':
            newPos = (startPos + distance + 100) % 100
            if (newPos == 0 || (newPos < startPos && startPos != 0)) {
                fmt.Printf("\tnewPos %d < startPos %d, distance %d\n", newPos, startPos, distance)
                countZero += 1
            }
        case 'L':
            newPos = (startPos - distance + 100) % 100
            if (newPos == 0 || (newPos > startPos && startPos != 0)) {
                fmt.Printf("\tnewPos %d > startPos %d, distance %d\n", newPos, startPos, distance)
                countZero += 1
            }
        }
        startPos = newPos
    }
	return countZero
}
