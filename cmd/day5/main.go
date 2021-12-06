package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	lines := readInput()
	part2(lines)
}

func part1(lines []line) {
	lines = removeNonHorizontalVertical(lines)
	m := make(map[point]int)
	nbPt := 0
	for _, l := range lines {
		pts := l.getPoints()
		for _, pt := range pts {
			if _, ok := m[pt]; !ok {
				m[pt] = 0
			}
			m[pt]++
		}
	}
	for _, v := range m {
		if v >= 2 {
			nbPt++
		}
	}
	fmt.Println("resp: ", nbPt)
}

func part2(lines []line) {
	m := make(map[point]int)
	nbPt := 0

	for _, l := range lines {
		pts := l.getPoints()
		if l.isHoriz() {
			// fmt.Println(l)
			// fmt.Println("horiz : ", pts)
		}
		for _, pt := range pts {
			if _, ok := m[pt]; !ok {
				m[pt] = 0
			}
			m[pt]++
		}
	}
	for _, v := range m {
		if v >= 2 {
			nbPt++
		}
	}
	fmt.Println("resp: ", nbPt)
}

func removeNonHorizontalVertical(lines []line) []line {
	var l []line
	for _, line := range lines {
		if line.p1.x == line.p2.x || line.p1.y == line.p2.y {
			l = append(l, line)
		}
	}
	return l
}

func readInput() []line {
	var lines []line
	strs := getInputStrings()
	for _, str := range strs {
		sp := strings.Split(strings.TrimSpace(str), "->")
		points := [2]point{}
		for i, s := range sp {
			s := strings.TrimSpace(s)
			sp2 := strings.Split(s, ",")
			n1, err := strconv.Atoi(sp2[0])
			if err != nil {
				panic(err)
			}
			n2, err := strconv.Atoi(sp2[1])
			if err != nil {
				panic(err)
			}
			points[i] = point{n1, n2}
		}
		lines = append(lines, line{points[0], points[1]})
	}
	return lines
}

func getInputStrings() []string {
	_, file, _, _ := runtime.Caller(0)
	b, err := os.ReadFile(strings.Replace(file, "main.go", "input.txt", 1))
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}

type line struct {
	p1, p2 point
}

func (l *line) isHoriz() bool {
	return l.p1.x != l.p2.x && l.p2.y != l.p1.y
}

func (l *line) getPoints() []point {
	var points []point
	origin := l.p1
	dest := l.p2
	if l.p1.x == l.p2.x {
		if l.p2.y < origin.y {
			origin = l.p2
			dest = l.p1
		}
		for i := origin.y; i <= dest.y; i++ {
			points = append(points, point{l.p1.x, i})
		}
	} else if l.p1.y == l.p2.y {
		if l.p2.x < origin.x {
			origin = l.p2
			dest = l.p1
		}
		for i := origin.x; i <= dest.x; i++ {
			points = append(points, point{i, l.p1.y})
		}
	} else {
		//HORIZONTAL
		if dest.x < origin.x {
			origin = l.p2
			dest = l.p1
		}
		increment := 1
		if dest.y < origin.y {
			increment = -1
		}
		for i := origin.x; i <= dest.x; i++ {
			points = append(points, point{
				x:i, 
				y:origin.y + ((i-origin.x) * increment),
			})
		}
	}
	return points
}

func (l *line) String() string {
	return fmt.Sprintf("P1:(%v,%v)\nP2:(%v,%v)\n", l.p1.x, l.p1.y, l.p2.x, l.p2.y)
}

type point struct {
	x, y int
}
