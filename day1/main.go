package main

import (
	"fmt"
	"os"
	"sort"
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
	fmt.Println(Part2(output))
}

func Part2(nums []int64) int64 {
	total := int64(0)
	for _, num := range nums {
		total += num
	}
	return total
}

func Part1(testInput string) ([]int64, error) {
	elfTotals := []int64{}
	currentTotal := int64(0)
	for _, line := range strings.Split(testInput, "\n") {
		if line == "" {
			elfTotals = append(elfTotals, currentTotal)
			currentTotal = 0
			continue
		}
		cal, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("bad input %s: %w", line, err)
		}
		currentTotal += cal
	}
	sort.Slice(elfTotals, func(i, j int) bool { return elfTotals[i] > elfTotals[j] })
	return elfTotals[:3], nil
}
