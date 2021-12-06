package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
)

var NONE = -100
var days = 256

func main() {
	strs := helpers.GetInputStrings("day6")
	// i := 0
	nbs := []int{}
	for _, str := range strs {
		n := helpers.ConvertStringToInts(str, ",")
		nbs = append(nbs, n...)
	}
	simulateAsync(nbs)
}

func simulateAsync(nbs []int) {
	ch := make(chan uint64, len(nbs))
	var sum uint64 = 0
	fmt.Println("LEN NS: ", len(nbs))
	for _, nb := range nbs {
		go simulateOne(nb, ch)
	}

	for i := 0; i < len(nbs); i++ {
		nb := <-ch
		fmt.Println("GET VALUE", i, ": ", nb)
		sum += nb
	}

	fmt.Println("sum: ", sum)
}

func simulateOne(nb int, c chan uint64) {
	// fmt.Println("SIMULATE: ",nb, ", days: ",days)

	nbs := [500000000]byte{}
	nbs[0] = byte(nb)
	count := 1

	for day := 0; day < days; day++ {
		// fmt.Println("day: ",day)
		toAdd := 0

		// Loop through numbers
		for i := 0; i < count; i++ {
			nbs[i]--
			if nbs[i] == 255 {
				nbs[i] = 6
				toAdd++
			}
		}

		// Add new numbers
		for i := 0; i < toAdd; i++ {
			// fmt.Println("ADD NEW NB: ",toAdd)
			nbs[i+count] = 8
		}
		count+=toAdd
	}
	c <- uint64(count)
}
