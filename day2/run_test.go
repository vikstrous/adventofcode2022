package main

import (
	"testing"
)

func Test(t *testing.T) {
	total, err := Part1(testInput)
	if err != nil {
		t.Errorf("%s", err)
	}
	if total != 15 {
		t.Errorf("wrong total: %d", total)
	}
}

const testInput = `A Y
B X
C Z`
