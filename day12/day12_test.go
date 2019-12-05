package day12

import (
	"testing"

	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

type testCase struct {
	JSON        string
	ExpectedSum int
}

func TestPart1(t *testing.T) {
	var testCases = []testCase{
		{JSON: `[1,2,3]`, ExpectedSum: 6},
		{JSON: `{"a":2,"b":4}`, ExpectedSum: 6},
		{JSON: `[[[3]]]`, ExpectedSum: 3},
		{JSON: `{"a":{"b":4},"c":-1}`, ExpectedSum: 3},
		{JSON: `{"a":[-1,1]}`, ExpectedSum: 0},
		{JSON: `[-1,{"a":1}]`, ExpectedSum: 0},
		{JSON: `[]`, ExpectedSum: 0},
		{JSON: `{}`, ExpectedSum: 0},
	}

	for _, tc := range testCases {
		var sum = Part1(tc.JSON)
		if sum != tc.ExpectedSum {
			t.Errorf("JSON: %v\tExpected: %v\tActual: %v\n", tc.JSON, tc.ExpectedSum, sum)
		}
	}
}

func TestPart1_solution(t *testing.T) {
	var tc = testCase{
		JSON:        load.InputFileAsString("input12.txt"),
		ExpectedSum: 191164,
	}

	var sum = Part1(tc.JSON)
	if sum != tc.ExpectedSum {
		t.Errorf("Expected: %v\tActual: %v\n", tc.ExpectedSum, sum)
	}
}

func TestPart2(t *testing.T) {
	var testCases = []testCase{
		{JSON: `[1,2,3]`, ExpectedSum: 6},
		{JSON: `[1,{"c":"red","b":2},3]`, ExpectedSum: 4},
		{JSON: `{"d":"red","e":[1,2,3,4],"f":5}`, ExpectedSum: 0},
		{JSON: `[1,"red",5]`, ExpectedSum: 6},
	}

	for _, tc := range testCases {
		var sum = Part2(tc.JSON)
		if sum != tc.ExpectedSum {
			t.Errorf("JSON: %v\tExpected: %v\tActual: %v\n", tc.JSON, tc.ExpectedSum, sum)
		}
	}
}

func TestPart2_solution(t *testing.T) {
	var tc = testCase{
		JSON:        load.InputFileAsString("input12.txt"),
		ExpectedSum: 87842,
	}

	var sum = Part2(tc.JSON)
	if sum != tc.ExpectedSum {
		t.Errorf("Expected: %v\tActual: %v\n", tc.ExpectedSum, sum)
	}
}
