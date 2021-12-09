package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// segments = createSegments()
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
		mapInitialNumbers()
		mapA()
		mapNumbersAndRunes()
		sum += calculateDigits(line.digits)
		fmt.Printf("mapped: %+v, numbers: %+v\n", mapped, numbers)

	}

	fmt.Println("P2: ", sum)
}

func calculateDigits(strs []string) int {
	str := ""
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
				str += ch
			}
		}

		if !found {
			panic(fmt.Sprintf("failed to find proper string for: %v\n", str))
		}
	}
	final, err := strconv.Atoi(str)
	if err != nil {
		panic("failed to convert to int: " + str)
	}
	return final
}

func mapInitialNumbers() {
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

	fmt.Println("INITIAL numbers: ", numbers)
}

func mapA() {
	removed := removeFromString(numbers[7], numbers[1])
	fmt.Printf("orig: %v, substr: %v, final:%v\n", numbers[7], numbers[1], removed)
	if len(removed) != 1 {
		panic("failed to map A: " + removed)
	}
	fmt.Println("REMOVED: ", removed)
	mapRune('a', removed)
}

func mapNumbersAndRunes() {
	remain := ""
	sixes := findStringsWithLength(seg, 6)
	fives := findStringsWithLength(seg, 5)
	// string of length 5
	for _, s := range fives {
		// 3
		if containsString(s,numbers[1]) {
			mapNumber(3,s)
		} else {
			remain = removeFromString(s,numbers[4])
			if len(remain) == 2 {
				// 5
				mapNumber(5,s)
			} else {
				// 2
				mapNumber(2,s)
			}
		}
	}

	// string of length 6
	for _, s := range sixes {
		if containsString(s,numbers[7]) {
			if containsString(s,numbers[4]) {
				// 9
				mapNumber(9,s)
			} else {
				// 0
				mapNumber(0,s)
			}
		} else {
			// 6
			mapNumber(6,s)
		}
	}

	sim := getSimilarCharactersCountArray(sixes, numbers[4], 3)
	// find the one that doesn't include all characters that 1 has (it's 6)
	found := false
	for _, str := range sim {
		if !containsString(str, numbers[7]) {
			// 6
			mapNumber(6, str)
			remain := removeFromString(numbers[7], str)

			// Map c and f
			mapRune('c',remain)
			remain = removeFromString(numbers[1],remain)
			mapRune('f',remain)

			// find 9
			for _, s := range sim {
				if !containsString(s,numbers[4]) && s != str {
					mapNumber(9,s)
					break
				}
			}
			// find 0
			found = false
			for _, s := range sixes {
				if s != numbers[9] && s != numbers[6] {
					mapNumber(0,s)
				}
			}

			// D
			remain = removeFromString(numbers[4],numbers[0])
			mapRune('d',remain)


			if !found {
				panic("failed to find 0")
			}
			mapRune('d')
			// map F
			strings
			for _, r := range str {
				// todo hereh
				if !strings.ContainsRune() {

				}
			}
			break
		}
	}
	if !found {
		panic("failed to find 6")
	}
	zeroStr := sim[0]
	mapNumber(0, zeroStr)
	//find the character that is in 4 that isn't in 1
	str := getSimilarString(numbers[4], numbers[1])
	str = removeFromString(str, numbers[1])
	if len(str) != 1 {
		panic("failed to find the character that should map to D")
	}

	mapRune('d', str)

	// Map B
	s := numbers[1] + string(str[0])
	n := findDifferentCharacters(numbers[4], s)
	if len(n) != 1 {
		panic("could not map B properly")
	}
	mapRune('b', str)

	//Map C
	contains := removeStringsContainingString(sixes, numbers[1])
	foundStr := ""
	for _, str := range contains {
		if strings.ContainsRune(str, mapped['d']) {
			foundStr = str
		}
	}
	if foundStr == "" {
		panic("failed to find 6")
	}

	mapNumber(6, foundStr)
	c := removeFromString(numbers[1], numbers[6])
	if len(c) != 1 {
		panic("failed to find C")
	}

	mapRune('c', str)

	//Map F
	f := removeFromString(numbers[1], c)
	if len(f) != 1 {
		panic("failed to find F")
	}

	mapRune('f', str)

	// Find 9
	nine := ""
	for _, str := range sixes {
		if str != numbers[6] && str != numbers[0] {
			nine = str
			break
		}
	}

	if nine == "" {
		panic("failed to find 9")
	}

	mapNumber(9, nine)
	found := false

	for _, r := range nine {
		if leftToMap(r) {
			mapRune('g', string(r))
			found = true
		}
	}
	if !found {
		panic("failed to find in remaing maps: ")
	}

	if len(toMap) != 1 {
		panic("only one rune should remain")
	}

	mapRune('e', string(toMap[0]))

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

func leftToMap(r rune) bool {
	for _, r2 := range toMap {
		if r == r2 {
			return true
		}
	}
	return false
}

var toMap []rune
var nbKeys []int

func resetToMap() {
	toMap = []rune{'a', 'b', 'c', 'd', 'e', 'f'}
	nbKeys = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func mapRune(r1 rune, str string) []rune {
	if len(str) != 1 {
		panic(fmt.Sprintf("wrong length: failed to map: %s to %s\n", string(r1), string(2)))
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

func getSimilarCharactersCount(strs []string, comparison string, count int) map[string]int {
	m := make(map[string]int)
	for _, str := range strs {
		count := countSimilarCharacters(str, comparison)
		m[str] = count
	}
	return m
}

func removeStringsContainingString(strs []string, comparison string) []string {
	final := []string{}
	for _, str := range strs {
		if !strings.Contains(str, comparison) {
			final = append(final, str)
		}
	}
	return final
}

func findDifferentCharacters(str1, str2 string) string {
	diff := str1 + str2
	init := str1
	second := str2
	if len(str2) < len(init) {
		init, second = second, init
	}
	for _, r := range init {
		if strings.ContainsRune(second, r) {
			diff = strings.ReplaceAll(diff, string(r), "")
		}
	}
	return diff
}

func getSimilarCharactersCountArray(strs []string, comparison string, count int) []string {
	final := make([]string, 0)
	m := getSimilarCharactersCount(strs, comparison, count)
	for k, _ := range m {
		final = append(strs, k)
	}
	return final
}

func getSimilarString(str1, str2 string) string {
	str := ""
	for _, r := range str1 {
		if strings.ContainsRune(str2, r) {
			str += string(r)
		}
	}
	return str
}

func findStringsWithSimilarCharactersAs(strs []string, comparison string, count int) []string {
	final := []string{}
	for _, str := range strs {
		ct := countSimilarCharacters(str, comparison)
		if ct == count {
			final = append(final, str)
		}
	}
	return final
}

func countSimilarCharacters(str1, str2 string) int {
	count := 0
	start, compare := str1, str2
	if len(str2) < len(start) {
		start, compare = compare, start
	}

	for _, r := range start {
		if strings.Contains(compare, string(r)) {
			count++
		}
	}

	return count
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
	s := []rune(original)
	indexes := make([]int, 0)
	for _, r := range substr {
		if strings.ContainsRune(original, r) {
			ind := strings.Index(original, string(r))
			indexes = append(indexes, ind)
		}
	}

	for i, ind := range indexes {
		tor := ind - i
		s = append(s[0:tor], s[tor+1:]...)
	}

	return string(s)
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

// func mapSegments(strs []string) map[rune]rune {
// 	m := make(map[rune]rune)
// 	original
// 	for _, str := range strs {

// 		if isUniqueLen(str) {
// 			for _, r := range []rune(str) {
// 				m
// 			}
// 		}
// 	}
// 	return m
// }

func replaceSegmentsWithMapped(segs, mapped map[int]string) map[int]string {
	// replace unique length segments with mapped
	for k, v := range mapped {
		segs[k] = sortString(v)
	}

	// replace non unique segments length with mapped {
	// for k,v := range segments {
	// 	if !isUniqueLen()
	// }
	return segs
}

// func getNumbersForStrings(strs []string) []int {
// 	var nbs []int
// 	for _, str := range strs {
// 		nbs = append(nbs, getSignal(str))
// 	}
// 	return nbs
// }

func getOutput(nbs []int) int {
	str := ""
	for _, nb := range nbs {
		s := strconv.Itoa(nb)
		str += s
	}
	out, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("ERROR ON NB: ", str)
		panic(err)
	}
	fmt.Println(out)
	return out
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

// func getSignal(str string) int {
// 	// check for unique number length
// 	if isUniqueLen(str) {
// 		for k, v := range segments {
// 			if len(str) == len(v) {
// 				return k
// 			}
// 		}
// 	}

// 	sorted := sortString(str)

// 	fmt.Println("before: ", str, ", after: ", sorted)
// 	// try to find segment
// 	for k, v := range segments {
// 		if v == sorted {
// 			return k
// 		}
// 	}
// fmt.Println("mapped segment: ",mappedSegments)
// 	// check in mapped segments
// 	for k,v := range mappedSegments {
// 		if v == sorted {
// 			return k
// 		}
// 	}

// 	panic(fmt.Sprintf("could not find a value for str: %v, and sorted: %v\n",str,sorted))
// 	return 0
// }

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
