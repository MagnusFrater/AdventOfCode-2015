package day03

import (
	"testing"

	"github.com/MagnusFrater/AdventOfCode-2015/internal/test"
)

func TestPart1(t *testing.T) {
	a := test.NewAssert(t)

	var testCases = []test.Case{
		{Input: ">", Expected: 2},
		{Input: "^>v<", Expected: 4},
		{Input: "^v^v^v^v^v", Expected: 2},
	}

	for _, tc := range testCases {
		a.Equals(tc, Part1(tc.Input.(string)))
	}
}

func TestPart2(t *testing.T) {
	a := test.NewAssert(t)

	var testCases = []test.Case{
		{Input: "^v", Expected: 3},
		{Input: "^>v<", Expected: 3},
		{Input: "^v^v^v^v^v", Expected: 11},
	}

	for _, tc := range testCases {
		a.Equals(tc, Part2(tc.Input.(string)))
	}
}
