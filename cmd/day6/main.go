package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
	"math"
)

var mi, mj = 0,0
var arrays [int(math.MaxInt)][int(math.MaxInt)]int

func main() {
	strs := helpers.GetInputStrings("day6")

	nbs := []int{}
	for i, str := range strs {
		nbs := helpers.ConvertStringToInts(str, ",")
		for _, nb := range nbs {
			arrays[mi][mj] = nb
			mj++
		}
	}
	// fmt.Println("After 0 days:", nbs)
	days := 256
	for i := 0; i < days; i++ {
		arrays = simulate(arrays)
		fmt.Println("DAY ",i)
		// fmt.Printf("After %v days: %v\n", i, nbs)
	}
	var total uint64 = 0

	for i := range arrays {
		for j := range arrays[i] {
			if arrays[i][j] != -100 {
				total++
			}
		}
	}

	fmt.Println("RESULT IS : ",total)
}

func simulate(slices [][]int) [][]int {

	for i := 0; i < mi;i++ {
		// continue here
		for j := 0;j< {
			if slices[i][j] == -100 {
				continue
			}
			slices[i][j]--
			if slices[i][j] < 0 {
				slices[i][j] = 6
					
				toAdd = append(toAdd, 8)
			}
		}
	}

	ind := len(slices) - 1
	for _, add := range toAdd {
		if len(slices[ind]) == math.MaxInt-1 {
			slices = append(slices, make([]int, 0))
			ind++
		}
		slices[ind] = append(slices[ind], add)
	}

	return slices
}

func addNumber(nb int) {
	if mj == int(math.MaxInt-1) {
		mj = 0
		mi++
	}
	arrays[mi][mj] = nb
	mi++
}

func getNewArray() []int {
	var arr [int(math.MaxInt)]

	for i := range arr {
		arr[i] = -100
	}
	return arr
}
