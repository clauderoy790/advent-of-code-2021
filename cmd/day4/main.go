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

func (b *board) markNumber(nb int) bool {
	marked := false
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.nbs[i][j] == nb {
				b.marked[i][j] = true
				marked = true
			}
		}
	}
	return marked
}

func (b *board) bingo() bool {
	return b.hasCompleteColumn() || b.hasCompleteRow()
}

func (b *board) hasCompleteRow() bool {
	for row := 0; row < b.size; row++ {
		completed := true
		for col := 0; col < b.size; col++ {
			if !b.marked[row][col] {
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
	for col := 0; col < b.size; col++ {
		completed := true
		for row := 0; row < b.size; row++ {
			if !b.marked[row][col] {
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
	b.iterateThroughBoard(func(row, col int) {
		if !b.marked[row][col] {
			sum += b.nbs[row][col]
		}
	})
	return sum
}

func (b *board) iterateThroughBoard(it func(int, int)) {
	for row, col := 0, 0; row < b.size; col++ {
		if col == b.size {
			col = 0
			row += 1
		}
		it(row, col)
	}
}

type bingoGame struct {
	nbs          []int
	boards       []board
	winner       *board
	calledNumber int
}

func (g *bingoGame) hasWinner() *board {
	for _, board := range g.boards {
		if board.bingo() {
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
	_, file, _, _ := runtime.Caller(0)
	b, err := os.ReadFile(strings.Replace(file, "main.go", "input.txt", 1))
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(b), "\n")
	size := 5
	bingo := bingoGame{}
	previousIsSpace := false
	currentRow := 0
	var board board
	var twoD [][]int = makeTwoD(size)
	for _, str := range strs {
		str := strings.TrimSpace(str)
		sep := ","
		if len(bingo.nbs) > 0 {
			sep = " "
			twoD = makeTwoD(size)
		}

		if strings.Contains(str, sep) {
			sl := strings.Split(str, sep)
			nbs := make([]int, 0)

			for _, s := range sl {
				nb, _ := strconv.Atoi(s)
				nbs = append(nbs, nb)
			}
			twoD[currentRow] = nbs
			fmt.Println("TWO DDDDD :",twoD)
			currentRow++
		}
		previousIsSpace = false
		if len(str) == 0 {
			previousIsSpace = true
		}
		if previousIsSpace {
			currentRow = 0
			board = newBingoBoard(size)
			bingo.boards = append(bingo.boards, board)
			board.nbs = twoD
			twoD = makeTwoD(size)
		}
	}

	bingo.boards = append(bingo.boards, board)

	return bingo
}

func makeTwoD(size int) [][]int {
	twoD := make([][]int, size)
	for i := 0; i < size; i++ {
		twoD[i] = make([]int, size)
	}
	fmt.Println("THE D : ",twoD)
	return twoD
}
