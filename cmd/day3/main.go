package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	findLifeSupportRating()
	// partOne()
}

func partOne() {
	nbs := readInput()
	lb := findLeftBit(nbs)
	gammaRate := 0
	epsilon := 0
	for i := lb; i >= 0; i-- {
		nbZero, nbOne := countZeroOne(nbs, i)

		if nbOne > nbZero {
			gammaRate += pow(2, i)
		} else if nbZero > nbOne {
			epsilon += pow(2, i)
		}
	}
	fmt.Println("gamma: ", gammaRate)
	fmt.Println("epsilon: ", epsilon)
	fmt.Println("resp is : ", (gammaRate * epsilon))

}

func isBitOne(nb, bit int) bool {
	bin := strconv.FormatInt(int64(nb), 2)
	if bit >= len(bin) {
		return false
	}
	runes := []rune(bin)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return runes[bit] == rune('1')
}

func findLifeSupportRating2() {
	nbs := readInput()
	leftBit := findLeftBit(nbs)
	m := make(map[bool]int)
	m[true] = 0
	m[false] = 0
	for k, _ := range m {
		for i := leftBit; len(nbs) > 1 && i >= 0; i-- {
			zero, one := countZeroOne(nbs, i)
			if k {
				if one >= zero {
					nbs = removeNbsWithBit(nbs, i, false)
				} else {
					nbs = removeNbsWithBit(nbs, i, true)
				}
			} else {
				if zero <= one {
					nbs = removeNbsWithBit(nbs, i, true)
				} else {
					nbs = removeNbsWithBit(nbs, i, false)
				}
			}
		}
		m[k] = nbs[0]
	}
	fmt.Println("fin : ", (m[true] * m[false]))

}
func findLifeSupportRating() {
	nbs := readInput()
	oxy := findOxygenGenerator(nbs)
	co2 := findCo2(nbs)
	fmt.Println("oxy: ", oxy)
	fmt.Println("co2: ", co2)

	fmt.Println("has to be it: ", oxy*co2)
}

func findOxygenGenerator(nbs []int) int {
	leftBit := findLeftBit(nbs)
	n := nbs
	for bit := leftBit; len(n) > 1 && bit >= 0; bit-- {
		zero, one := countZeroOne(n, bit)
		if one >= zero {
			n = removeNbsWithBit(n, bit, false)
		} else {
			n = removeNbsWithBit(n, bit, true)
		}
	}
	if len(n) > 1 {
		panic("oxygen not valid")
	}
	return n[0]
}

func findCo2(nbs []int) int {
	leftBit := findLeftBit(nbs)
	n := nbs
	for bit := leftBit; len(n) > 1 && bit >= 0; bit-- {
		zero, one := countZeroOne(n, bit)
		if zero <= one {
			n = removeNbsWithBit(n, bit, true)
		} else {
			n = removeNbsWithBit(n, bit, false)
		}
	}
	if len(n) > 1 {
		panic("co2 not valid")
	}
	return n[0]
}

func removeNbsWithBit(nbs []int, bit int, removeOne bool) []int {
	n := []int{}
	for _, nb := range nbs {
		if isBitOne(nb, bit) != removeOne {
			n = append(n, nb)
		}

	}
	return n
}

func countZeroOne(nbs []int, bit int) (int, int) {
	zero, one := 0, 0
	for _, nb := range nbs {
		if isBitOne(nb, bit) {
			one++
		} else {
			zero++
		}
	}
	return zero, one
}

func findLeftBit(nbs []int) int {
	left := 0
	for _, nb := range nbs {
		b := findFirstBit(nb)
		if b > left {
			left = b - 1
		}
	}
	return left
}

func findFirstBit(nb int) int {
	firstBit := 0
	for i := 0; firstBit == 0; i++ {
		n := pow(2, i)
		if n > nb {
			firstBit = i
		}
	}
	return firstBit
}

func pow(base, pow int) int {
	return int(math.Pow(float64(base), float64(pow)))
}

func readInput() []int {
	_, file, _, _ := runtime.Caller(0)
	b, err := os.ReadFile(strings.Replace(file, "main.go", "input.txt", 1))
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(b), "\n")
	nbs := make([]int, 0)
	for _, str := range strs {
		str = strings.TrimSpace(str)
		if len(str) == 0 {
			continue
		}
		nb := 0
		if err != nil {
			panic(err)
		}

		for i := len(str) - 1; i >= 0; i-- {
			add := float64(0)
			if string(str[i]) == "1" {
				add = math.Pow(float64(2), float64(math.Abs(float64(len(str)-1-i))))
				nb += int(add)
			}
		}

		nbs = append(nbs, nb)
	}
	return nbs
}
