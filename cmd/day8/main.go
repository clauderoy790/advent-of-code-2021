package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var lines []line
	strs := helpers.GetInputStrings("day8")
	for _, str := range strs {
		l := makeLine(str)
		lines = append(lines, l)
	}
	// part1(lines)
	part2(lines)
}

func part1(lines []line) {
	m := make(map[int]int)
	uniques := []int{1, 4, 7, 8}
	for _, n := range uniques {
		m[n] = 0
	}

	count := 0
	for _, line := range lines {
		for _, digit := range line.digits {
			l := getNumber(digit)
			if !skip(l) {
				count++
				m[l]++
			}
		}
	}
	fmt.Println("count: ", count)
	fmt.Println("digits")
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
}

var mapped map[rune]rune
var numbers map[int]string
var seg []string

func part2(lines []line) {
	mapped = make(map[rune]rune)
	sum := 0
	// determine numbers
	for _, line := range lines {
		resetToMap()
		numbers = make(map[int]string)
		seg = line.segments
		mapNumbers()
		mapRunes()
		sum += calculateDigits(line.digits)

	}

	fmt.Println("P2: ", sum)
}

func calculateDigits(strs []string) int {
	finalStr := ""
	if len(nbKeys) != 0 {
		panic(fmt.Sprintf("some numbers are still need mapping: %v\n", nbKeys))
	}
	for _, str := range strs {
		found := false
		str = sortString(str)

		for k, v := range numbers {
			if v == str {
				found = true
				ch := strconv.Itoa(k)
				finalStr += ch
			}
		}

		if !found {
			panic(fmt.Sprintf("failed to find proper string for: %v\n", str))
		}
	}
	final, err := strconv.Atoi(finalStr)
	if err != nil {
		panic("failed to convert to int: " + finalStr)
	}
	return final
}

func mapNumbers() {
	// Map initial numbers
	for _, str := range seg {
		switch len(str) {
		case 2:
			mapNumber(1, str)
		case 3:
			mapNumber(7, str)
		case 4:
			mapNumber(4, str)
		case 7:
			mapNumber(8, str)
		}
	}

	remain := ""
	sixes := findStringsWithLength(seg, 6)
	fives := findStringsWithLength(seg, 5)
	// string of length 5
	for _, s := range fives {
		// 3
		if containsString(s, numbers[1]) {
			mapNumber(3, s)
		} else {
			remain = removeFromString(s, numbers[4])
			if len(remain) == 2 {
				// 5
				mapNumber(5, s)
			} else {
				// 2
				mapNumber(2, s)
			}
		}
	}

	// string of length 6
	for _, s := range sixes {
		if containsString(s, numbers[7]) {
			if containsString(s, numbers[4]) {
				// 9
				mapNumber(9, s)
			} else {
				// 0
				mapNumber(0, s)
			}
		} else {
			// 6
			mapNumber(6, s)
		}
	}
}

func mapRunes() {
	// A
	remain := removeFromString(numbers[7], numbers[1])
	mapRune('a', remain)

	// B
	remain = removeFromString(numbers[4], numbers[3])
	mapRune('b', remain)

	// C
	remain = removeFromString(numbers[8], numbers[6])
	mapRune('c', remain)

	// D
	remain = removeFromString(numbers[8], numbers[0])
	mapRune('d', remain)

	// E
	remain = removeFromString(numbers[8], numbers[9])
	mapRune('e', remain)

	// F
	remain = removeFromString(numbers[7], numbers[2])
	mapRune('f', remain)

	// G
	remain = removeFromString(numbers[9], numbers[4])
	remain = removeFromString(remain, numbers[7])
	mapRune('g', remain)
}

func mapNumber(nb int, str string) {
	if _, ok := numbers[nb]; ok {
		panic(fmt.Sprintf("nb: %v is already mapped to :%v\n", nb, numbers[nb]))
	}

	for i, n := range nbKeys {
		if n == nb {
			nbKeys = append(nbKeys[:i], nbKeys[i+1:]...)
		}
	}
	numbers[nb] = sortString(str)
}

var toMap []rune
var nbKeys []int

func resetToMap() {
	toMap = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	nbKeys = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func mapRune(r1 rune, str string) []rune {
	if len(str) != 1 {
		panic(fmt.Sprintf("wrong length: failed to map: %s to %s\n", string(r1), str))
	}
	r2 := []rune(str)[0]
	mapped[r1] = r2
	removed := false

	for i, r := range toMap {
		if r == r2 {
			toMap = append(toMap[:i], toMap[i+1:]...)
			removed = true
			break
		}
	}

	if !removed {
		panic(fmt.Sprintf("failed to map: %s to %s\n", string(r1), string(r2)))
	}

	return toMap
}

func containsString(original, substr string) bool {
	for _, r := range substr {
		if !strings.Contains(original, string(r)) {
			return false
		}
	}
	return true
}

func removeFromString(original, substr string) string {
	for _, r := range substr {
		if strings.ContainsRune(original, r) {
			ind := strings.Index(original, string(r))
			s := []rune(original)
			original = string(append(s[:ind], s[ind+1:]...))
		}
	}

	return original
}

func findStringsWithLength(strs []string, length int) []string {
	var final []string
	for _, str := range strs {
		if len(str) == length {
			final = append(final, str)
		}
	}
	return final
}

type T struct {
	r []rune
}

func (t *T) Swap(i, j int) {
	t.r[i], t.r[j] = t.r[j], t.r[i]
}
func (t *T) Less(i, j int) bool {
	return uint32(t.r[i]) < uint32(t.r[j])
}
func (t *T) Len() int {
	return len(t.r)
}
func sortString(str string) string {
	runes := []rune(str)
	t := T{runes}
	sort.Sort(&t)
	return string(t.r)
}


func isUniqueLen(str string) bool {
	l := len(str)
	return l == 1 || l == 4 || l == 7 || l == 8
}

var segments map[int]string
var mappedSegments map[int]string

func skip(nb int) bool {
	return nb != 1 && nb != 4 && nb != 7 && nb != 8
}
func getNumber(str string) int {
	for k, v := range segments {
		if len(str) == len(v) {
			return k
		}
	}
	return 0
}

type line struct {
	segments []string
	digits   []string
}

func makeLine(str string) line {
	l := line{}
	strs := strings.Split(str, "|")
	f := make([][]string, 2)
	for i := 0; i < 2; i++ {
		strs[i] = strings.TrimSpace(strs[i])
		f[i] = strings.Split(strs[i], " ")
	}
	l.segments = f[0]
	l.digits = f[1]

	return l
}

// func createSegments() map[int]string {
// 	segments[0] = "abcefg"
// 	segments[1] = "cf"
// 	segments[2] = "acdeg"
// 	segments[3] = "acdfg"
// 	segments[4] = "bcdf"
// 	segments[5] = "abdfg"
// 	segments[6] = "abdefg"
// 	segments[7] = "acf"
// 	segments[8] = "abcdefg"
// 	segments[9] = "abcdfg"
// 	return segments
// }
