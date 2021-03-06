package main

import (
	"fmt"
	"os"
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

func newBoardGroup() boardGroup {
	return boardGroup{}
}

type boardGroup struct {
	boards []board
}

func (b *boardGroup) markNumber(nb int) {
	for _, board := range b.boards {
		board.markNumber(nb)
	}
}
func newBingoBoard(size int) board {
	b := board{
		size: size,
	}
	marked := make([][]bool, size)
	for i := 0; i < size; i++ {
		marked[i] = make([]bool, size)
	}
	b.marked = marked

	return b
}

type board struct {
	nbs    [][]int
	marked [][]bool
	size   int
}

func (b *board) markNumber(nb int)  {
	for i := 0; i < b.size;i++ {
		for j:=0;j<b.size;j++ {
			if b.nbs[i][j] == nb {
				b.marked[i][j] = true
			}
		}
	}
}

func (b *board) hasCompleteRow() bool {
	for i := 0; i < b.size; i++ {
		completed := true
		for _, m  := range b.marked[i] {
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
			if !b.marked[j][i] {
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

type bingoGame struct {
	nbs          []int
	boards       []board
	winner       *board
	calledNumber int
}

func (g *bingoGame) hasWinner() *board {
	for _, board := range g.boards {
		if board.hasCompleteColumn() || board.hasCompleteRow() {
			return &board
		}
	}
	return nil
}

// returns if game is over
func (g *bingoGame) playTurn() bool {
	if len(g.nbs) > 0 {
		g.calledNumber = g.nbs[0]
		g.nbs = g.nbs[1:]
	} else {
		return true
	}
	for _, b := range g.boards {
		b.markNumber(g.calledNumber)
	}
	if win := g.hasWinner(); win != nil {
		g.winner = win
		return true
	}
	return false
}

func (g *bingoGame) calculateResult() int {
	if g.winner == nil {
		return 0
	}
	unmarked := g.winner.unmarkedSum()
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
