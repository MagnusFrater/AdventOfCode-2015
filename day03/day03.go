package day03

import (
	"fmt"
	"log"
	"strconv"
)

// Point is a coordinate on the cartesian plane.
type Point struct {
	x, y int
}

func (p Point) String() string {
	return strconv.Itoa(p.x) + ":" + strconv.Itoa(p.y)
}

func (p *Point) add(p2 Point) {
	p.x += p2.x
	p.y += p2.y
}

func record(log map[string]int, key string) {
	if _, prs := log[key]; prs {
		log[key]++
	} else {
		log[key] = 1
	}
}

func move(r rune) (Point, error) {
	p := Point{}

	if r == '^' {
		p.y = -1
	}

	if r == 'v' {
		p.y = 1
	}

	if r == '<' {
		p.x = -1
	}

	if r == '>' {
		p.x = 1
	}

	if p.x == 0 && p.y == 0 {
		return p, fmt.Errorf("Invalid move: %v", r)
	}

	return p, nil
}

// Part1 returns the output for Day 03 part 1.
func Part1(s string) int {
	book := make(map[string]int)

	santa := Point{}
	record(book, santa.String())

	for _, r := range s {
		p, err := move(r)
		if err != nil {
			log.Fatal(err)
		}

		santa.add(p)
		record(book, santa.String())
	}

	return len(book)
}

// Part2 returns the output for Day 03 part 2.
func Part2(s string) int {
	book := make(map[string]int)

	santa := Point{}
	record(book, santa.String())

	roboSanta := Point{}
	record(book, roboSanta.String())

	for i, r := range s {
		p, err := move(r)
		if err != nil {
			log.Fatal(err)
		}

		if i%2 == 0 {
			santa.add(p)
			record(book, santa.String())
		} else {
			roboSanta.add(p)
			record(book, roboSanta.String())
		}
	}

	return len(book)
}
