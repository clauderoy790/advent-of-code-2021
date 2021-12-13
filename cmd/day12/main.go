package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
	"reflect"
	"strings"
	"unicode"
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
		addPath(spl[1], spl[0])
	}
	// part1()
	part2()
}

// func part1() {
// 	fmt.Println("paths: ", paths)
// 	_, _ = getPathRecursive([]string{"start"})
// 	fmt.Println("P1: ", len(alreadyExplored))
// }

func part2() {
	// fmt.Println("paths: ", paths)
	_, _ = getPathRecursive2([]string{"start"})
	fmt.Println("P2: ", len(alreadyExplored))
}

// // Returns true when can't continue and []string will not be nil if path is possible
// func getPathRecursive(currentPath []string) ([]string, bool) {
// 	lastNode := currentPath[len(currentPath)-1]

// 	// when at end
// 	if lastNode == "end" {
// 		fmt.Println("FOUDN PATH : ", currentPath)
// 		alreadyExplored = append(alreadyExplored, currentPath)
// 		return currentPath, true
// 	}

// 	// find all possible paths from this path and call the function on it
// 	possiblePaths := getPossiblePaths(currentPath)
// 	for _, possible := range possiblePaths {
// 		// skip is small cave is already contained in path
// 		if isSmallCave(possible) && containsSmallCave(currentPath, possible) {
// 			fmt.Println("SKIP SMALL CAVE: ", possible, " because contained in: ", currentPath)
// 			continue
// 		}
// 		newP := append(currentPath, possible)
// 		_, _ = getPathRecursive(newP)
// 	}

// 	return nil, true
// }

// Returns true when can't continue and []string will not be nil if path is possible
func getPathRecursive2(currentPath []string) ([]string, bool) {
	lastNode := currentPath[len(currentPath)-1]

	// when at end
	if lastNode == "end" {
		fmt.Println("FOUND PATH : ", currentPath)
		alreadyExplored = append(alreadyExplored, currentPath)
		return currentPath, true
	}

	// find all possible paths from this path and call the function on it
	possiblePaths := getPossiblePaths2(currentPath)
	if len(possiblePaths) == 0 {
		return nil, false
	}
	for _, possible := range possiblePaths {
		newP := append(currentPath, possible)
		_, _ = getPathRecursive2(newP)
	}

	return nil, true
}

func countSmallCaves(currentPath []string) map[string]int {
	small := make(map[string]int)
	for _, path := range currentPath {
		// skip non small caves
		if !isSmallCave(path) || path == "start" {
			continue
		}
		if _, ok := small[path]; !ok {
			small[path] = 0
		}
		small[path]++
	}
	return small
}

// func getPossiblePaths(current []string) []string {
// 	var possible []string
// 	lastNode := current[len(current)-1]

// 	// path going in from current to next path
// 	for _, path := range paths[lastNode] {
// 		if !isSmallCave(path) {
// 			possible = append(possible, path)
// 		} else if !containsSmallCave(current, path) {
// 			possible = append(possible, path)
// 		}
// 	}

// 	return possible
// }

func getPossiblePaths2(current []string) []string {
	var possible []string
	lastNode := current[len(current)-1]

	// path going in from current to next path
	for _, nextStep := range paths[lastNode] {
		if !isSmallCave(nextStep) {
			possible = append(possible, nextStep)
		} else {
			smallCaves := countSmallCaves(current)

			wasVisitedTwice := false
			twiceCave := ""
			for k, v := range smallCaves {
				if v == 2 {
					wasVisitedTwice = true
					twiceCave = k
					break
				}
			}
			// allow certain sc to be visited twice
			if !wasVisitedTwice || (twiceCave != nextStep && smallCaves[nextStep] == 0) {
				possible = append(possible, nextStep)
			}
		}
	}

	return possible
}

func addPath(p1, p2 string) {
	// ignore inversed start/end
	if p1 == "end" || p2 == "start" {
		return
	}

	contains := false
	for _, p := range paths[p1] {
		if p == p1 {
			contains = true
			break
		}
	}

	if !contains {
		paths[p1] = append(paths[p1], p2)
	}
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

// Amit soltion below
func Day12() {
	inp := helpers.GetInputStrings("day12")
	edges := map[string][]string{}

	for _, line := range inp {
		parts := strings.Split(line, "-")

		if parts[0] != "end" && parts[1] != "start" {
			edges[parts[0]] = append(edges[parts[0]], parts[1])
		}
		if parts[1] != "end" && parts[0] != "start" {
			edges[parts[1]] = append(edges[parts[1]], parts[0])
		}
	}
	start := "start"
	end := "end"

	val := CountDFS(edges, [][]string{}, start, end, 1)
	val2 := CountDFS(edges, [][]string{}, start, end, 2)

	fmt.Println(val)
	fmt.Println(val2)
}

func CountDFS(edges map[string][]string, pathStack [][]string, start string, end string, part int) int {

	if start == end {
		return 1
	}

	pathStack = append(pathStack, []string{start})
	partStack := make([]int, len(pathStack))
	partStack[0] = part
	count := 0
	curNode := "start"

	for len(pathStack) > 0 {

		origPart := partStack[len(pathStack)-1]
		partStack = partStack[:len(pathStack)-1]

		curPath := pathStack[len(pathStack)-1]
		pathStack = pathStack[:len(pathStack)-1]

		if len(curPath) > 0 {
			curNode = curPath[len(curPath)-1]
		}

		if curNode == end {
			count++
			continue
		}

		for _, e := range edges[curNode] {
			skip := false
			part = origPart
			for _, v := range curPath {
				if e == v {
					if v == "start" {
						skip = true
						break
					}
					if unicode.IsLower(rune(v[0])) {
						if part == 1 {
							skip = true
							break
						}
						part = 1
						break
					}
				}
			}
			if !skip {
				temp := make([]string, len(curPath))
				copy(temp, curPath)
				temp = append(temp, e)
				pathStack = append(pathStack, temp)
				partStack = append(partStack, part)
			}
		}

	}
	return count
}
