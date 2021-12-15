package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var template string
var instructions map[string]string

func main() {
	strs := helpers.GetInputStrings("day14")
	tp := parseInput(strs)
	part1(10, tp)
}

var steps = 0

func part1(nbSteps int, template string) {
	steps = nbSteps
	newStr := template
	for i := 0; i < steps; i++ {
		toApply := []instruction{}
		temp := newStr
		// iterate string
		pattern := ""
		longerTo := ""
		for j := 0; len(temp) >= 2; j++ {
			from, to, err := getInstruction(temp)
			if err == nil {
				pattern += temp[0:2]
				longerTo += to
				if _, ok := instructions[pattern]; !ok {
					fmt.Printf("adding new instruction: %v: %v\n", pattern, longerTo)
					addInstruction(pattern, longerTo)
				}
				toApply = append(toApply, instruction{from: from, to: to, index: j + len(from)/2})
			} else {
				pattern = ""
				longerTo = ""
			}
			temp = temp[len(from)/2:]
		}

		//apply
		nbInserted := 0
		for _, ins := range toApply {
			newStr = insertStringAt(newStr, ins.to, ins.index+nbInserted)
			nbInserted += len(ins.to)
		}
		// fmt.Printf("After step %v: %v\n", (i + 1), newStr)
		fmt.Printf("After step %v\n", (i + 1))
	}

	m := findMostCommonLetter(newStr)
	if len(m) != 2 {
		panic("invalid common letters length")
	}
	most, least := uint64(1), uint64(1)
	for _, v := range m {
		if most == uint64(1) {
			most = v
			continue
		}
		if least == uint64(1) {
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

var maxInstructionLength = 0

func addInstruction(key, val string) {
	if len(key) > maxInstructionLength {
		maxInstructionLength = len(key)
	}
	instructions[key] = val
}

func getInstruction(str string) (string, string, error) {
	max := len(str)
	if max%2 != 0 {
		max--
	}

	if maxInstructionLength < max {
		max = maxInstructionLength
	}
	for keyLen := max; keyLen >= 2; keyLen -= 2 {
		key := str[:keyLen]
		for k, v := range instructions {
			if key == k {
				return k, v, nil
			}
		}
	}
	return "", "", errors.New("failed to find")
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

func findMostCommonLetter(str string) map[rune]uint64 {
	m := make(map[rune]uint64)
	most := make(map[rune]uint64)
	for _, ru := range str {
		m[ru]++
	}

	r := ' '
	c := uint64(0)
	// find most common
	for k, v := range m {
		if v > c {
			c = v
			r = k
		}
	}
	most[r] = c

	r = ' '
	c = uint64(10000000000000000000)
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
	instructions = make(map[string]string)
	for _, str := range strs {
		if len(str) == 0 || !strings.Contains(str, "->") {
			continue
		}

		sp := strings.Split(str, "->")
		for i := range sp {
			sp[i] = strings.TrimSpace(sp[i])
		}
		addInstruction(sp[0], sp[1])
	}
	fmt.Println("instructions: ", instructions)
	return template
}

type instruction struct {
	from, to string
	index    int
}
