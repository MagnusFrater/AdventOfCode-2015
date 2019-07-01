package day04

import (
	"testing"

	"github.com/MagnusFrater/AdventOfCode-2015/internal/test"
)

func TestPart1(t *testing.T) {
	a := test.NewAssert(t)

	var testCases = []test.Case{
		{Input: "abcdef", Expected: 609043},
		{Input: "pqrstuv", Expected: 1048970},
	}

	for _, tc := range testCases {
		a.Equals(tc, Part1(tc.Input.(string), 1))
	}
}
