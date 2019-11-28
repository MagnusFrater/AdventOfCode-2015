package day10

import (
	"strconv"
	"testing"

	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

type testCase struct {
	Input          string
	ExpectedLength int
}

func TestLookAndSay(t *testing.T) {
	var testCases = []struct {
		Input  int
		Output string
	}{
		{Input: 1, Output: "11"},
		{Input: 11, Output: "21"},
		{Input: 21, Output: "1211"},
		{Input: 1211, Output: "111221"},
		{Input: 111221, Output: "312211"},
	}

	for _, tc := range testCases {
		var output = lookAndSay(strconv.Itoa(tc.Input))
		if output != tc.Output {
			t.Errorf("Input: %v\tExpected: %v\tActual: %v\n", tc.Input, tc.Output, output)
		}
	}
}

func TestPart1_solution(t *testing.T) {
	var tc = testCase{
		Input:          load.InputFileAsString("input10.txt"),
		ExpectedLength: 252594,
	}

	var length = Part1(tc.Input)
	if length != tc.ExpectedLength {
		t.Errorf("Input: %v\tExpected: %v\tActual: %v\n", tc.Input, tc.ExpectedLength, length)
	}
}

// func TestPart2_solution(t *testing.T) {
// 	var tc = testCase{
// 		Input:          load.InputFileAsString("input10.txt"),
// 		ExpectedLength: 3579328,
// 	}

// 	var length = Part2(tc.Input)
// 	if length != tc.ExpectedLength {
// 		t.Errorf("Input: %v\tExpected: %v\tActual: %v\n", tc.Input, tc.ExpectedLength, length)
// 	}
// }
