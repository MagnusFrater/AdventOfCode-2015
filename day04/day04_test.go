package day04

import (
	"testing"
)

func TestPart1(t *testing.T) {
	type testCase struct {
		SecretKey string
		Expected  int
	}

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
