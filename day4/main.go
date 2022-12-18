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

// true if a contains b
// aaaaaaaaaaaa
//
//	bbbb
func contains(a, b Assignment) bool {
	return a.First <= b.First && a.Last >= b.Last
}

func Part1(testInput string) (int, error) {
	contained := 0
	for _, line := range strings.Split(testInput, "\n") {
		assigments := []Assignment{}
		assignmentsStr := strings.Split(line, ",")
		for _, assignment := range assignmentsStr {
			startAndEnd := strings.Split(assignment, "-")
			a := Assignment{First: must(strconv.Atoi(startAndEnd[0])), Last: must(strconv.Atoi(startAndEnd[1]))}
			assigments = append(assigments, a)
		}
		if contains(assigments[0], assigments[1]) || contains(assigments[1], assigments[0]) {
			contained += 1
		}
	}
	return contained, nil
}

// true if a overlaps b
// aaaaaaaaaaaa
//
//	bbbb
func overlaps(a, b Assignment) bool {
	return !(a.Last < b.First || a.First > b.Last)
}

func Part2(testInput string) (int, error) {
	contained := 0
	for _, line := range strings.Split(testInput, "\n") {
		assigments := []Assignment{}
		assignmentsStr := strings.Split(line, ",")
		for _, assignment := range assignmentsStr {
			startAndEnd := strings.Split(assignment, "-")
			a := Assignment{First: must(strconv.Atoi(startAndEnd[0])), Last: must(strconv.Atoi(startAndEnd[1]))}
			assigments = append(assigments, a)
		}
		if overlaps(assigments[0], assigments[1]) {
			contained += 1
		}
	}
	return contained, nil
}
