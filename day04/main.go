package main

import (
	"fmt"
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
    // exmple line input ..@@.@@@@.
    grid := make([][]rune, len(lines))
    for r, line := range lines {
        runeRow := []rune(line)
        grid[r] = runeRow
    }

    count := 0
    for r, row := range grid {
        for c, char := range row {
            if char == '@' {
                if canAccess(grid, r, c) {
                    count += 1
                }
            }
        }
    }
    return count
}

func solvePart2(lines []string) int {
    // exmple line input ..@@.@@@@.
    grid := make([][]rune, len(lines))
    for r, line := range lines {
        runeRow := []rune(line)
        grid[r] = runeRow
    }

    count := 0

    for {
        rollsOfPaperToBeRemoved := findAccessibleRollsOfPaper(grid)

        if len(rollsOfPaperToBeRemoved) == 0 {
            break
        }
        count += len(rollsOfPaperToBeRemoved)
        removeRollsOfPaper(grid, rollsOfPaperToBeRemoved)
    }

    return count
}

type Coord struct {
	R int
	C int
}

func canAccess(grid [][]rune, r int, c int) bool {
    // fewer than four rolls of paper in the eight adjacent positions
    dirs := []Coord{
        {-1, -1},
        {-1, 0},
        {-1, 1},
        {0, -1},
        {0, 1},
        {1, -1},
        {1, 0},
        {1, 1},
    }

    countRollPaper := 0
    lenRow, lenCol := len(grid), len(grid[0])

    for _, dir := range dirs {
        new_r := r + dir.R
        new_c := c + dir.C
        if new_r >= 0 && new_r < lenRow && new_c >= 0 && new_c < lenCol && grid[new_r][new_c] == '@' {
            countRollPaper += 1
        }
    }
    return countRollPaper < 4
}

func findAccessibleRollsOfPaper(grid [][]rune) []Coord {
    accessibleRollsOfPaper := []Coord{}

    for r, row := range grid {
        for c, char := range row {
            if char == '@' {
                if canAccess(grid, r, c) {
                    accessibleRollsOfPaper = append(accessibleRollsOfPaper, Coord{r, c})
                }
            }
        }
    }

    return accessibleRollsOfPaper
}

func removeRollsOfPaper(grid [][]rune, rollsOfPaperCoords []Coord) {
    for _, coord := range rollsOfPaperCoords {
        grid[coord.R][coord.C] = 'x'
    }
}
