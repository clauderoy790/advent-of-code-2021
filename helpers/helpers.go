package helpers

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func GetInputStrings(dir string) []string {
	_, file, _, _ := runtime.Caller(0)
	b, err := os.ReadFile(strings.Replace(file, "/helpers/helpers.go", fmt.Sprintf("/cmd/%s/input.txt", dir), 1))
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(b), "\n")
	for i, _ := range strs {
		strs[i] = strings.TrimSpace(strs[i])
	}
	return strs
}

func ConvertStringToInts(str string, sep string) []int {
	str = strings.TrimSpace(str)
	sl := strings.Split(str, sep)
	nbs := make([]int, 0)

	for _, s := range sl {
		nb, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			continue
		}
		nbs = append(nbs, nb)
	}
	return nbs
}

func ConvertStringsToInts(strs []string) []int {
	nbs := []int{}
	for _, s := range strs {
		nb, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			continue
		}
		nbs = append(nbs, nb)
	}
	return nbs
}

func ReadInputToGrid(dir string) Grid {
	pts := [][]*Point{}
	strs := GetInputStrings(dir)
	height, width := len(strs), len(strs[0])
	for i := 0; i < height; i++ {
		row := []*Point{}
		for j := 0; j < width; j++ {
			val, err := strconv.Atoi(string(strs[i][j]))
			if err != nil {
				panic("cannot conv to str: " + string(strs[i][j]))
			}
			pt := &Point{j, i, val}
			row = append(row, pt)
		}
		pts = append(pts, row)
	}
	return Grid{width, height, pts}
}

type Grid struct {
	Width, Height int
	Points        [][]*Point
}

var adjacentsDiag = []Point{{X: -1, Y: -1}, {X: -1, Y: 1}, {X: -1, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: -1}, {X: 1, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: -1}}
var adjacents = []Point{{X: -1, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: -1}, {X: 1, Y: 0}}

func (g *Grid) GetAdjacents(x, y int) []*Point {
	if x >= 0 && x < g.Width && y >= 0 && y < g.Height {
		panic("GetAdjacents: invalid sarting point")
	}
	var adjs []*Point
	p := g.Points[x][y]
	for _, adj := range adjacents {
		pt := Point{X: p.X + adj.X, Y: p.Y + adj.Y}
		if pt.isInBounds(g) {
			adjs = append(adjs, g.Points[pt.X][pt.Y])
		}
	}
	return adjs
}

type Point struct {
	X, Y, Val int
}

func (p *Point) String() string {
	return fmt.Sprintf("X: %v, Y: %v, Val: %v", p.X, p.Y, p.Val)
}

func (p *Point) isInBounds(g *Grid) bool {
	return p.X >= 0 && p.X < g.Width && p.Y >= 0 && p.Y < g.Height
}
