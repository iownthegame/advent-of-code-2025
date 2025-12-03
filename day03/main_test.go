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
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}

	tests := []testCase{
		{
			inputLines: testInput,
			expectedP1: 357,
			expectedP2: 3121910778619,
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
