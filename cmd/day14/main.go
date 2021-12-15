package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
	"strconv"
	"strings"
)

var template string
var instructions []instruction

func main() {
	strs := helpers.GetInputStrings("day14")
	tp := parseInput(strs)
	part1(40, tp)
}

var steps = 0

func part1(nbSteps int, template string) {
	steps = nbSteps
	newStr := template
	for i := 0; i < steps; i++ {
		toApply := []instruction{}
		temp := newStr
		for j := 0; len(temp) >= 2; j += 2 / 2 {
			if ins := getInstruction(temp[0:2]); ins != nil {
				toApply = append(toApply, instruction{from: ins.from, to: ins.to, index: j + 2/2})
			}
			temp = temp[2/2:]
		}
		// toApply = sortApply(toApply, newStr)
		//apply
		nbInserted := 0
		for _, ins := range toApply {
			newStr = insertStringAt(newStr, ins.to, ins.index+nbInserted)
			nbInserted += len(ins.to)
		}
		fmt.Printf("After step %v: %v\n", (i + 1), "coucou")
	}

	m := findMostCommonLetter(newStr)
	if len(m) != 2 {
		panic("invalid common letters length")
	}
	most, least := -1, -1
	for _, v := range m {
		if most == -1 {
			most = v
			continue
		}
		if least == -1 {
			least = v
			continue
		}
	}

	if most < least {
		most, least = least, most
	}
	fmt.Println("most: ", m)
	fmt.Println("most: ", most, ", least: ", least)

	fmt.Println("P1: ", (most - least))
}

func getInstruction(str string) *instruction {
	for _, ins := range instructions {
		if ins.from == str {
			return &ins
		}
	}
	return nil
}

func getSmallestTemplateLength() int {
	smallest := 100000
	for _, ins := range instructions {
		if len(ins.from) < smallest {
			smallest = len(ins.from)
		}
	}
	return smallest
}

func insertStringAt(str string, sub string, pos int) string {
	if pos < 0 || pos >= len(str) {
		panic("invalid pos: " + strconv.Itoa(pos))
	}
	runes := []rune(str)
	temp := append(runes[:pos], []rune(sub)...)
	runes = []rune(str)
	temp = append(temp, runes[pos:]...)

	return string(temp)
}

func sortApply(ins []instruction, str string) []instruction {
	sorted := []instruction{}
	for len(ins) > 0 {
		smallest := 10000
		index := 0
		for i, in := range ins {
			ind := strings.Index(str, in.from)
			if ind == -1 {
				panic("not supped to has -1 index here")
			}
			if ind < smallest {
				smallest = ind
				index = i
			}
		}

		sorted = append(sorted, ins[index])
		ins = append(ins[:index], ins[index+1:]...)
	}
	// fmt.Println("sorted: ", sorted)
	return sorted
}

func findMostCommonLetter(str string) map[rune]int {
	m := make(map[rune]int)
	most := make(map[rune]int)
	for _, ru := range str {
		m[ru]++
	}

	r := ' '
	c := 0
	// find most common
	for k, v := range m {
		if v > c {
			c = v
			r = k
		}
	}
	most[r] = c

	r = ' '
	c = 1000000
	// find least
	for k, v := range m {
		if v < c {
			c = v
			r = k
		}
	}
	most[r] = c
	return most
}

func parseInput(strs []string) string {
	template = strings.TrimSpace(strs[0])
	for _, str := range strs {
		if len(str) == 0 || !strings.Contains(str, "->") {
			continue
		}

		sp := strings.Split(str, "->")
		for i := range sp {
			sp[i] = strings.TrimSpace(sp[i])
		}

		ins := instruction{from: sp[0], to: sp[1]}
		instructions = append(instructions, ins)
	}
	return template
}

type instruction struct {
	from, to string
	index    int
}
