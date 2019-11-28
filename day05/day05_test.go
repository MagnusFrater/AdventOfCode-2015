package day05

import (
	"strings"
	"testing"

	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

type testCase struct {
	String            string
	ExpectedNiceCount int
}

func TestPart1(t *testing.T) {
	var testCases = []testCase{
		{String: "ugknbfddgicrmopn", ExpectedNiceCount: 1},
		{String: "aaa", ExpectedNiceCount: 1},
		{String: "jchzalrnumimnmhp", ExpectedNiceCount: 0},
		{String: "haegwjzuvuyypxyu", ExpectedNiceCount: 0},
		{String: "dvszwmarrgswjxmb", ExpectedNiceCount: 0},
	}

	for _, tc := range testCases {
		var niceCount = Part1([]string{tc.String})
		if niceCount != tc.ExpectedNiceCount {
			t.Errorf("String: %v\tExpected: %v\tActual: %v\n", tc.String, tc.ExpectedNiceCount, niceCount)
		}
	}
}

func TestPart1_solution(t *testing.T) {
	var testCase = struct {
		Strings           []string
		ExpectedNiceCount int
	}{
		Strings:           strings.Fields(load.InputFileAsString("input05.txt")),
		ExpectedNiceCount: 255,
	}

	var niceCount = Part1(testCase.Strings)
	if niceCount != testCase.ExpectedNiceCount {
		t.Errorf("Expected: %v\tActual: %v\n", testCase.ExpectedNiceCount, niceCount)
	}
}

func TestPart2(t *testing.T) {
	var testCases = []testCase{
		{String: "qjhvhtzxzqqjkmpb", ExpectedNiceCount: 1},
		{String: "xxyxx", ExpectedNiceCount: 1},
		{String: "uurcxstgmygtbstg", ExpectedNiceCount: 0},
		{String: "ieodomkazucvgmuy", ExpectedNiceCount: 0},
	}

	for _, tc := range testCases {
		var niceCount = Part2([]string{tc.String})
		if niceCount != tc.ExpectedNiceCount {
			t.Errorf("String: %v\tExpected: %v\tActual: %v\n", tc.String, tc.ExpectedNiceCount, niceCount)
		}
	}
}

func TestPart2_solution(t *testing.T) {
	var testCase = struct {
		Strings           []string
		ExpectedNiceCount int
	}{
		Strings:           strings.Fields(load.InputFileAsString("input05.txt")),
		ExpectedNiceCount: 55,
	}

	var niceCount = Part2(testCase.Strings)
	if niceCount != testCase.ExpectedNiceCount {
		t.Errorf("Expected: %v\tActual: %v\n", testCase.ExpectedNiceCount, niceCount)
	}
}
