package day12

import (
	"encoding/json"
	"regexp"
	"strconv"
)

var validNumber = regexp.MustCompile(`(-|)[0-9]+`)

// Part1 returns the output for Day 12 part 1.
func Part1(JSON string) int {
	var matches = validNumber.FindAllString(JSON, -1)

	var sum int
	for _, match := range matches {
		i, err := strconv.Atoi(match)
		if err != nil {
			panic(err)
		}
		sum += i
	}

	return sum
}

// Part2 returns the output for Day 12 part 2.
func Part2(JSON string) int {

	var f interface{}
	err := json.Unmarshal([]byte(JSON), &f)
	if err != nil {
		panic(err)
	}

	return sumItem(f)
}

func sumItem(i interface{}) int {
	switch t := i.(type) {
	case []interface{}:
		return sumSlice(t)
	case map[string]interface{}:
		return sumMap(t)
	case float64:
		return int(t)
	default:
		return 0
	}
}

func sumSlice(s []interface{}) int {
	var sum int

	for _, n := range s {
		sum += sumItem(n)
	}

	return sum
}

func sumMap(m map[string]interface{}) int {
	var sum int

	for _, v := range m {
		if v == "red" {
			return 0
		}

		sum += sumItem(v)
	}

	return sum
}
