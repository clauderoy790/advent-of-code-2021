package main

import (
	"clauderoy790/advent-of-code-2021/helpers"
	"fmt"
	"sync"
)

var days = 18

func main() {
	strs := helpers.GetInputStrings("day6")
	nbs := []int{}
	for _, str := range strs {
		n := helpers.ConvertStringToInts(str, ",")
		nbs = append(nbs, n...)
	}
	// simulateAsync(nbs)
	simulateMap(nbs)
}

var count uint64

func simulateMap(nbs []int) {
	wg := new(sync.WaitGroup)
	count = uint64(len(nbs))
	for _, nb := range nbs {
		go simulateM(0, nb, wg)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println("ans: ", count)
}

func simulateM(startDay int, startVal int, wg *sync.WaitGroup) {
	defer wg.Done()
	day := startDay
	if day >= days {
		return
	}
	if day == 5 && startVal == 8 {
		fmt.Println("debug")
	}
	for startVal > 0 {
		if day >= days {
			return
		}
		startVal--
		day++
	}
	
	var ct uint64 = 0
	additions := (days - day) / 7
	for i := 0; i <= additions; i++ {
		startDay := day + 1 + (i * 7)
		go simulateM(startDay, 8, wg)
		wg.Add(1)
		ct++
	}
	fmt.Println("CT: ",ct)
	inc(ct)
}

var mu sync.Mutex
func inc(val uint64) {
	mu.Lock()
	defer mu.Unlock()
	count += val
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
	additional := (days - startDay - 1) / 7
	count += additional
	for i := 0; i <= additional; i++ {
		start := startDay + (i * 7) + 1
		count += countAddedFrom(start, 8)
	}

	return count
}
