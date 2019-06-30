package day02

// Part1 returns the output for Day 02 part 1.
func Part1(s []string) int {
	var sum float64

	for _, desc := range s {
		rp := newRectangularCuboid(desc)
		sum += rp.SurfaceArea()
		sum += rp.SmallestFace().Area()
	}

	return int(sum)
}

// Part2 returns the output for Day 02 part 2.
func Part2(s []string) int {
	var sum float64

	for _, desc := range s {
		rp := newRectangularCuboid(desc)
		sum += rp.Volume()
		sum += rp.SmallestFace().Perimeter()
	}

	return int(sum)
}
