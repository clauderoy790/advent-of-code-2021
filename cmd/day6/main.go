package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
)

var days = 256

func main() {
	// get the input and transform to a int array
	strs := helpers.GetInputStrings("day6")
	nbs := []int{}
	for _, str := range strs {
		n := helpers.ConvertStringToInts(str, ",")
		nbs = append(nbs, n...)
	}

	simulate(nbs)
}

func simulate(initialFishes []int) {
	m := initMap() // initialize a map with keys 0 to 8

	// initialize the map with the input fishes
	for _, nb := range initialFishes {
		m[nb]++
	}

	// run the simulation for X days
	for day := 0; day < days; day++ {
		newM := initMap() // initialize a map with keys 0 to 8
		for day, fishCount := range m {
			if day == 0 { // if we get to day 0, the current fish resets at 6 days and creates a new fish at 8 days
				newM[8] += fishCount
				newM[6] += fishCount
			} else { // decrease our day value by one with the same fish count
				newM[day-1] += fishCount
			}
		}
		m = newM // override previous day map
	}

	var totalFishCount uint64 = 0
	for _, nbFish := range m {
		totalFishCount += uint64(nbFish)
	}
	fmt.Println("the number of fish is: ", totalFishCount)
}

func initMap() map[int]int {
	newM := make(map[int]int)
	for i := 0; i <= 8; i++ {
		newM[i] = 0
	}
	return newM
}
