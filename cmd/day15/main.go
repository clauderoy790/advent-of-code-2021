package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
)

func main() {
	g := helpers.ReadInputToGrid("day15")
	fmt.Println(g.Points)
	part1(g)
}

var calculatedFrom []*helpers.Point

func part1(g helpers.Grid) {
	paths := [][]*helpers.Point{}
	// loop grid
	p := calculatePathFrom(nil, g.GetAdjacents(0, 0))
	paths = append(paths, p)

	// find lowest
	lowestRisk := 100000000
	for _, path := range paths {
		risk := 0
		for _, point := range path {
			risk += point.Val
		}
		if risk < lowestRisk {
			lowestRisk = risk
		}
	}
}

func calculatePathFrom(previous, points []*helpers.Point) [][]*helpers.Point {
	paths := [][]*helpers.Point{}
	// add previous points
	for i, path := range paths {
		paths[i] = append(previous, path[i])
	}
	return paths
}

func hasAlreadyCalculatedFrom(p *helpers.Point) bool {
	for _, path := range calculatedFrom {
		if path.X == p.X && path.Y == p.Y {
			return true
		}
	}
	return false
}
