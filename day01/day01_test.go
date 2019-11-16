package day01

import (
	"testing"
)

type testCase struct {
	Instructions  string
	ExpectedFloor int
}

func TestPart1(t *testing.T) {
	var testCases = []testCase{
		{Instructions: "(())", ExpectedFloor: 0},
		{Instructions: "()()", ExpectedFloor: 0},
		{Instructions: "(((", ExpectedFloor: 3},
		{Instructions: "(()(()(", ExpectedFloor: 3},
		{Instructions: "))(((((", ExpectedFloor: 3},
		{Instructions: "())", ExpectedFloor: -1},
		{Instructions: ")))", ExpectedFloor: -3},
		{Instructions: ")())())", ExpectedFloor: -3},
	}

	for _, tc := range testCases {
		var floor = Part1(tc.Instructions)
		if floor != tc.ExpectedFloor {
			t.Errorf("Input: %v:\tExpected: %v\tActual:%v\n", tc.Instructions, tc.ExpectedFloor, floor)
		}
	}
}

func TestPart2(t *testing.T) {
	testCases := []testCase{
		{Instructions: ")", ExpectedFloor: 1},
		{Instructions: "()())", ExpectedFloor: 5},
	}

	for _, tc := range testCases {
		var floor = Part2(tc.Instructions)
		if floor != tc.ExpectedFloor {
			t.Errorf("Input: %v:\tExpected: %v\tActual:%v\n", tc.Instructions, tc.ExpectedFloor, floor)
		}
	}
}
