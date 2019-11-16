package day03

import (
	"testing"
)

type testCase struct {
	Directions     string
	ExpectedHouses int
}

func TestPart1(t *testing.T) {
	var testCases = []testCase{
		{Directions: ">", ExpectedHouses: 2},
		{Directions: "^>v<", ExpectedHouses: 4},
		{Directions: "^v^v^v^v^v", ExpectedHouses: 2},
	}

	for _, tc := range testCases {
		var houses = Part1(tc.Directions)
		if houses != tc.ExpectedHouses {
			t.Errorf("Directions: %v\tExpected: %v\tActual: %v\n", tc.Directions, tc.ExpectedHouses, houses)
		}
	}
}

func TestPart2(t *testing.T) {
	var testCases = []testCase{
		{Directions: "^v", ExpectedHouses: 3},
		{Directions: "^>v<", ExpectedHouses: 3},
		{Directions: "^v^v^v^v^v", ExpectedHouses: 11},
	}

	for _, tc := range testCases {
		var houses = Part2(tc.Directions)
		if houses != tc.ExpectedHouses {
			t.Errorf("Directions: %v\tExpected: %v\tActual: %v\n", tc.Directions, tc.ExpectedHouses, houses)
		}
	}
}
