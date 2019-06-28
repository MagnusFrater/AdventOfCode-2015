package day01

// Part1 returns the result of Day 01 part 1.
func Part1(input string) int {
	var floor int

	for _, c := range input {
		floor += travel(c)
	}

	return floor
}

// Part2 returns the result of Day 01 part 2.
func Part2(input string) int {
	var floor int

	for i, c := range input {
		floor += travel(c)

		if floor < 0 {
			return i + 1
		}
	}

	return -1
}

func travel(c rune) int {
	if c == '(' {
		return 1
	} else if c == ')' {
		return -1
	}

	return 0
}
