package day09

// Part1 returns the output for Day 09 part 1.
func Part1(edges []string) int {
	var g = make(graph)
	g.addEdges(edges)
	return g.getShortestDistance()
}

// Part2 returns the output for Day 09 part 2.
func Part2(edges []string) int {
	var g = make(graph)
	g.addEdges(edges)
	return g.getLongestDistance()
}
