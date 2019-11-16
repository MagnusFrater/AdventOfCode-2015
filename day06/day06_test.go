package day06

import (
	"testing"
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
