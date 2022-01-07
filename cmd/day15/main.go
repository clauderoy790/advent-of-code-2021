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
var lowestRisk = 1000000000000000000

func part1(g helpers.Grid) {
	paths := [][]*helpers.Point{}
	// loop grid
	p := calculatePathFrom([]*helpers.Point{g.Points[0][0]}, g.GetAdjacents(0, 0))
	paths = append(paths, p)

	// find lowest
	lowestRisk := 1000000000000000000
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

func calculateRisk(points []*helpers.Point) int {
	risk := 0
	for _, p := range points {
		risk += p.Val
	}
	return risk
}

func calculatePathFrom(previousRisk int, from *helpers.Point) []*helpers.Point {
	p := []*helpers.Point{}
	p = append(p, previous...)

	for i, path := range p {
		p[i] = append(previous, path[i])
	}
	return p
}

func hasAlreadyCalculatedFrom(p *helpers.Point) bool {
	for _, path := range calculatedFrom {
		if path.X == p.X && path.Y == p.Y {
			return true
		}
	}
	return false
}

// find coming from point
// pass current risk level
