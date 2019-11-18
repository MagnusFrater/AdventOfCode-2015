package day08

import (
	"strings"
	"testing"

	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

type testCase struct {
	File          []string
	ExpectedCount int
}

func TestGetInMemoryCharacterCount(t *testing.T) {
	var testCases = []struct {
		Line          string
		ExpectedCount int
	}{
		{Line: `""`, ExpectedCount: 0},
		{Line: `"abc"`, ExpectedCount: 3},
		{Line: `"aaa\"aaa"`, ExpectedCount: 7},
		{Line: `"\x27"`, ExpectedCount: 1},
	}

	for _, tc := range testCases {
		var count = getInMemoryCharacterCount(tc.Line)
		if count != tc.ExpectedCount {
			t.Errorf("Line: %v\tExpected count: %v\tActual count: %v\n", tc.Line, tc.ExpectedCount, count)
		}
	}
}

func TestGetEncodedCharacterCount(t *testing.T) {
	var testCases = []struct {
		Line, EncodedLine string
		ExpectedCount     int
	}{
		{Line: `""`, EncodedLine: `"\"\""`, ExpectedCount: 6},
		{Line: `"abc"`, EncodedLine: `"\"abc\""`, ExpectedCount: 9},
		{Line: `"aaa\"aaa"`, EncodedLine: `"\"aaa\\\"aaa\""`, ExpectedCount: 16},
		{Line: `"\x27"`, EncodedLine: `"\"\\x27\""`, ExpectedCount: 11},
	}

	for _, tc := range testCases {
		var count = getEncodedCharacterCount(tc.Line)
		if count != tc.ExpectedCount {
			t.Errorf("Line: %v\tEncoded line: %v\tExpected count: %v\tActual count: %v\n",
				tc.Line, tc.EncodedLine, tc.ExpectedCount, count)
		}
	}
}

func TestPart1(t *testing.T) {
	var tc = testCase{
		File: []string{
			`""`,
			`"abc"`,
			`"aaa\"aaa"`,
			`"\x27"`,
		},
		ExpectedCount: 12,
	}

	var count = Part1(tc.File)
	if count != tc.ExpectedCount {
		t.Errorf("Expected count: %v\tActual count: %v\n", tc.ExpectedCount, count)
	}
}

func TestPart1_solution(t *testing.T) {
	var tc = testCase{
		File:          strings.Fields(load.InputFileAsString("input08.txt")),
		ExpectedCount: 1333,
	}

	var count = Part1(tc.File)
	if count != tc.ExpectedCount {
		t.Errorf("Expected count: %v\tActual count: %v\n", tc.ExpectedCount, count)
	}
}

func TestPart2(t *testing.T) {
	var tc = testCase{
		File: []string{
			`""`,
			`"abc"`,
			`"aaa\"aaa"`,
			`"\x27"`,
		},
		ExpectedCount: 19,
	}

	var count = Part2(tc.File)
	if count != tc.ExpectedCount {
		t.Errorf("Expected count: %v\tActual count: %v\n", tc.ExpectedCount, count)
	}
}

func TestPart2_solution(t *testing.T) {
	var tc = testCase{
		File:          strings.Fields(load.InputFileAsString("input08.txt")),
		ExpectedCount: 2046,
	}

	var count = Part2(tc.File)
	if count != tc.ExpectedCount {
		t.Errorf("Expected count: %v\tActual count: %v\n", tc.ExpectedCount, count)
	}
}
