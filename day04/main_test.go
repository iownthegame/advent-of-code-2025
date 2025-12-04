package main

import (
	"testing"
	"fmt"
)

type testCase struct {
	inputLines []string

	expectedP1 int

	expectedP2 int
}

func TestDay01(t *testing.T) {
	testInput := []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}

	tests := []testCase{
		{
			inputLines: testInput,
			expectedP1: 13,
			expectedP2: 43,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {

			resultP1 := solvePart1(tt.inputLines)
			if resultP1 != tt.expectedP1 {
				t.Errorf("Part 1 FAILED. Expected %d, Got %d", tt.expectedP1, resultP1)
			}

			resultP2 := solvePart2(tt.inputLines)
			if resultP2 != tt.expectedP2 {
				t.Errorf("Part 2 FAILED. Expected %d, Got %d", tt.expectedP2, resultP2)
			}
		})
	}
}
