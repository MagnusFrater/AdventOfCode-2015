package day04

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

const span int = 1000000

func hash(s, goal string, id, maxRoutines int, ch chan int) {
	min := span * id
	max := min + span
	len := len(goal)

	for i := min; i < max; i++ {
		hash := md5.Sum([]byte(s + strconv.Itoa(i)))
		if fmt.Sprintf("%x", hash)[:len] == goal {
			ch <- i
			return
		}
	}

	go hash(s, goal, id+maxRoutines, maxRoutines, ch)
}

// Part1 returns the output for Day 04 part 1.
func Part1(s string, maxRoutines int) int {
	const goal = "00000"
	ch := make(chan int)

	for i := 0; i < maxRoutines; i++ {
		go hash(s, goal, i, maxRoutines, ch)
	}

	return <-ch
}

// Part2 returns the output for Day 04 part 2.
func Part2(s string, maxRoutines int) int {
	const goal = "000000"
	ch := make(chan int)

	for i := 0; i < maxRoutines; i++ {
		go hash(s, goal, i, maxRoutines, ch)
	}

	return <-ch
}
