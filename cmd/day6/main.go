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
	simulateAsync(nbs)
}

func simulateAsync(nbs []int) {
	ch := make(chan uint64, len(nbs))
	var sum uint64 = uint64(len(nbs))
	for _, nb := range nbs {
		go simulateOne(nb, ch)
	}

	for i := 0; i < len(nbs); i++ {
		nb := <-ch
		sum += nb
	}

	fmt.Println("sum: ", sum)
}

func simulateOne(nb int, c chan uint64) {
	count := countAddedFrom(0, nb)
	c <- uint64(count)
}

func countAddedFrom(startDay, startValue int) int {
	// return if no new number
	if startDay+startValue >= days {
		return 0
	}

	// trim to even value
	count := 1
	startDay += startValue

	// recursively call other numbers
	additional := (days - startDay -1) / 7
	count += additional
	for i := 0; i <= additional;i++ {
		start := startDay+(i*7)+1
		count += countAddedFrom(start,8)
	}

	return count
}
