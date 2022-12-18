package main

import (
	"errors"
	"fmt"
	"os"
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

func allDifferent(seq []rune) bool {
	seen := map[rune]struct{}{}
	for _, s := range seq {
		_, ok := seen[s]
		if ok {
			return false
		}
		seen[s] = struct{}{}
	}
	return true
}

func Part1(testInput string) (int, error) {
	buf := []rune{}
	for i, c := range testInput {
		buf = append(buf, c)
		if len(buf) > 4 {
			buf = buf[1:]
		}
		if len(buf) == 4 {
			if allDifferent(buf) {
				fmt.Println(buf)
				return i + 1, nil
			}
		}
	}
	return 0, errors.New("no signal")
}

func Part2(testInput string) (int, error) {
	buf := []rune{}
	for i, c := range testInput {
		buf = append(buf, c)
		if len(buf) > 14 {
			buf = buf[1:]
		}
		if len(buf) == 14 {
			if allDifferent(buf) {
				fmt.Println(buf)
				return i + 1, nil
			}
		}
	}
	return 0, errors.New("no signal")
}
