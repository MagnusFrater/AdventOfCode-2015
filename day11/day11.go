package day11

// Part1 returns the output for Day 11 part 1.
func Part1(previous string) string {
	var password = []byte(previous)
	password = incrementPassword(password)
	for !(isValid(password)) {
		password = incrementPassword(password)
	}

	return string(password)
}

// Part2 returns the output for Day 11 part 2.
func Part2(previous string) string {
	previous = Part1(previous)
	return Part1(previous)
}

func incrementPassword(password []byte) []byte {
	var index = len(password) - 1

	if password[index] == 'z' {
		for password[index] == 'z' {
			password[index] = 'a'
			index--
		}
		password[index]++
		return password
	}

	password[index]++
	return password
}

func isValid(password []byte) bool {
	var length = len(password)
	var foundStraight bool
	var pairsIndices = []int{}

	for i, c := range password {
		// check invalid letters
		switch c {
		case 'i', 'o', 'l':
			return false
		}

		// check increasing straight
		if i < length-2 {
			if password[i+1] == c+1 && password[i+2] == c+2 {
				foundStraight = true
			}
		}

		// check for distinct pairs
		if i < length-1 && password[i+1] == c {
			var isOverlapping bool
			for _, n := range pairsIndices {
				if n == i || n == i+1 {
					isOverlapping = true
					break
				}
			}

			if !isOverlapping {
				pairsIndices = append(pairsIndices, i, i+1)
			}
		}
	}

	return foundStraight && len(pairsIndices) >= 4
}
