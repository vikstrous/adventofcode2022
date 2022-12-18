package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	for i, tc := range testInputs {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result, err := Part1(tc.Input)
			if err != nil {
				t.Error(err)
			}
			if result != tc.Output {
				t.Error("wrong result", result)
			}
		})
	}
}

type TC struct {
	Input  string
	Output int
}

var testInputs = []TC{
	{Input: `bvwbjplbgvbhsrlpgdmjqwftvncz`, Output: 5},
	{Input: `nppdvjthqldpwncqszvftbrmjlhg`, Output: 6},
	{Input: `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`, Output: 10},
	{Input: `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`, Output: 11},
}
