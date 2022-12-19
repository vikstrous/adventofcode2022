package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	trees := ParseTrees(testInput)
	fmt.Println(trees)
	// SetVisible(trees)
	// fmt.Println(trees)
	// PrettyPrint(trees)
	// fmt.Println(CountVisible(trees))
	// fmt.Println(CalcScenicScores(3, 2, trees))
	// fmt.Println(CalcScenicScoreDirection(3, 2, trees, 1, 0))
	// fmt.Println(CalcScenicScore(CalcScenicScores(2, 1, trees)))
	fmt.Println(CalcAllScenicScores(trees))
	// fmt.Println(HighestScenicScore(CalcAllScenicScores(trees)))
}

const testInput = `30373
25512
65332
33549
35390`
