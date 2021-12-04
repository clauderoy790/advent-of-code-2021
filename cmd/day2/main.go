package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	sub := submarine{}
	commands := readInput()

	for _, cmd := range commands {
		switch cmd.direction {
		case "forward":
			sub.forward(cmd.value)
		case "up":
			sub.up(cmd.value)
		case "down":
			sub.down(cmd.value)
		}
	}
	fmt.Println("final", (sub.x * sub.depth))
}

type submarine struct{ x, aim, depth int }

func (s *submarine) forward(x int) {
	s.x += x
	s.depth += s.aim * x
}
func (s *submarine) down(y int) {
	s.aim += y
	//s.depth += y
}
func (s *submarine) up(y int) {
	s.aim -= y
	//s.depth -= y
}

func readInput() []command {
	_, file, _, _ := runtime.Caller(0)
	b, err := os.ReadFile(strings.Replace(file, "main.go", "input.txt", 1))
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(b), "\n")
	commands := []command{}
	for _, str := range strs {
		if err != nil {
			panic(err)
		}
		cmd := strings.Split(str, " ")

		val, err := strconv.Atoi(cmd[1])
		if err != nil {
			panic(err)
		}

		commands = append(commands, command{direction: cmd[0], value: val})
	}
	return commands
}

type command struct {
	direction string
	value     int
}
