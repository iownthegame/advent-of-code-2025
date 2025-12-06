package main

import (
	"fmt"
    "slices"
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

type IngredientRange struct {
	Start int
	End int
}

func solvePart1(lines []string) int {
    // exmple line input few lines of "3-5", empty line, few lines of "4"
    ingredientRanges := []IngredientRange{}
    rangesParsed := false
    count := 0

    for _, line := range lines {
        if line == "" {
            rangesParsed = true
            continue
        }

        if rangesParsed {
            // parse ingredient ids
            ingredientId, _ := strconv.Atoi(line)
            if isIngredientFresh(ingredientId, ingredientRanges) {
                count += 1
            }
        } else {
            // parse ingredient ranges
            parts := strings.Split(line, "-")
            rangeStart, _ := strconv.Atoi(parts[0])
            rangeEnd, _ := strconv.Atoi(parts[1])
            ingredientRanges = append(ingredientRanges, IngredientRange{rangeStart, rangeEnd})
        }
    }

    return count
}

type StringSet map[int]struct{} // Set definition

func solvePart2TLE(lines []string) int {
    // exmple line input few lines of "3-5", empty line, few lines of "4"
    freshIngredientIds := StringSet{}

    for _, line := range lines {
        // no need to parse the ingredient ids anymore
        if line == "" {
            break
        }

        // parse ingredient ranges
        parts := strings.Split(line, "-")
        rangeStart, _ := strconv.Atoi(parts[0])
        rangeEnd, _ := strconv.Atoi(parts[1])

        for i := rangeStart; i <= rangeEnd; i++ {
            freshIngredientIds[i] = struct{}{}
        }
    }

    return len(freshIngredientIds)
}

func solvePart2(lines []string) int {
    // exmple line input few lines of "3-5", empty line, few lines of "4"
    ingredientRanges := []IngredientRange{}

    for _, line := range lines {
        // no need to parse the ingredient ids anymore
        if line == "" {
            break
        }

        // parse ingredient ranges
        parts := strings.Split(line, "-")
        rangeStart, _ := strconv.Atoi(parts[0])
        rangeEnd, _ := strconv.Atoi(parts[1])
        ingredientRanges = append(ingredientRanges, IngredientRange{rangeStart, rangeEnd})
    }

    // sort ingredient ranges
    slices.SortFunc(ingredientRanges, func(a IngredientRange, b IngredientRange) int {
        if a.Start != b.Start {
            return a.Start - b.Start
        }

        return a.End - b.End
    })

    // merge ingredient ranges
    mergedRanges := mergeIngredientRanges(ingredientRanges)

    count := 0
    for _, ingredientRange := range mergedRanges {
        count += ingredientRange.End - ingredientRange.Start + 1
    }

    return count
}

func mergeIngredientRanges(ingredientRanges []IngredientRange) []IngredientRange {
    mergedRanges := []IngredientRange{ingredientRanges[0]}

    for i := 1; i < len(ingredientRanges); i++ {
        currentRange := ingredientRanges[i]

        lastMergedIdx := len(mergedRanges) - 1
        lastMergedRange := mergedRanges[lastMergedIdx]

        if currentRange.Start <= lastMergedRange.End {
            if currentRange.End > lastMergedRange.End {
                mergedRanges[lastMergedIdx] = IngredientRange{lastMergedRange.Start, currentRange.End}
            }
        } else {
            mergedRanges = append(mergedRanges, currentRange)
        }
    }

    return mergedRanges
}

func isIngredientFresh(ingredientId int, ingredientRanges []IngredientRange) bool {
    for _, ingredientRange := range ingredientRanges {
        if ingredientId >= ingredientRange.Start && ingredientId <= ingredientRange.End {
            return true
        }
    }
    return false
}
