package day06

import (
	"testing"

	"github.com/MagnusFrater/AdventOfCode-2015/internal/test"
)

func TestPart1(t *testing.T) {
	a := test.NewAssert(t)

	var testCases = []test.Case{
		{Input: "turn on 0,0 through 999,999", Expected: 1000000},
		{Input: "toggle 0,0 through 999,0", Expected: 1000},
		{Input: "turn on 499,499 through 500,500", Expected: 4},
	}

	for _, tc := range testCases {
		input := []string{tc.Input.(string)}
		a.Equals(tc, Part1(input))
	}
}

func TestPart2(t *testing.T) {
	a := test.NewAssert(t)

	var testCases = []test.Case{
		{Input: "turn on 0,0 through 0,0", Expected: 1},
		{Input: "toggle 0,0 through 999,999", Expected: 2000000},
	}

	for _, tc := range testCases {
		input := []string{tc.Input.(string)}
		a.Equals(tc, Part2(input))
	}
}
