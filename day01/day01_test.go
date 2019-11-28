package day01

import (
	"testing"

	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
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

func TestPart1_solution(t *testing.T) {
	var testCase = struct {
		Instructions  string
		ExpectedFloor int
	}{
		Instructions:  load.InputFileAsString("input01.txt"),
		ExpectedFloor: 74,
	}

	var floor = Part1(testCase.Instructions)
	if floor != testCase.ExpectedFloor {
		t.Errorf("Expected: %v\tActual:%v\n", testCase.ExpectedFloor, floor)
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

func TestPart2_solution(t *testing.T) {
	var testCase = struct {
		Instructions  string
		ExpectedIndex int
	}{
		Instructions:  load.InputFileAsString("input01.txt"),
		ExpectedIndex: 1795,
	}

	var index = Part2(testCase.Instructions)
	if index != testCase.ExpectedIndex {
		t.Errorf("Expected: %v\tActual:%v\n", testCase.ExpectedIndex, index)
	}
}
