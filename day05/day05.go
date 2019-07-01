package day05

import (
	"strings"
)

// Part1 returns the output for Day 05 part 1.
func Part1(s []string) int {
	var count int

	for _, str := range s {
		if !hasForbiddenStrings(str) && countVowels(str) >= 3 && hasDoubleLetters(str) {
			count++
		}
	}

	return count
}

// Part2 returns the output for Day 05 part 2.
func Part2(s []string) int {
	var count int

	for _, str := range s {
		if hasDoubleDouble(str) && hasLetterSandwich(str) {
			count++
		}
	}

	return count
}

func hasForbiddenStrings(s string) bool {
	forbidden := []string{"ab", "cd", "pq", "xy"}

	for _, f := range forbidden {
		if strings.Count(s, f) > 0 {
			return true
		}
	}

	return false
}

func countVowels(s string) int {
	var count int

	for _, r := range s {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}

	return count
}

func hasDoubleLetters(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func hasDoubleDouble(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if strings.Contains(s[i+2:], s[i:i+2]) {
			return true
		}
	}
	return false
}

func hasLetterSandwich(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}
