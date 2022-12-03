package main

import (
	"fmt"
	"os"
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

var youScores = map[string]int{
	"X": 1, // rock
	"Y": 2, // paper
	"Z": 3, // scissors
}

var gameScores = map[string]int{
	// rock
	"AX": 3,
	"AY": 6,
	"AZ": 0,
	// paper
	"BX": 0,
	"BY": 3,
	"BZ": 6,
	// scissors
	"CX": 6,
	"CY": 0,
	"CZ": 3,
}

func Part1(testInput string) (int, error) {
	score := 0
	for _, line := range strings.Split(testInput, "\n") {
		plays := strings.Split(line, " ")
		oponent := plays[0]
		you := plays[1]
		score += youScores[you]
		score += gameScores[oponent+you]
	}
	return score, nil
}

var winScores = map[string]int{
	"X": 0, // rock
	"Y": 3, // paper
	"Z": 6, // scissors
}

var chosenOutcomeScores = map[string]int{
	// rock
	"AX": 3, // lose -> scisors
	"AY": 1, // tie -> rock
	"AZ": 2, // win -> paper
	// paper
	"BX": 1, // lose -> rock
	"BY": 2, // tie -> paper
	"BZ": 3, // win -> scissors
	// scissors
	"CX": 2, // lose -> paper
	"CY": 3, // tie -> scissors
	"CZ": 1, // win -> rock
}

func Part2(testInput string) (int, error) {
	score := 0
	for _, line := range strings.Split(testInput, "\n") {
		plays := strings.Split(line, " ")
		oponent := plays[0]
		win := plays[1]
		score += winScores[win]
		score += chosenOutcomeScores[oponent+win]
	}
	return score, nil
}
