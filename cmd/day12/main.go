package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
	"strings"
)

var paths map[string][]string
var smallCaves map[string]bool
var alreadyCrossed []string

func main() {
	paths = make(map[string][]string)
	smallCaves = make(map[string]bool)
	strs := helpers.GetInputStrings("day12")
	for _, str := range strs {
		spl := strings.Split(str, "-")
		addPath(spl[0], spl[1])
	}
	part1()
}

var fullPaths []string

func part1() {
	fullPaths = make([]string, 0)

	start := getPossiblePaths("start")
	for _, s := range start {
		createPathsFrom([]string{s})
	}

	fmt.Println("P1: ", len(fullPaths))
}

func createPathsFrom(initials []string) {
	paths := getPossiblePaths(initials[len(initials)-1])
	//create additional paths
	for _, path := range paths {
		if path == "end" {
	_, ok := smallCaves[str]
	return ok
}

func visitedSmallCave(str string) bool {
	return smallCaves[str]
}

func visitSmallCave(str string) {
	if !isSmallCave(str) {
		panic("cannot visit this cave as small cave: " + str)
	}

	if smallCaves[str] {
		panic("already visited this cave for path " + str)
	}

	smallCaves[str] = true
}

func addSmallCave(str string) {
	if !isSmallCave(str) {
		panic("cannot add this cave as big cave: " + str)
	}

	if containsSmallCave(str) {
		panic("cannot add already present: " + str)
	}

	smallCaves[str] = false
}

func resetSmallCaves() {
	for k, _ := range smallCaves {
		smallCaves[k] = false
	}
}

func isSmallCave(str string) bool {
	return strings.ToUpper(str) != str
}
