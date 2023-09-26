package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
)

var days = 256

func main() {
	strs := helpers.GetInputStrings("day6")
	nbs := []int{}
	for _, str := range strs {
		n := helpers.ConvertStringToInts(str, ",")
		nbs = append(nbs, n...)
	}
	simulateWorking(nbs)
}

func simulateWorking(nbs []int) {
	m := make(map[int]int)
	// init
	for i := 0; i <= 8; i++ {
		m[i] = 0
	}
	for _, nb := range nbs {
		m[nb]++
	}

	for day := 0; day < days; day++ {
		newM := initMap()
		for k, v := range m {
			if k == 0 {
				newM[8] += v
				newM[6] += v
			} else {
				newM[k-1] += v
			}
		}
		m = copyMap(newM)
	}

	var count uint64 = 0
	for _, v := range m {
		count += uint64(v)
	}
	fmt.Println("RESP IS: ", count)
}

func initMap() map[int]int {
	newM := make(map[int]int)
	for i := 0; i <= 8; i++ {
		newM[i] = 0
	}
	return newM
}

func copyMap(m map[int]int) map[int]int {
	newM := make(map[int]int)
	for k, v := range m {
		newM[k] = v
	}
	return newM
}
