package day02

import (
	"strings"
	"testing"

	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

func TestPart1(t *testing.T) {
	var testCases = []struct {
		Dimensions    string
		ExpectedPaper int
	}{
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

func TestPart1_solution(t *testing.T) {
	var testCase = struct {
		Dimensions    []string
		ExpectedPaper int
	}{
		Dimensions:    strings.Fields(load.InputFileAsString("input02.txt")),
		ExpectedPaper: 1586300,
	}

	var paper = Part1(testCase.Dimensions)
	if paper != testCase.ExpectedPaper {
		t.Errorf("Expected: %v\tActual: %v\n", testCase.ExpectedPaper, paper)
	}
}

func TestPart2(t *testing.T) {
	var testCases = []struct {
		Dimensions     string
		ExpectedRibbon int
	}{
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

func TestPart2_solution(t *testing.T) {
	var testCase = struct {
		Dimensions     []string
		ExpectedRibbon int
	}{
		Dimensions:     strings.Fields(load.InputFileAsString("input02.txt")),
		ExpectedRibbon: 3737498,
	}

	var ribbon = Part2(testCase.Dimensions)
	if ribbon != testCase.ExpectedRibbon {
		t.Errorf("Expected: %v\tActual: %v\n", testCase.ExpectedRibbon, ribbon)
	}
}
