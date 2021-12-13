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
		addPath(spl[1], spl[0])
	}
	fmt.Println("PATHS: ", paths)
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
		fmt.Println("EXIT on node: ", lastNode)
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
	fmt.Println("GET POSS PATH: ", current)
	var possible []string
	lastNode := current[len(current)-1]
	if len(current) == 3 && lastNode == "c" {
		fmt.Println("db")
	}

	// path going in from current to next path
	for _, nextStep := range paths[lastNode] {
		if !isSmallCave(nextStep) {
			possible = append(possible, nextStep)
		} else {
			smallCaves := countSmallCaves(current)

			// wtf
			for _, v := range smallCaves {
				if v > 2 {
					fmt.Println("WTF")
				}
			}
			wasVisitedTwice := false
			twiceCave := ""
			for k, v := range smallCaves {
				if v == 2 {
					wasVisitedTwice = true
					twiceCave = k
					break
				}
			}

			// if wasVisitedTwice {
			// 	fmt.Println("TWICE")
			// }

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
