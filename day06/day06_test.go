package day06

import (
	"strings"
	"testing"

	"github.com/MagnusFrater/AdventOfCode-2015/internal/load"
)

type testCase struct {
	Instructions   string
	ExpectedLights int
}

func TestPart1(t *testing.T) {
	var testCases = []testCase{
		{Instructions: "turn on 0,0 through 999,999", ExpectedLights: 1000000},
		{Instructions: "toggle 0,0 through 999,0", ExpectedLights: 1000},
		{Instructions: "turn on 499,499 through 500,500", ExpectedLights: 4},
	}

	for _, tc := range testCases {
		var lights = Part1([]string{tc.Instructions})
		if lights != tc.ExpectedLights {
			t.Errorf("Instructions: %v\tExpected lights: %v\tActual lights: %v\n", tc.Instructions, tc.ExpectedLights, lights)
		}
	}
}

func TestPart1_solution(t *testing.T) {
	var testCase = struct {
		Instructions   []string
		ExpectedLights int
	}{
		Instructions:   strings.Split(load.InputFileAsString("input06.txt"), "\n"),
		ExpectedLights: 377891,
	}

	var lights = Part1(testCase.Instructions)
	if lights != testCase.ExpectedLights {
		t.Errorf("Expected lights: %v\tActual lights: %v\n", testCase.ExpectedLights, lights)
	}
}

func TestPart2(t *testing.T) {
	var testCases = []testCase{
		{Instructions: "turn on 0,0 through 0,0", ExpectedLights: 1},
		{Instructions: "toggle 0,0 through 999,999", ExpectedLights: 2000000},
	}

	for _, tc := range testCases {
		var lights = Part2([]string{tc.Instructions})
		if lights != tc.ExpectedLights {
			t.Errorf("Instructions: %v\tExpected lights: %v\tActual lights: %v\n", tc.Instructions, tc.ExpectedLights, lights)
		}
	}
}

func TestPart2_solution(t *testing.T) {
	var testCase = struct {
		Instructions   []string
		ExpectedLights int
	}{
		Instructions:   strings.Split(load.InputFileAsString("input06.txt"), "\n"),
		ExpectedLights: 14110788,
	}

	var lights = Part2(testCase.Instructions)
	if lights != testCase.ExpectedLights {
		t.Errorf("Expected lights: %v\tActual lights: %v\n", testCase.ExpectedLights, lights)
	}
}
