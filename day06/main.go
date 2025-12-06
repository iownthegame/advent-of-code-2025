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

type Problem struct {
	Numbers []int
	Operator rune
}

func solvePart1(lines []string) int {
    // exmple line input "123 328  51 64 ", "*   +   *   +  "
    problems := []Problem{}

    for lineIdx, line := range lines {
        isSymbolLine := lineIdx == len(lines) - 1
        isFirstLine := lineIdx == 0

        parts := strings.Split(line, " ")
        problemIdx := 0

        for _, part := range parts {
            if part == "" {
                continue
            }

            if isSymbolLine {
                // set symbol of the problem
                problems[problemIdx].Operator = []rune(part)[0]

            } else {
                number, _ := strconv.Atoi(part)
                if isFirstLine {
                    // initial problem with one number in numbers and operator null
                    problems = append(problems, Problem{
                        Numbers: []int{number},
                        Operator: 0, // null
                    })
                } else {
                    // append a number to the numbers
                    problems[problemIdx].Numbers = append(problems[problemIdx].Numbers, number)
                }
            }
            problemIdx += 1
        }
    }

    total := 0
    for _, problem := range problems {
        total += calculateProblem(problem.Numbers, problem.Operator)
    }
    return total
}

func calculateProblem(numbers []int, operator rune) int {
    total := numbers[0]

    for i := 1; i < len(numbers); i++ {
        number := numbers[i]

        switch operator {
        case '+':
            total += number
        case '*':
            total *= number
        }
    }

    return total
}

func solvePart2(lines []string) int {
    /*
		"123 328  51 64 ",
 		" 45 64  387 23 ",
  		"  6 98  215 314",
		"*   +   *   +  ",
    */

    total := 0
    problemOperator := rune(0)
    problemNumbers := []int{}
    symbolLineIdx := len(lines) - 1

    // parse from the right most line
    j := len(lines[0]) - 1
    for {
        if j < 0 {
            break
        }
        currentNumberString := ""
        for i := 0; i < len(lines); i ++ {
            if i == symbolLineIdx {
                currentNumber, _ := strconv.Atoi(currentNumberString)
                problemNumbers = append(problemNumbers, currentNumber)

                problemOperator = rune(lines[i][j])
                if problemOperator == ' ' {
                    // continue with next time
                    j--
                } else {
                    // end of problem
                    // fmt.Printf("problemNumbers %d\n", problemNumbers)
                    total += calculateProblem(problemNumbers, problemOperator)
                    // reset problem
                    problemNumbers = nil
                    problemOperator = 0
                    j -= 2 // skip the next empty line
                }
            } else {
                if (lines[i][j] != ' ') {
                    currentNumberString += string(lines[i][j])
                }
            }
        }
    }

    return total
}
