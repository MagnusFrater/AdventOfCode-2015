package day02

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	type testCase struct {
		Dimensions    string
		ExpectedPaper int
	}

	var testCases = []testCase{
		{Dimensions: "2x3x4", ExpectedPaper: 58},
		{Dimensions: "1x1x10", ExpectedPaper: 43},
	}

	for _, tc := range testCases {
		var input = strings.Split(tc.Dimensions, "\n")
		var paper = Part1(input)
		if paper != tc.ExpectedPaper {
			t.Errorf("Dimensions: %v\tExpected: %v\tActual: %v\n", tc.Dimensions, tc.ExpectedPaper, paper)
		}
	}
}

func TestPart2(t *testing.T) {
	type testCase struct {
		Dimensions     string
		ExpectedRibbon int
	}

	var testCases = []testCase{
		{Dimensions: "2x3x4", ExpectedRibbon: 34},
		{Dimensions: "1x1x10", ExpectedRibbon: 14},
	}

	for _, tc := range testCases {
		var input = strings.Split(tc.Dimensions, "\n")
		var ribbon = Part2(input)
		if ribbon != tc.ExpectedRibbon {
			t.Errorf("Dimensions: %v\tExpected: %v\tActual: %v\n", tc.Dimensions, tc.ExpectedRibbon, ribbon)
		}
	}
}
