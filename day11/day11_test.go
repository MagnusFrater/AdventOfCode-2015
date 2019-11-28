package day11

import (
	"testing"

	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

type testCase struct {
	PreviousPassword string
	ExpectedPassword string
}

func TestIncrementPassword(t *testing.T) {
	var testCases = []struct {
		PreviousPassword            string
		ExpectedIncrementedPassword string
	}{
		{PreviousPassword: "xx", ExpectedIncrementedPassword: "xy"},
		{PreviousPassword: "xy", ExpectedIncrementedPassword: "xz"},
		{PreviousPassword: "xz", ExpectedIncrementedPassword: "ya"},
		{PreviousPassword: "ya", ExpectedIncrementedPassword: "yb"},
		{PreviousPassword: "yb", ExpectedIncrementedPassword: "yc"},
		{PreviousPassword: "azzzz", ExpectedIncrementedPassword: "baaaa"},
	}

	for _, tc := range testCases {
		var incrementedPassword = string(incrementPassword([]byte(tc.PreviousPassword)))
		if incrementedPassword != tc.ExpectedIncrementedPassword {
			t.Errorf("Input: %v\tExpected: %v\tActual: %v\n",
				tc.PreviousPassword, tc.ExpectedIncrementedPassword, incrementedPassword)
		}
	}
}

func TestIsValid(t *testing.T) {
	var testCases = []struct {
		Password        string
		ExpectedIsValid bool
	}{
		{Password: "hijklmmn", ExpectedIsValid: false},
		{Password: "abbceffg", ExpectedIsValid: false},
		{Password: "abbcegjk", ExpectedIsValid: false},
		{Password: "abcdffaa", ExpectedIsValid: true},
		{Password: "ghjaabcc", ExpectedIsValid: true},
	}

	for _, tc := range testCases {
		var isValid = isValid([]byte(tc.Password))
		if isValid != tc.ExpectedIsValid {
			t.Errorf("Input: %v\tExpected: %v\tActual: %v\n", tc.Password, tc.ExpectedIsValid, isValid)
		}
	}
}

func TestPart1(t *testing.T) {
	var testCases = []testCase{
		{
			PreviousPassword: "abcdefgh",
			ExpectedPassword: "abcdffaa",
		},
		{
			PreviousPassword: "ghijklmn",
			ExpectedPassword: "ghjaabcc",
		},
	}

	for _, tc := range testCases {
		var password = Part1(tc.PreviousPassword)
		if password != tc.ExpectedPassword {
			t.Errorf("Input: %v\tExpected: %v\tActual: %v\n", tc.PreviousPassword, tc.ExpectedPassword, password)
		}
	}
}

func TestPart1_solution(t *testing.T) {
	var tc = testCase{
		PreviousPassword: load.InputFileAsString("input11.txt"),
		ExpectedPassword: "hepxxyzz",
	}

	var password = Part1(tc.PreviousPassword)
	if password != tc.ExpectedPassword {
		t.Errorf("Input: %v\tExpected: %v\tActual: %v\n", tc.PreviousPassword, tc.ExpectedPassword, password)
	}
}

func TestPart2_solution(t *testing.T) {
	var tc = testCase{
		PreviousPassword: load.InputFileAsString("input11.txt"),
		ExpectedPassword: "heqaabcc",
	}

	var password = Part2(tc.PreviousPassword)
	if password != tc.ExpectedPassword {
		t.Errorf("Input: %v\tExpected: %v\tActual: %v\n", tc.PreviousPassword, tc.ExpectedPassword, password)
	}
}
