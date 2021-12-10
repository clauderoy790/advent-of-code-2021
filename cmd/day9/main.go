package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
	"sort"
	"strconv"
)

var points [][]*point
var width, height int

func main() {
	strs := helpers.GetInputStrings("day9")
	width = len(strs[0])
	height = len(strs)
	points = make([][]*point, width)
	for i := 0; i < width; i++ {
		points[i] = make([]*point, height)
		for j := 0; j < height; j++ {
			pt := point{x: i, y: j}
			val := string(strs[pt.y][pt.x])
			pt.setHeight(string(val))
			points[pt.x][pt.y] = &pt
		}
	}
	part1()
	part2()

}

func part2() {
	basinLocations = make(map[string]bool)
	basins := make([]*basin, 0)
	for _, pts := range points {
		for _, pt := range pts {
			if pt.isLowPoint() {
				b := basin{}
				b.lowPoint = pt
				basins = append(basins, &b)
			}
		}
	}
	// calculate basins
	for _, basin := range basins {
		basin.calculatePoints()
	}

	// Count top 3
	top3 := make([]*basin, 0)
	for _, basin := range basins {
		if len(top3) < 3 {
			top3 = append(top3, basin)
			sort.Sort(B(top3))
		} else {
			for i, top := range top3 {
				if basin.size() > top.size() {
					top3[i] = basin
					break
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

type B []*basin

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

func isInBasin(pt *point) bool {
	// 9 are never in basins
	if pt.height == 9 {
		return false
	}
	str := strconv.Itoa(pt.x) + strconv.Itoa(pt.y)
	return basinLocations[str]
}

type basin struct {
	points   []*point
	lowPoint *point
}

func (b *basin) calculatePoints() {
	b.calculateRecursive(b.lowPoint)
}

func (b *basin) calculateRecursive(p *point) {
	adjacents := []point{{x: -1, y: 0}, {x: 1, y: 0}, {x: 0, y: 1}, {x: 0, y: -1}}
	for _, ad := range adjacents {
		x, y := p.x+ad.x, p.y+ad.y
		if isInBounds(x, y) {
			newPt := points[x][y]
			if b.canAddPoint(newPt) {
				b.add(newPt)
				b.calculateRecursive(newPt)
			}
		}
	}
}

func (b *basin) canAddPoint(pt *point) bool {
	return pt.height != 9 && !isInBasin(pt)
}

func (b *basin) add(pt *point) {
	if pt.height == 9 {
		panic("height 9 should not be in basin")
	}
	str := strconv.Itoa(pt.x) + strconv.Itoa(pt.y)
	basinLocations[str] = true
	b.points = append(b.points, pt)
}

func (b *basin) size() int {
	return len(b.points)
}

func part1() {
	// find low
	risk := 0
	for _, pts := range points {
		for _, pt := range pts {
			if pt.isLowPoint() {
				risk += pt.height + 1
			}
		}
	}
	fmt.Println("P1: ", risk)
}

func isInBounds(x, y int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

type point struct {
	x, y   int
	height int
}

func (p *point) isInBounds() bool {
	return p.x >= 0 && p.x < width && p.y >= 0 && p.y < height
}

func (p *point) isLowPoint() bool {
	for _, pt := range p.getAdjacents() {
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

func (p *point) String() string {
	return fmt.Sprintf("x:%v, y:%v\n", p.x, p.y)
}

var adjacents = []point{{x: -1, y: 0}, {x: 1, y: 0}, {x: 0, y: 1}, {x: 0, y: -1}}

func (p *point) getAdjacents() []*point {
	var adj []*point
	for _, ad := range adjacents {
		pt := point{x: p.x + ad.x, y: p.y + ad.y}
		if pt.isInBounds() {
			adj = append(adj, points[pt.x][pt.y])
		}
	}
	return adj
}
