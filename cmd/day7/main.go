package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
	"math"
)

func main() {
	strs := helpers.GetInputStrings("day7")
	nbs := []int{}
	for _, str := range strs {
		nbs = append(nbs, helpers.ConvertStringToInts(str, ",")...)
	}
	smallest := findSmallest(nbs)
	largest := findLargest(nbs)
	m := make(map[int]int)
	for i := smallest; i <= largest; i++ {
		n := countFuelP2(nbs, i)
		m[i] = n
	}

	pos := 0
	cheapest := math.MaxInt
	for k, v := range m {
		if v < cheapest {
			pos = k
			cheapest = v
		}
	}
	fmt.Println("cheapest is in row : ", pos, ", with fuel: ", cheapest)
}

func countFuelP2(nbs []int, pos int) int {
	fuel := 0
	for _, nb := range nbs {
		cost := 1
		start, end := nb, pos
		if end < start {
			start, end = end, start
		}
		for i := start; i < end; i++ {
			fuel += cost
			cost++
		}
	}
	return fuel
}

// func countFuel(nbs []int, pos int) int {
// 	fuel := 0
// 	for _, nb := range nbs {
// 		f := int(math.Abs(float64(nb - pos)))
// 		fuel += f
// 	}
// 	return fuel
// }

func findSmallest(nbs []int) int {
	small := nbs[0]
	for _, nb := range nbs {
		if nb < small {
			small = nb
		}
	}
	return small
}
func findLargest(nbs []int) int {
	largest := nbs[0]
	for _, nb := range nbs {
		if nb > largest {
			largest = nb
		}
	}
	return largest
}
