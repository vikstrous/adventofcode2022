package main

import (
	"testing"
)

func Test(t *testing.T) {
	result, err := Part1(testInput)
	if err != nil {
		t.Error(err)
	}
	if result != 2 {
		t.Error("wrong result", result)
	}
}

const testInput = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
