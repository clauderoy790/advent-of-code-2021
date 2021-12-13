package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
	"strconv"
	"strings"
)

var paper map[point]bool
var foldInstructions []foldInstruction
var width, height int
var strs []string

func main() {
	strs = helpers.GetInputStrings("day13")
	makePaper(strs)
	part1()
	part2()
}

func part1() {
	for _, ins := range foldInstructions {
		if ins.up {
			foldUp(ins.val)
		} else {
			foldLeft(ins.val)
		}
	}
	fmt.Println("P1: ", countPaperValues())
}

func part2() {
	for _, ins := range foldInstructions {
		if ins.up {
			foldUp(ins.val)
		} else {
			foldLeft(ins.val)
		}
	}

	// recount the size
	recountSize()

	//print the paper
	for i := 0; i < height; i++ {
		str := ""
		for j := 0; j < width; j++ {
			pt := point{j, i}
			sub := "."
			if paper[pt] {
				sub = "#"
			}
			str += sub
		}
		fmt.Println(str)
	}
}

func countPaperValues() int {
	count := 0
	for _, v := range paper {
		if v {
			count++
		}
	}
	return count
}

func foldLeft(x int) {
	pointsNeededFold := make([]point, 0)

	for k, _ := range paper {
		if k.x > x {
			pointsNeededFold = append(pointsNeededFold, k)
		}
	}

	for _, pt := range pointsNeededFold {
		// delete old key
		delete(paper, pt)

		// count new X position
		newX := x - (pt.x - x)
		newPt := point{newX, pt.y}
		paper[newPt] = true
	}
}

func foldUp(y int) {
	pointsNeededFold := make([]point, 0)

	for k, _ := range paper {
		if k.y > y {
			pointsNeededFold = append(pointsNeededFold, k)
		}
	}

	for _, pt := range pointsNeededFold {
		// delete old key
		delete(paper, pt)

		// count new Y position
		newY := y - (pt.y - y)
		newPt := point{pt.x, newY}
		paper[newPt] = true
	}
}

func makePaper(strs []string) {
	paper = make(map[point]bool)
	makeFolds := false
	w, h := 0, 0
	for _, str := range strs {
		if len(str) == 0 {
			makeFolds = true
			continue
		}
		//
		if makeFolds {
			ins := foldInstruction{}
			ins.up = strings.Contains(str, "y=")
			equalInd := strings.Index(str, "=")
			sub := str[equalInd+1:]
			nb, err := strconv.Atoi(sub)
			if err != nil {
				panic("failed to covert to nb: " + sub)
			}
			ins.val = nb
			foldInstructions = append(foldInstructions, ins)
			continue
		}

		spl := strings.Split(str, ",")
		tempW, err := strconv.Atoi(spl[0])
		if err != nil {
			panic("fail to parse width :" + spl[0])
		}
		tempH, err := strconv.Atoi(spl[1])
		p := point{tempW, tempH}
		paper[p] = true
		if err != nil {
			panic("fail to parse width :" + spl[1])
		}
		if tempW+1 > w {
			w = tempW + 1
		}
		if tempH+1 > h {
			h = tempH + 1
		}
	}
	width = w
	height = h
}

func recountSize() {
	w, h := 0, 0
	for k, v := range paper {
		if v {
			if k.x+1 > w {
				w = k.x + 1
			}
			if k.y+1 > h {
				h = k.y + 1
			}
		}
	}
	width = w
	height = h
}

type point struct {
	x, y int
}

type foldInstruction struct {
	up  bool
	val int
}
