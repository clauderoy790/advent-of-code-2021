package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
	"strconv"
	"strings"
)

var pairs map[string]int
var instructions map[string]string

func main() {
	strs := helpers.GetInputStrings("day14")
	tp := parseInput(strs)
	part1(10, tp)
	fmt.Println("FINAL PAIRAS: ", pairs)
}

// Template:     NNCB NN: 1, NC: 1, CB: 1 | N: 2, C:1, B:1
// After step 1: NCNBCHB
// After step 2: NBCCNBBBCBHCB BB:2 BC:2 BH: 1 CB:2 CC: 1 CN:1 HC:1 NB:2 | N:2 B:6 C:4 H:1
// After step 3: NBBBCNCCNBBNBNBBCHBHHBCHB nb bb bc bn nc cn bb nb bn nb bb ch
// After step 4: NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB
var steps = 0

func part1(nbSteps int, template string) {
	steps = nbSteps
	pairs = stringToPairs(template)
	for i := 0; i < steps; i++ {
		addPairs()
		applyRules()
		// fmt.Printf("After step %v: %v\n", (i + 1), newStr)
		fmt.Printf("After step %v\n", (i + 1))
	}

	most, least := getMostLeast(pairs)
	fmt.Println("most: ", most, ", least: ", least)
	fmt.Println("P1: ", (most - least))
}

func stringToPairs(str string) map[string]int {
	newPairs := map[string]int{}

	for i := 0; i < len(str)-1; i++ {
		newPairs[str[i:i+2]]++
	}
	return newPairs
}

func addPairs() {
	var p = map[string]int{}
	// loop through pairs
	for k, _ := range pairs {
		// loop through instructions to modify string
		for from, to := range instructions {
			if k == from {
				// pairs[k]++
				str := insertStringAt(from, to, 1)
				newPairs := stringToPairs(str)
				for k, v := range newPairs {
					p[k] += v
				}
				break
			}
		}
	}

	// add new pairs
	pairs = p
	fmt.Println("NEW PAIRS: ", pairs)
}

func applyRules() {

}

func getMostLeast(pairs map[string]int) (uint64, uint64) {
	most, least := uint64(1), uint64(10000000000000000000)
	uniqueLetters := findUniqueLetters(pairs)

	count := map[rune]int{}
	for _, r := range uniqueLetters {
		for k, v := range pairs {
			if strings.ContainsRune(k, r) {
				count[r] += v
			}
		}
	}

	//divide
	for k, v := range count {
		ct := v / 2
		if v%2 != 0 {
			ct++
		}
		count[k] = ct
	}

	// count least/most
	for _, v := range count {
		val := uint64(v)
		if val < least {
			least = uint64(v)
		}
		if val > most {
			most = val
		}
	}

	return most, least
}

func findUniqueLetters(pairs map[string]int) []rune {
	runes := []rune{}
	for k, _ := range pairs {
		for _, ru := range k {
			contains := false
			for _, r := range runes {
				if r == ru {
					contains = true
					break
				}
			}
			if !contains {
				runes = append(runes, ru)
			}

		}
	}
	return runes
}

var maxInstructionLength = 0

func addInstruction(key, val string) {
	if len(key) > maxInstructionLength {
		maxInstructionLength = len(key)
	}
	instructions[key] = val
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

func findMostCommonLetter(pairs map[string]int) map[rune]uint64 {
	m := make(map[rune]uint64)
	most := make(map[rune]uint64)
	for k, v := range pairs {
		for _, r := range k {
			m[r] += uint64(v)
		}
	}
	fmt.Println("appearaces: ", m)

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
	str := strings.TrimSpace(strs[0])
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
	return str
}
