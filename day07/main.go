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

type Coord struct {
	R int
	C int
}
type CoordSet map[Coord]struct{}

func solvePart1(lines []string) int {
    // exmple line input ".......S.......", "...............", ".......^......."

    grid := make([][]rune, len(lines))
    for r, line := range lines {
        runeRow := []rune(line)
        grid[r] = runeRow
    }

    endLineIndex := len(lines) - 1
    entry := findEntry(grid)
    queue := []Coord{entry}
    visited := make(CoordSet)
    visited[entry] = struct{}{}
    count := 0

    checkAndEnqueue := func(newCoord Coord) {
        if _, ok := visited[newCoord]; !ok {
            queue = append(queue, newCoord)
            visited[newCoord] = struct{}{}
        }
    }

    for len(queue) > 0 {
        coord := queue[0]
        queue = queue[1:]

        if coord.R == endLineIndex {
            continue
        }

        if grid[coord.R + 1][coord.C] == '^' {
            // splitter, add left and right
            count += 1

            if coord.C > 0 {
                checkAndEnqueue(Coord{coord.R + 1, coord.C - 1})
            }
            if coord.C < len(lines[0]) - 1 {
                checkAndEnqueue(Coord{coord.R + 1, coord.C + 1})
            }
        } else {
            // continue going down
            checkAndEnqueue(Coord{coord.R + 1, coord.C})
        }
    }

    return count
}

func findEntry(grid [][]rune) Coord{
    for r, row := range grid {
        for c, char := range row {
            if char == 'S' {
                return Coord{r, c}
            }
        }
    }
    return Coord{-1, -1}
}


func solvePart2TLE(lines []string) int {
    // exmple line input ".......S.......", "...............", ".......^......."

    grid := make([][]rune, len(lines))
    for r, line := range lines {
        runeRow := []rune(line)
        grid[r] = runeRow
    }

    endLineIndex := len(lines) - 1
    entry := findEntry(grid)
    queue := []Coord{entry}
    count := 0

    for len(queue) > 0 {
        coord := queue[0]
        queue = queue[1:]
        fmt.Printf("queue size: %d, R: %d\n", len(queue), coord.R)

        if coord.R == endLineIndex {
            count += 1
            continue
        }

        if grid[coord.R + 1][coord.C] == '^' {
            // splitter, add left and right

            if coord.C > 0 {
                queue = append(queue, Coord{coord.R + 1, coord.C - 1})
            }
            if coord.C < len(lines[0]) - 1 {
                queue = append(queue, Coord{coord.R + 1, coord.C + 1})
            }
        } else {
            // continue going down
            queue = append(queue, Coord{coord.R + 1, coord.C})
        }
    }

    return count
}

type VisitedCounter map[Coord]int

func solvePart2(lines []string) int {
    // exmple line input ".......S.......", "...............", ".......^......."

    grid := make([][]rune, len(lines))
    for r, line := range lines {
        runeRow := []rune(line)
        grid[r] = runeRow
    }

    endLineIndex := len(lines) - 1
    entry := findEntry(grid)
    queue := []Coord{entry}
    visitedCounter := make(VisitedCounter)
    visitedCounter[entry] = 1
    count := 0

    checkAndEnqueue := func(newR int, newC int, currentVisitCount int) {
        key := Coord{R: newR, C: newC}
        currentMapCount, exists := visitedCounter[key]

        if !exists {
            queue = append(queue, Coord{newR, newC})
            visitedCounter[key] = currentVisitCount

        } else {
            visitedCounter[key] = currentMapCount + currentVisitCount
        }
    }

    for len(queue) > 0 {
        coord := queue[0]
        // fmt.Printf("queue size: %d, R: %d, C: %d, visitCount: %d\n", len(queue), coord.R, coord.C, visitedCounter[coord])
        queue = queue[1:]

        if coord.R == endLineIndex {
            count += visitedCounter[coord]
            continue
        }

        if grid[coord.R + 1][coord.C] == '^' {
            // splitter, add left and right

            if coord.C > 0 {
                checkAndEnqueue(coord.R + 1, coord.C - 1, visitedCounter[coord])
            }
            if coord.C < len(lines[0]) - 1 {
                checkAndEnqueue(coord.R + 1, coord.C + 1, visitedCounter[coord])
            }
        } else {
            // continue going down
            checkAndEnqueue(coord.R + 1, coord.C, visitedCounter[coord])
        }
    }

    return count
}
