package day08

import (
	"regexp"
)

var r = regexp.MustCompile(`(\\\\)|(\\\")|(\\x[0-9a-f]{2})`)

// Part1 returns the output for Day 08 part 1.
func Part1(file []string) int {
	var total int

	for _, line := range file {
		var codeCharacterCount = len(line)
		var inMemoryCharacterCount = getInMemoryCharacterCount(line)
		total += codeCharacterCount - inMemoryCharacterCount
	}

	return total
}

// Part2 returns the output for Day 08 part 2.
func Part2(file []string) int {
	var total int

	for _, line := range file {
		var encodedCharacterCount = getEncodedCharacterCount(line)
		var codeCharacterCount = len(line)
		total += encodedCharacterCount - codeCharacterCount
	}

	return total
}

// "aaa\"aaa" -> aaa\"aaa
func getLineData(line string) string {
	return line[1 : len(line)-1]
}

// "\"" -> [\"], "\\" -> [\\], "\x27" -> [\x27]
func getEscapeSequences(line string) []string {
	return r.FindAllString(line, -1)
}

func getInMemoryCharacterCount(line string) int {
	var count int

	for _, seq := range getEscapeSequences(getLineData(line)) {
		count += len(seq) - 1
	}

	return len(line) - 2 - count
}

func getEncodedCharacterCount(line string) int {
	var count int

	for i := range line {
		if line[i] == '\\' || line[i] == '"' {
			count++
		}
	}

	return len(line) + 2 + count
}
