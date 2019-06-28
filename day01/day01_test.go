package day01

import "testing"

func TestPart1(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, c := range cases {
		actual := Part1(c.input)

		if actual != c.expected {
			t.Errorf("Input: %q\tExpected: %d\tActual: %d", c.input, c.expected, actual)
		}
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{")", 1},
		{"()())", 5},
	}

	for _, c := range cases {
		actual := Part2(c.input)

		if actual != c.expected {
			t.Errorf("Input: %q\tExpected: %d\tActual: %d", c.input, c.expected, actual)
		}
	}
}
