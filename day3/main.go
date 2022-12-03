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

func Priority(r rune) int {
	if r >= 97 && r <= 122 {
		return int(r-97) + 1
	}
	return int(r-65) + 27
}

func PriorityToLetter(p int) rune {
	if p < 27 {
		return rune(p + 97 - 1)
	}
	return rune(p + 65 - 27)
}

func findBadge(threeLines []string) int {
	a := countsByPriorityForLine(threeLines[0])
	b := countsByPriorityForLine(threeLines[1])
	c := countsByPriorityForLine(threeLines[2])
	for i := range a {
		if a[i] > 0 && b[i] > 0 && c[i] > 0 {
			return i
		}
	}
	panic("no badge found")
}

func countsByPriorityForLine(line string) []int {
	countsByPriority := make([]int, 27*2)
	for _, r := range line {
		priority := Priority(r)
		countsByPriority[priority] += 1
	}
	return countsByPriority
}

func findDuplicateItemPriority(line string) int {
	countsByPriority := make([]int, 27*2)
	for i, r := range line {
		priority := Priority(r)
		// first half
		if i < len(line)/2 {
			if countsByPriority[priority] > 0 {
				return priority
			}
			countsByPriority[priority] -= 1
		} else {
			if countsByPriority[priority] < 0 {
				return priority
			}
			countsByPriority[priority] += 1
		}
	}
	panic("no duplicate")
}

func Part1(testInput string) (int, error) {
	sum := 0
	for _, line := range strings.Split(testInput, "\n") {
		priority := findDuplicateItemPriority(line)
		fmt.Println(priority, string(PriorityToLetter(priority)))
		sum += priority
	}
	return sum, nil
}

func Part2(testInput string) (int, error) {
	sum := 0
	threeLines := []string{}
	for _, line := range strings.Split(testInput, "\n") {
		threeLines = append(threeLines, line)
		if len(threeLines) < 3 {
			continue
		}
		priority := findBadge(threeLines)
		fmt.Println(priority, string(PriorityToLetter(priority)))
		sum += priority
		threeLines = []string{}
	}
	return sum, nil
}
