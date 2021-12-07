package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	nbs := readInput()
	timeIncrease := increasedFromPreviousCount(nbs)
	secondCount := increasedFromPreviousCount2(nbs)
	fmt.Println(timeIncrease)
	fmt.Println("second: ", secondCount)
}

func increasedFromPreviousCount(nbs []int) int {
	previous, increase := -1, 0
	for _, nb := range nbs {
		if previous != -1 && nb > previous {
			increase++
		}
		previous = nb
	}
	return increase
}

func increasedFromPreviousCount2(nbs []int) int {
	increase, previous, current := 0, -1, 0
	window := window{size: 4, measurements: 3}
	for i := 0; i < len(nbs); i++ {
		if len(nbs) < i+window.measurements {
			break
		}
		for j := 0; j < window.measurements; j++ {
			current += nbs[i+j]
		}
		if previous > 0 && current > previous {
			increase++
		}

		previous = current
		current = 0
	}
	return increase
}

type window struct {
	size         int
	measurements int
}

func readInput() []int {
	_, file, _, _ := runtime.Caller(0)
	b, err := os.ReadFile(strings.Replace(file, "main.go", "input.txt", 1))
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(b), "\n")
	nbs := []int{}
	for _, str := range strs {
		nb, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			panic(err)
		}
		nbs = append(nbs, nb)
	}
	return nbs
}
