package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
	"strconv"
	"strings"
)

var pairs map[string]uint64
var instructions map[string]string

func main() {
	strs := helpers.GetInputStrings("day14")
	tp := parseInput(strs)
	part2(40, tp)
	// real answers
	// 2188189693529
	// my answers
	// 2188189693529
	// after 40 steps:
	// 2188189693529
	// 2911561572630
}

// Template:     NNCB NN: 1, NC: 1, CB: 1 | N: 2, C:1, B:1
// After step 1: NCNBCHB | N:2 H:1
// After step 2: NBCCNBBBCBHCB BB:2 BC:2 BH: 1 CB:2 CC: 1 CN:1 HC:1 NB:2 | N:2 B:6 C:4 H:1
// After step 3: NBBBCNCCNBBNBNBBCHBHHBCHB nb bb bc bn nc cn bb nb bn nb bb ch | B: 10 H: 3 | BB:
// After step 4: NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB | N: B: H:

func part2(steps int, template string) {
	pairs = stringToPairs(template)
	for i := 0; i < steps; i++ {
		pairs = processInstructions()
		fmt.Printf("after %v step\n", (i + 1))
		printRunes()
	}

	// count runes
	count := getRuneCount()
	_, most := countMostRune(count)
	_, least := countLeastRune(count)
	fmt.Println("P2: ", (most - least))
}

func processInstructions() map[string]uint64 {
	p := map[string]uint64{}
	// for k, v := range pairs {
	// 	p[k] = v
	// }

	for k, v := range pairs {
		if ins, ok := instructions[k]; ok {
			newStr := insertStringAt(k, ins, 1)
			newP := stringToPairs(newStr)
			// multiply pair by values
			for k2, v2 := range newP {
				p[k2] += v2 * v
			}
		}
	}
	return p
}

func printRunes() {
	runeCount := getRuneCount()
	rm, m := countMostRune(runeCount)
	rl, l := countLeastRune(runeCount)
	fmt.Printf("The most present letter is %v with %v times and the least present letter is %v with %v times \n", rm, m, rl, l)

}

func stringToPairs(str string) map[string]uint64 {
	newPairs := map[string]uint64{}
	for i := 0; i < len(str)-1; i++ {
		newPairs[str[i:i+2]]++
	}
	return newPairs
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

func getRuneCount() map[string]uint64 {
	m := map[string]uint64{}

	for k, v := range pairs {
		for _, r := range k {
			m[string(r)] += v
		}
	}

	//real count
	for k, v := range m {
		if v%2 != 0 {
			m[k]++
		}

		m[k] /= 2
	}

	return m
}

func countMostRune(runes map[string]uint64) (string, uint64) {
	r := ""
	m := uint64(0)
	for k, v := range runes {
		if v > m {
			m = v
			r = k
		}
	}
	return r, m
}

func countLeastRune(runes map[string]uint64) (string, uint64) {
	r := ""
	var l uint64 = 10000000000000000000
	for k, v := range runes {
		if v < l {
			l = v
			r = k
		}
	}
	return r, l
}

func findMostCommonLetter(pairs map[string]uint64) map[rune]uint64 {
	m := make(map[rune]uint64)
	most := make(map[rune]uint64)
	for k, v := range pairs {
		for _, r := range k {
			m[r] += uint64(v)
		}
	}
	// divide
	for k, v := range m {
		if v%2 == 1 {
			m[k]++
		}
		m[k] /= 2
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
		instructions[sp[0]] = sp[1]
	}
	return str
}
