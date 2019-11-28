package day06

import (
	"strconv"
	"strings"
)

// Part1 returns the output for Day 06 part 1.
func Part1(s []string) int {
	lights := createLights()

	for _, line := range s {
		stateChange := getStateChange(line)
		minX, minY, maxX, maxY := getRange(line)

		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				switch stateChange {
				case 1:
					lights[y][x] = 1
				case -1:
					lights[y][x] = 0
				default:
					if lights[y][x] == 0 {
						lights[y][x] = 1
					} else {
						lights[y][x] = 0
					}
				}
			}
		}
	}

	return getTotalLightsOn(lights)
}

// Part2 returns the output for Day 06 part 2.
func Part2(s []string) int {
	lights := createLights()

	for _, line := range s {
		stateChange := getStateChange(line)
		minX, minY, maxX, maxY := getRange(line)

		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				switch stateChange {
				case 1:
					lights[y][x]++
				case -1:
					if lights[y][x] > 0 {
						lights[y][x]--
					}
				default:
					lights[y][x] += 2
				}
			}
		}
	}

	return getTotalBrightness(lights)
}

func createLights() [][]int {
	lights := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		lights[i] = make([]int, 1000)
	}
	return lights
}

func getStateChange(line string) int {
	// 'toggle' === 0
	var stateChange int

	if strings.HasPrefix(line, "turn on") {
		stateChange = 1
	} else if strings.HasPrefix(line, "turn off") {
		stateChange = -1
	}

	return stateChange
}

func getRange(line string) (minX, minY, maxX, maxY int) {
	split := strings.Split(line, " ")
	minRange := strings.Split(split[len(split)-3], ",")
	maxRange := strings.Split(split[len(split)-1], ",")

	minX, _ = strconv.Atoi(minRange[0])
	minY, _ = strconv.Atoi(minRange[1])
	maxX, _ = strconv.Atoi(maxRange[0])
	maxY, _ = strconv.Atoi(maxRange[1])
	return
}

func getTotalLightsOn(lights [][]int) int {
	var count int

	for y := 0; y < len(lights); y++ {
		for x := 0; x < len(lights[0]); x++ {
			if lights[y][x] == 1 {
				count++
			}
		}
	}

	return count
}

func getTotalBrightness(lights [][]int) int {
	var brightness int

	for y := 0; y < len(lights); y++ {
		for x := 0; x < len(lights[0]); x++ {
			brightness += lights[y][x]
		}
	}

	return brightness
}
