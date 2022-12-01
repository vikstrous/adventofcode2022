package main

import (
	"testing"
)

func Test(t *testing.T) {
	total, err := Part1(testInput)
	if err != nil {
		t.Errorf("%s", err)
	}
	if total[0] != 24000 {
		t.Errorf("wrong total: %d", total)
	}
}

const testInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
