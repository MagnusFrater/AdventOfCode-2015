package day04

import (
	"testing"

	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

type testCase struct {
	SecretKey string
	Expected  int
}

func TestPart1(t *testing.T) {
	var testCases = []testCase{
		{SecretKey: "abcdef", Expected: 609043},
		{SecretKey: "pqrstuv", Expected: 1048970},
	}

	for _, tc := range testCases {
		var actual = Part1(tc.SecretKey, 1)
		if actual != tc.Expected {
			t.Errorf("Secret Key: %v\tExpected: %v\tActual: %v\n", tc.SecretKey, tc.Expected, actual)
		}
	}
}

func TestPart1_solution(t *testing.T) {
	var tc = testCase{
		SecretKey: load.InputFileAsString("input04.txt"),
		Expected:  117946,
	}

	var actual = Part1(tc.SecretKey, 1)
	if actual != tc.Expected {
		t.Errorf("Secret Key: %v\tExpected: %v\tActual: %v\n", tc.SecretKey, tc.Expected, actual)
	}
}

func TestPart2_solution(t *testing.T) {
	var tc = testCase{
		SecretKey: load.InputFileAsString("input04.txt"),
		Expected:  3938038,
	}

	var actual = Part2(tc.SecretKey, 1)
	if actual != tc.Expected {
		t.Errorf("Secret Key: %v\tExpected: %v\tActual: %v\n", tc.SecretKey, tc.Expected, actual)
	}
}
