package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	segments = createSegments()
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

func part2(lines []line) {
	nbs := []int{}
	// determine numbers
	for _, line := range lines {
		// segments := getNumbersForStrings(line.segments)
		mappedSegments = mapSegments(line.segments)
		segments = replaceSegmentsWithMapped(segments,mappedSegments)
		digits := getNumbersForStrings(line.digits)
		output := getOutput(digits)
		fmt.Println("DIGIT:", line.digits, ", value is: ", output)

		nbs = append(nbs, output)
	}

	sum := 0
	for _, nb := range nbs {
		sum += nb
	}
	fmt.Println("P2: ", sum)
}

func mapSegments(strs []string) map[rune]rune {
	m := make(map[rune]rune)
	original
	for _, str := range strs {
		
		if isUniqueLen(str) {
			for _, r := range []rune(str) {
				m
			}
		}
	}
	return m
}

func replaceSegmentsWithMapped(segments,mapped map[int]string) map[int]string {
	// replace unique length segments with mapped
	for k,v := range mapped {
		segments[k] = sortString(v)	
	}

	// replace non unique segments length with mapped {
for k,v := range segments {
	if !isUniqueLen()
}	
	return segments
}

func getNumbersForStrings(strs []string) []int {
	var nbs []int
	for _, str := range strs {
		nbs = append(nbs, getSignal(str))
	}
	return nbs
}

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
func getSignal(str string) int {
	// check for unique number length
	if isUniqueLen(str) {
		for k, v := range segments {
			if len(str) == len(v) {
				return k
			}
		}
	}

	sorted := sortString(str)

	fmt.Println("before: ", str, ", after: ", sorted)
	// try to find segment 
	for k, v := range segments {
		if v == sorted {
			return k
		}
	}
fmt.Println("mapped segment: ",mappedSegments)
	// check in mapped segments
	for k,v := range mappedSegments {
		if v == sorted {
			return k
		}
	}

	panic(fmt.Sprintf("could not find a value for str: %v, and sorted: %v\n",str,sorted))
	return 0
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
func createSegments() map[int]string {
	segments := make(map[int]string)
	segments[0] = "abcefg"
	segments[1] = "cf"
	segments[2] = "acdeg"
	segments[3] = "acdfg"
	segments[4] = "bcdf"
	segments[5] = "abdfg"
	segments[6] = "abdefg"
	segments[7] = "acf"
	segments[8] = "abcdefg"
	segments[9] = "abcdfg"
	return segments
}
