package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
	"sort"
)

func main() {
	mapCharacters()
	part2()
}

var scores map[rune]int
var mapped map[rune]rune

func mapCharacters() {
	mapped = make(map[rune]rune)
	mapped['('] = ')'
	mapped['['] = ']'
	mapped['{'] = '}'
	mapped['<'] = '>'

	scores = make(map[rune]int)
	scores[')'] = 1
	scores[']'] = 2
	scores['}'] = 3
	scores['>'] = 4
}

var openings map[int]rune

func part1() {
	strs := helpers.GetInputStrings("day10")
	score := 0
	for _, str := range strs {
		openings = make(map[int]rune)
		if s, ok := isCorrupted(str); ok {
			score += s
		}
	}
	fmt.Println("SCORE: ", score)
}

func part2() {
	strs := helpers.GetInputStrings("day10")
	var scores []int
	for _, str := range strs {
		openings = make(map[int]rune)
		if s, ok := isIncomplete(str); ok {
			scores = append(scores, s)
		}
	}
	sort.Ints(scores)
	fmt.Println("part 2: ", scores[len(scores)/2])
}

func isCorrupted(str string) (int, bool) {
	openings := make(map[int]rune)

	// while all the characters arent all opening
	for countOpening(str) != len(str) {
		// loop through the string and find openings
		for i, r := range str {
			if isOpening(r) {
				openings[i] = r
			} else {
				// check if character before is opening
				if i > 0 && isOpening(rune(str[i-1])) {
					previous := rune(str[i-1])
					matching := isMatchingClosing(previous, r)
					if !matching {
						return getScore(r), true
					} else {
						// remove 2 characters
						runes := []rune(str)
						runes = append(runes[:i-1], runes[i+1:]...)
						str = string(runes)
						break
					}
				}

			}
		}
	}

	return 0, false
}

func isIncomplete(str string) (int, bool) {
	openings := make(map[int]rune)

	// while all the characters arent all opening
	for countOpening(str) != len(str) {
		// loop through the string and find openings
		for i, r := range str {
			if isOpening(r) {
				openings[i] = r
			} else {
				// check if character before is opening
				if i > 0 && isOpening(rune(str[i-1])) {
					previous := rune(str[i-1])
					matching := isMatchingClosing(previous, r)
					if !matching {
						// is corrupted
						return 0, false
					} else {
						// remove 2 characters
						runes := []rune(str)
						runes = append(runes[:i-1], runes[i+1:]...)
						str = string(runes)
						break
					}
				}

			}
		}
	}
	closing := ""
	for i := len(str) - 1; i >= 0; i-- {
		closing += string(getMatchingClosing(rune(str[i])))
	}

	// calculate score
	score := 0
	for _, s := range closing {
		score *= 5
		score += getScore(s)
	}

	return score, true
}

func getScore(r rune) int {
	if s, ok := scores[r]; ok {
		return s
	}
	panic("failed to find score")
}

func findMatchingOpening(str string, i int, closing rune) (rune, error) {
	if !isClosing(closing) {
		return ' ', fmt.Errorf("%v is not a closing char\n", closing)
	}
	
	return ' ', nil
}

func countOpening(str string) int {
	count := 0
	for _, r := range str {
		if isOpening(r) {
			count++
		}
	}
	return count
}

func isOpening(r rune) bool {
	for k, _ := range mapped {
		if r == k {
			return true
		}
	}
	return false
}

func getMatchingClosing(r rune) rune {
	for k, v := range mapped {
		if r == k {
			return v
		}
	}
	panic("failed to find matching closing character for : " + string(r))
}

func isClosing(r rune) bool {
	for _, v := range mapped {
		if r == v {
			return true
		}
	}
	return false
}

func isMatchingClosing(r, closing rune) bool {
	return mapped[r] == closing
}
