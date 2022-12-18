package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	output, err := Part1(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
	output, err = Part2(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

type Assignment struct {
	First, Last int
}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

type Move struct {
	Amount int
	From   int
	To     int
}

func (m Move) apply(stacks [][]rune) {
	for i := 0; i < m.Amount; i++ {
		fromStack := stacks[m.From]
		toStack := stacks[m.To]
		lastCrate := fromStack[len(fromStack)-1]
		fromStack = fromStack[:len(fromStack)-1]
		stacks[m.From] = fromStack
		toStack = append(toStack, lastCrate)
		stacks[m.To] = toStack
	}
}

func (m Move) apply9001(stacks [][]rune) {
	tempStack := []rune{}
	fromStack := stacks[m.From]
	for i := 0; i < m.Amount; i++ {
		lastCrate := fromStack[len(fromStack)-m.Amount+i]
		tempStack = append(tempStack, lastCrate)
	}
	fromStack = fromStack[:len(fromStack)-m.Amount]
	stacks[m.From] = fromStack
	toStack := stacks[m.To]
	toStack = append(toStack, tempStack...)
	stacks[m.To] = toStack
}

func ReadInput(testInput string) ([][]rune, []Move) {
	parts := strings.Split(testInput, "\n\n")
	return ReadInputStacks(parts[0]), ReadInputMoves(parts[1])
}

func ReadInputStacks(testInput string) [][]rune {
	lines := strings.Split(testInput, "\n")
	lastLine := lines[len(lines)-1]
	numStacks := (len(lastLine) + 1) / 4
	stacks := make([][]rune, numStacks)
	for i := len(lines) - 2; i >= 0; i-- {
		for stackNum := 0; stackNum < len(stacks); stackNum++ {
			crate := rune(lines[i][(stackNum*4)+1])
			if crate != ' ' {
				stacks[stackNum] = append(stacks[stackNum], crate)
			}
		}
	}
	return stacks
}

func ReadInputMoves(testInput string) []Move {
	lines := strings.Split(testInput, "\n")
	moves := []Move{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		moves = append(moves, Move{Amount: must(strconv.Atoi(parts[1])), From: must(strconv.Atoi(parts[3])) - 1, To: must(strconv.Atoi(parts[5])) - 1})
	}
	return moves
}

func Part1(testInput string) (string, error) {
	stacks, moves := ReadInput(testInput)
	for _, m := range moves {
		m.apply(stacks)
	}
	topCrates := []rune{}
	for _, stack := range stacks {
		lastCrate := stack[len(stack)-1]
		topCrates = append(topCrates, lastCrate)
	}
	return string(topCrates), nil
}

func Part2(testInput string) (string, error) {
	stacks, moves := ReadInput(testInput)
	for _, m := range moves {
		m.apply9001(stacks)
	}
	topCrates := []rune{}
	for _, stack := range stacks {
		lastCrate := stack[len(stack)-1]
		topCrates = append(topCrates, lastCrate)
	}
	return string(topCrates), nil
}
