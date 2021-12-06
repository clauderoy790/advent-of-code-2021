package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	game := readInput()
	for {
		if game.playTurn() {
			break
		}
	}
	fmt.Println("result! ", game.calculateResult())
}

func newBingoBoard(size int) board {
	b := board{
		size: size,
	}
	marked := make([][]bool, size)
	market2 := make([][]bool, size)
	for i := 0; i < size; i++ {
		marked[i] = make([]bool, size)
		market2[i] = make([]bool, size)
	}
	b.marked = marked
	b.markedTemp = market2

	return b
}

type board struct {
	nbs        [][]int
	marked     [][]bool
	markedTemp [][]bool
	size       int
}

func (b *board) markNumber(nb int) {
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.nbs[i][j] == nb {
				b.markedTemp[i][j] = true
			}
		}
	}
}

func (b *board) saveMarkTemp() {
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			b.marked[i][j] = b.markedTemp[i][j]
		}
	}
}

func (b *board) bingo() bool {
	return b.hasCompleteColumn() || b.hasCompleteRow()
}

func (b *board) hasCompleteRow() bool {
	for i := 0; i < b.size; i++ {
		completed := true
		for _, m := range b.markedTemp[i] {
			if !m {
				completed = false
				break
			}
		}
		if completed {
			return true
		}
	}
	return false
}

func (b *board) hasCompleteColumn() bool {
	for i := 0; i < b.size; i++ {
		completed := true
		for j := 0; j < b.size; j++ {
			if !b.markedTemp[j][i] {
				completed = false
				break
			}
		}
		if completed {
			return true
		}
	}
	return false
}

func (b *board) unmarkedSum() int {
	sum := 0
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if !b.marked[i][j] {
				sum += b.nbs[i][j]
			}
		}
	}
	return sum
}

func (b *board) String() string{
	return fmt.Sprintf("NBS: %v \nMarked: %v\nMarkedTemp: %v\n",b.nbs,b.marked, b.markedTemp)
}

type bingoGame struct {
	nbs            []int
	boards         []board
	winner         *board
	previousWinner *board
	calledNumber   int
}

// returns if game is over
func (g *bingoGame) playTurn() bool {
	called := 0
	if len(g.nbs) > 0 {
		called = g.nbs[0]
		g.nbs = g.nbs[1:]
	} else {
		fmt.Println("DONE")
		return true
	}
	fmt.Println("MARK : ", called)
	for _, b := range g.boards {
		b.markNumber(called)
	}
	bi := g.checkForBingos(called)
	if bi {
		return bi
	}
	
	return len(g.boards) == 0
}

func (g *bingoGame) checkForBingos(called int) bool {
	boards := []board{}
	for _, b := range g.boards {
		if b.bingo() {
			b.saveMarkTemp()
			fmt.Println("BINGO: ",called)
			fmt.Println(&b)
			g.previousWinner = &b
			g.winner = &b
			g.calledNumber = called

			boards = append(boards, b)
			if len(g.boards) == 1 {
				fmt.Println("WINNER IS: ", g.winner)
				return true
			}
		}
	}
	for _, b := range boards {
		in := -1
		for i, bo := range g.boards {
			if reflect.DeepEqual(b, bo) {
				in = i
				break
			}
		}
		if in >= 0 {
			g.boards = append(g.boards[:in], g.boards[in+1:]...)
		}
	}
	return false
}

func (g *bingoGame) calculateResult() int {
	if g.winner == nil {
		return 0
	}
	fmt.Println("WINNER: ", g.winner)
	unmarked := g.winner.unmarkedSum()
	fmt.Println("KED : ", unmarked)
	fmt.Println("called: ", g.calledNumber)
	return unmarked * g.calledNumber
}

func readInput() bingoGame {
	strs := getInputStrings()
	size := 5
	bingo := bingoGame{}
	bingo.nbs = convertStringToInts(strs[0], ",")
	currentRow := 0
	var board board = newBingoBoard(size)
	var twoD [][]int = makeTwoD(size)
	for _, str := range strs[1:] {
		if len(str) == 0 {
			continue
		}

		twoD[currentRow] = convertStringToInts(str, " ")
		currentRow++
		if currentRow == size {
			currentRow = 0
			board = newBingoBoard(size)
			board.nbs = twoD
			bingo.boards = append(bingo.boards, board)
			twoD = makeTwoD(size)
		}
	}

	return bingo
}

func makeTwoD(size int) [][]int {
	twoD := make([][]int, size)
	for i := 0; i < size; i++ {
		twoD[i] = make([]int, size)
	}
	return twoD
}

func convertStringToInts(str string, sep string) []int {
	str = strings.TrimSpace(str)
	sl := strings.Split(str, sep)
	nbs := make([]int, 0)

	for _, s := range sl {
		nb, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			continue
		}
		nbs = append(nbs, nb)
	}
	if len(nbs) > 5 && !strings.Contains(str, ",") {
		panic(nbs)
	}
	return nbs
}

func getInputStrings() []string {
	_, file, _, _ := runtime.Caller(0)
	b, err := os.ReadFile(strings.Replace(file, "main.go", "input.txt", 1))
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}
