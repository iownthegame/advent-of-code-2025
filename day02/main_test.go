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
		"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",
	}

	tests := []testCase{
		{
			inputLines: testInput,
			expectedP1: 1227775554,
			expectedP2: 4174379265,
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
