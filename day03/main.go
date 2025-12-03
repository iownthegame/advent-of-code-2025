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
    // exmple line input 987654321111111
    joltage := 0
    for _, line := range lines {
        bank := parseBatteryBank(line)
        joltage += findJoltage(bank)
    }

    return joltage
}

func solvePart2(lines []string) int {
    // exmple line input 987654321111111
    joltage := 0
    for _, line := range lines {
        bank := parseBatteryBank(line)
        joltage += findJoltageWithTotalDigits(bank, 12)
    }

    return joltage
}

func parseBatteryBank(s string) []int {
    var bank []int

    for _, c := range s {
        battery, _ := strconv.Atoi(string(c))
        bank = append(bank, battery)
    }

    return bank
}

func findJoltage(bank []int) int {
    maxIndex := 0
    maxDigit := 0
    secondMaxDigit := 0
    // secondMaxIndex := 0

    for i := 0; i < len(bank)-1; i ++ {
        if (bank[i] > maxDigit) {
            maxDigit = bank[i]
            maxIndex = i
        }
    }

    for i := maxIndex + 1; i < len(bank); i ++ {
        if (bank[i] > secondMaxDigit) {
            secondMaxDigit = bank[i]
            // secondMaxIndex = i
        }
    }

    return maxDigit * 10 + secondMaxDigit
}

func findJoltageWithTotalDigits(bank []int, digits int) int {
    joltage := 0
    startIndex := 0
    for n := 0; n < digits; n++ {
        endIndex := len(bank) - digits + n
        maxDigit, maxIndex := findMaxDigit(bank, startIndex, endIndex)
        joltage = joltage * 10 + maxDigit
        startIndex = maxIndex + 1
    }
    return joltage
}

func findMaxDigit(s []int, startIndex int, endIndex int) (int, int) {
    maxIndex := 0
    maxDigit := 0
    for i := startIndex; i <= endIndex; i ++ {
        if (s[i] > maxDigit) {
            maxDigit = s[i]
            maxIndex = i
        }
    }
    return maxDigit, maxIndex
}
