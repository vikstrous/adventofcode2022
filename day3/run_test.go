package main

import (
	"fmt"
	"testing"
)

func TestPriority(t *testing.T) {
	if Priority('a') != 1 {
		t.Errorf("wrong priority for a: %d", Priority('a'))
	}
	if Priority('z') != 26 {
		t.Errorf("wrong priority for z: %d", Priority('z'))
	}
	if Priority('A') != 27 {
		t.Errorf("wrong priority for A: %d", Priority('A'))
	}
	if Priority('Z') != 52 {
		t.Errorf("wrong priority for Z: %d", Priority('Z'))
	}
	for i := 1; i < 27*2; i++ {
		if i != Priority(PriorityToLetter(i)) {
			t.Errorf("wrong result for %d: %s", i, string(PriorityToLetter(i)))
		}
	}
}

func Test(t *testing.T) {
	total, err := Part1(testInput)
	if err != nil {
		t.Errorf("%s", err)
	}
	if total != 157 {
		t.Errorf("wrong total: %d", total)
	}
	fmt.Println("next")
	total, err = Part2(testInput)
	if err != nil {
		t.Errorf("%s", err)
	}
	if total != 70 {
		t.Errorf("wrong total: %d", total)
	}
}

var testInput = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
