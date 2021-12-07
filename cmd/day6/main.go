package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
)

var days = 18

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
	var sum uint64 = 0
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
	fmt.Println("SIMULATE: ", nb, ", days: ", days)

	count := countAddedFrom(0, nb)
	c <- uint64(count)
}

func countAddedFrom(startDay, startValue int) int {
	init := startValue
	nbInitDay := 0
	check := 2
	dayToCheck := 7
	if startDay == dayToCheck {
		fmt.Println("FUCKKKKK")
	}
	if startValue == check {
		fmt.Println("here")
	}
	count := 0

	// trim to even value
	for startValue != 6 {
		startValue--
		nbInitDay++
		if startValue == 0 || startValue == 7 {
			startValue = 6
			count++
			break
		}
		startDay++

		// At the end of the days, return 0
		if startDay >= days {
			return 0
		}
	}

	// calculate day diff
	d := days - startDay

	// total added days
	addedCount := d / 7
	if init == check {

		fmt.Println("d: ", d)
	}
	// if d%7 != 0 {
	// 	addedCount++
	// }

	// recursively count the number of days that will be added
	count += addedCount
	for i := 0; i < addedCount; i++ {
		newStartDay := init + startDay + (i * 7)
		count += countAddedFrom(newStartDay, 8)
		if init == check {
			fmt.Println("new star tday: ", newStartDay)
		}
	}

	if init == check {
		fmt.Println("added count: ", count)
	}

	return count
}
