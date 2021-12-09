package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	strs := helpers.GetInputStrings("day9")
	points := make([][]point, len(strs))
	adjacents := []point{{x: -1, y: 0}, {x: 1, y: 0}, {x: 0, y: 1}, {x: 0, y: -1}}
	for i, str := range strs {
		points[i] = make([]point, len(str))
		for j := 0; j < len(str); j++ {
			pt := point{x: i, y: j}
			pt.setHeight(string(strs[i][j]))
			for _, adj := range adjacents {
				x, y := i+adj.x, j+adj.y
				if x >= 0 && x < len(strs) && y >= 0 && y < len(str) {
					newPt := point{x: x, y: y}
					newPt.setHeight(string(strs[x][y]))
					pt.addAdjacent(newPt)
				}
			}
			points[i][j] = pt
		}
	}
	part1(points)

}

func part2(points [][]point) {
	basinLocations = make(map[string]bool)
	basins := []basin{}
	for _, pts := range points {
		for _, pt := range pts {
			if pt.isLowPoint() {
				b := basin{}
				basins = append(basins, b)
			}
		}
	}

	// Count top 3
	top3 := make([]basin, 3)
	for _, basin := range basins {
		if len(top3) < 3 {
			top3 = append(top3, basin)
			sort.Sort(B(top3))
		} else {
			for i, top := range top3 {
				if basin.size() > top.size() {
					top3[i] = basin
				}
			}
		}
	}

	// Multiply
	p2 := 1
	for _, b := range top3 {
		p2 *= b.size()
	}
	fmt.Println("P2: ", p2)
}

type B []basin

func (b B) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b B) Len() int {
	return len(b)
}

func (b B) Less(i, j int) bool {
	return b[i].size() < b[j].size()
}

var basinLocations map[string]bool

func addToBasin(pt point) {
	if pt.height == 9 {
		panic("height 9 should not be in basin")
	}
	str := strconv.Itoa(pt.x) + strconv.Itoa(pt.y)
	basinLocations[str] = true
}

func isInBasin(pt point) bool {
	// 9 are never in basins
	if pt.height == 9 {
		return false
	}
	str := strconv.Itoa(pt.x) + strconv.Itoa(pt.y)
	return basinLocations[str]
}

type basin struct {
	points []point
}

func (b *basin) calculatePoints(points [][]point) {
}

func (b *basin) size() int {
	return len(b.points)
}

func part1(points [][]point) {
	// find low
	risk := 0
	for _, pts := range points {
		for _, pt := range pts {
			if pt.isLowPoint() {
				risk += pt.height + 1
			}
		}
	}
	fmt.Println("P!: ", risk)
}

type point struct {
	x, y      int
	height    int
	adjacents []point
}

func (p *point) isLowPoint() bool {
	for _, pt := range p.adjacents {
		if pt.height <= p.height {
			return false
		}
	}
	return true
}

func (p *point) setHeight(height string) {
	if len(height) != 1 {
		panic("invalid height: " + height)
	}
	nb, err := strconv.Atoi(height)
	if err != nil {
		panic("failed to convert height to int: " + height)
	}
	p.height = nb
}

func (p *point) addAdjacent(pt point) {
	p.adjacents = append(p.adjacents, pt)
}
