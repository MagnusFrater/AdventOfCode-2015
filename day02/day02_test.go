package day02

import (
	"strings"
	"testing"

	"github.com/MagnusFrater/AdventOfCode-2015/internal/test"
)

func TestPart1(t *testing.T) {
	a := test.NewAssert(t)

	var testCases = []test.Case{
		{Input: "2x3x4", Expected: 58},
		{Input: "1x1x10", Expected: 43},
	}

	for _, tc := range testCases {
		var input = strings.Split(tc.Input.(string), "\n")
		a.Equals(tc, Part1(input))
	}
}

func TestPart2(t *testing.T) {
	a := test.NewAssert(t)

	var testCases = []test.Case{
		{Input: "2x3x4", Expected: 34},
		{Input: "1x1x10", Expected: 14},
	}

	for _, tc := range testCases {
		var input = strings.Split(tc.Input.(string), "\n")
		a.Equals(tc, Part2(input))
	}
}
