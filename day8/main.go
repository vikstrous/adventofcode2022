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

type Tree struct {
	Height  int
	Visible bool
}

func ParseTrees(input string) [][]Tree {
	lines := strings.Split(input, "\n")
	trees := [][]Tree{}
	for _, line := range lines {
		treeRow := []Tree{}
		for _, c := range line {
			tree := int(c - rune('0'))
			treeRow = append(treeRow, Tree{
				Height: tree,
			})
		}
		trees = append(trees, treeRow)
	}
	return trees
}

func CountVisible(trees [][]Tree) int {
	vis := 0
	for _, treeRow := range trees {
		for _, tree := range treeRow {
			if tree.Visible {
				vis++
			}
		}
	}
	return vis
}

func PrettyPrint(trees [][]Tree) {
	for _, treeRow := range trees {
		for _, tree := range treeRow {
			if tree.Visible {
				fmt.Printf("T")
			} else {
				fmt.Printf("_")
			}
		}
		fmt.Println()
	}
}

func SetVisible(trees [][]Tree) {
	// Going by row
	for r, treeRow := range trees {
		// From the left
		treeRow[0].Visible = true
		lastVisibleHeight := treeRow[0].Height
		for j, tree := range treeRow {
			if tree.Height > lastVisibleHeight {
				fmt.Println("left visible", tree.Height, lastVisibleHeight, r, j)
				treeRow[j].Visible = true
				lastVisibleHeight = treeRow[j].Height
			}
		}
		// From the right
		treeRow[len(treeRow)-1].Visible = true
		lastVisibleHeight = treeRow[len(treeRow)-1].Height
		for j := len(treeRow) - 1; j >= 0; j-- {
			tree := treeRow[j]
			if tree.Height > lastVisibleHeight {
				fmt.Println("right visible", r, j)
				treeRow[j].Visible = true
				lastVisibleHeight = treeRow[j].Height
			}
		}
	}
	// Going by column
	for col := 0; col < len(trees[0]); col++ {
		// Up to down
		trees[0][col].Visible = true
		lastVisibleHeight := trees[0][col].Height
		for j := 0; j < len(trees); j++ {
			tree := trees[j][col]
			if tree.Height > lastVisibleHeight {
				fmt.Println("up visible", j, col)
				trees[j][col].Visible = true
				lastVisibleHeight = trees[j][col].Height
			}
		}

		// Down to up
		trees[len(trees)-1][col].Visible = true
		fmt.Println("down border visible", len(trees)-1, col)
		lastVisibleHeight = trees[len(trees)-1][col].Height
		for j := len(trees) - 1; j >= 0; j-- {
			tree := trees[j][col]
			if tree.Height > lastVisibleHeight {
				fmt.Println("down visible", j, col)
				trees[j][col].Visible = true
				lastVisibleHeight = trees[j][col].Height
			}
		}
	}
}

func notOutOfBounds(currentX, currentY int, trees [][]Tree) bool {
	return currentX < len(trees[0]) && currentX >= 0 && currentY < len(trees) && currentY >= 0
}

func HighestScenicScore(scores [][]int) int {
	highest := scores[0][0]
	for _, row := range scores {
		for _, score := range row {
			if score > highest {
				highest = score
			}
		}
	}
	return highest
}

func CalcScenicScoreDirection(startX, startY int, trees [][]Tree, xOffset, yOffset int) int {
	if !notOutOfBounds(startX+xOffset, startY+yOffset, trees) {
		return 0
	}
	startHeight := trees[startY][startX].Height
	treeCount := 0
	currentX := startX + xOffset
	currentY := startY + yOffset
	for notOutOfBounds(currentX, currentY, trees) {
		treeHeight := trees[currentY][currentX].Height
		if treeHeight < startHeight {
			treeCount++
		} else {
			treeCount++
			break
		}
		currentX += xOffset
		currentY += yOffset
	}
	return treeCount
}

func CalcAllScenicScores(trees [][]Tree) [][]int {
	scores := [][]int{}
	for y := 0; y < len(trees); y++ {
		scoresRow := []int{}
		for x := 0; x < len(trees[0]); x++ {
			scoresRow = append(scoresRow, CalcScenicScore(CalcScenicScores(x, y, trees)))
		}
		scores = append(scores, scoresRow)
	}
	return scores
}

func CalcScenicScore(a, b, c, d int) int {
	return a * b * c * d
}

func CalcScenicScores(startX, startY int, trees [][]Tree) (int, int, int, int) {
	rightScore := CalcScenicScoreDirection(startX, startY, trees, 1, 0)
	leftScore := CalcScenicScoreDirection(startX, startY, trees, -1, 0)
	downScore := CalcScenicScoreDirection(startX, startY, trees, 0, 1)
	upScore := CalcScenicScoreDirection(startX, startY, trees, 0, -1)
	return upScore, leftScore, downScore, rightScore
}

func Part1(input string) (int, error) {
	t := ParseTrees(input)
	SetVisible(t)
	return CountVisible(t), nil
}

func Part2(input string) (int, error) {
	// 2352 too low
	// 20592 too low
	return HighestScenicScore(CalcAllScenicScores(ParseTrees(input))), nil
}
