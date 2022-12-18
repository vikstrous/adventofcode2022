package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	stacks, moves := ReadInput(testInput)
	for i, stack := range stacks {
		fmt.Printf("stack %d ", i)
		for _, c := range stack {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
	fmt.Printf("%#v", moves)

	result, err := Part1(testInput)
	if err != nil {
		t.Error(err)
	}
	if result != "CMZ" {
		t.Error("wrong result", result)
	}
	result, err = Part2(testInput)
	if err != nil {
		t.Error(err)
	}
	if result != "MCD" {
		t.Error("wrong result", result)
	}
}

const testInput = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
