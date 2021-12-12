package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
	"strconv"
)

var steps = 100
var flashes = 0
var width, height int = 10, 10
var points [10][10]*point

func main() {
	strs := helpers.GetInputStrings("day11")
	createPoints(strs)
	// part1()
	part2()
}

func createPoints(strs []string) {
	// height
	for i, str := range strs {
		// width
		for j, r := range str {
			nb, err := strconv.Atoi(string(r))
			if err != nil {
				panic("failed to convert to int: " + string(r))
			}
			pt := &point{x: j, y: i, energy: nb}
			points[j][i] = pt
		}
	}
}

func part1() {
	flashes = 0
	flashB := 0
	// fmt.Println("5,2: ", points[5][2])
	// fmt.Println("2,9: ", points[2][9])
	for i := 0; i < steps; i++ {
		flashB = flashes
		fmt.Println("AFTER ", i, " steps: ", flashes)
		resetFlashed()
		for _, pts := range points {
			for _, pt := range pts {
				if pt.flashed && pt.increased {
					continue
				}
				pt.increaseEnergy()
			}
		}
		fmt.Println("FLASHES THIS TURN: ", (flashes - flashB))
	}

	fmt.Println("P1: ", flashes)
}

func part2() {
	step := 0
	flash = make(map[string]bool)
	initFlashed()
	for !allFlashed() {
		initFlashed()
		resetFlashed()
		for _, pts := range points {
			for _, pt := range pts {
				if pt.flashed && pt.increased {
					continue
				}
				pt.increaseEnergy()
			}
		}
		step++
	}

	fmt.Println("P2: ", step)
}

var flash map[string]bool

func allFlashed() bool {
	for _, v := range flash {
		if !v {
			return false
		}
	}
	return true
}

func initFlashed() {
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			flash[getKey(points[i][j])] = false
		}
	}
}

func flashed(p *point) bool {
	return flash[getKey(p)]
}

func getKey(p *point) string {
	str := fmt.Sprintf("%v%v", p.x, p.y)
	return str
}

func resetFlashed() {
	for _, pts := range points {
		for _, pt := range pts {
			pt.resetFlash()
		}
	}
}

type point struct {
	x, y, energy int
	flashed      bool
	increased    bool
}

func (p *point) resetFlash() {
	p.flashed = false
	p.increased = false
}

func (p *point) increaseEnergy() {
	// can only flash once per turn
	if p.flashed {
		return
	}
	// if p.increased {
	// 	return
	// }
	p.increased = true
	p.energy++

	// flash
	if p.energy > 9 {
		p.energy = 0
		flashes++
		p.flashed = true
		flash[getKey(p)] = true
		adjs := p.getAdjacents()
		for _, adj := range adjs {
			adj.increaseEnergy()
		}
	}
}

var adjacents = []point{{x: -1, y: -1}, {x: -1, y: 1}, {x: -1, y: 0}, {x: 0, y: 1}, {x: 0, y: -1}, {x: 1, y: 1}, {x: 1, y: 0}, {x: 1, y: -1}}

func (p *point) getAdjacents() []*point {
	var adjs []*point
	for _, adj := range adjacents {
		pt := point{x: p.x + adj.x, y: p.y + adj.y}
		if pt.isInBounds() {
			adjs = append(adjs, points[pt.x][pt.y])
		}
	}

	return adjs
}

func (p *point) isInBounds() bool {
	return p.x >= 0 && p.x < width && p.y >= 0 && p.y < height
}
