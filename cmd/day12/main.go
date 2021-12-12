package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
	"reflect"
	"strings"
)

var paths map[string][]string
var alreadyExplored [][]string

func main() {
	alreadyExplored = make([][]string, 0)
	paths = make(map[string][]string)
	strs := helpers.GetInputStrings("day12")
	for _, str := range strs {
		spl := strings.Split(str, "-")
		addPath(spl[0], spl[1])
	}
	part1()
}

func part1() {
	fmt.Println("paths: ", paths)
	_, _ = getPathRecursive([]string{"start"})
	fmt.Println("P1: ", len(alreadyExplored))
}

// Returns true when can't continue and []string will not be nil if path is possible
func getPathRecursive(currentPath []string) ([]string, bool) {
	lastNode := currentPath[len(currentPath)-1]

	// when at end
	if lastNode == "end" {
		fmt.Println("FOUDN PATH : ", alreadyExplored)
		alreadyExplored = append(alreadyExplored, currentPath)
		return currentPath, true
	}

	// find all possible paths from this path and call the function on it
	possiblePaths := paths[lastNode]
	for _, possible := range possiblePaths {
		// skip is small cave is already contained in path
		if isSmallCave(possible) && containsSmallCave(currentPath, possible) {
			continue
		}
		newP := append(currentPath, possible)
		_, _ = getPathRecursive(newP)
	}

	return nil, true
}

func addPath(p1, p2 string) {
	if _, ok := paths[p1]; !ok {
		paths[p1] = make([]string, 0)
	}
	paths[p1] = append(paths[p1], p2)
}

func alreadyContainsPath(path []string) bool {
	for _, explored := range alreadyExplored {
		if reflect.DeepEqual(path, explored) {
			return true
		}
	}
	return false
}

func containsSmallCave(path []string, smallCave string) bool {
	for _, p := range path {
		if p == smallCave {
			return true
		}
	}
	return false
}

func isSmallCave(str string) bool {
	return strings.ToUpper(str) != str
}
