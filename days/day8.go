package days

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func fillArray(s string) []int {
	split := strings.Split(s, "")
	column := []int{}
	for _, t := range split {
		height, err := strconv.Atoi(t)
		if err != nil {
			panic("something went wrong")
		}
		column = append(column, height)
	}
	return column
}

func getVisibleTrees(treeMap [][]int) int {
	checkEdgeVisibility := func(treeMap [][]int, x, y int) bool {
		if x == 0 || x == len(treeMap)-1 || y == 0 || y == len(treeMap[x])-1 {
			return true
		}
		return false
	}

	checkInteriorVisibility := func(treeMap [][]int, x, y, h int) bool {
		visibleSides := 4
		var i, j int
		for i = x; i >= 0; i-- {
			if i != x && treeMap[i][y] >= h {
				visibleSides -= 1
				break
			}
		}
		for i = x; i <= len(treeMap)-1; i++ {
			if i != x && treeMap[i][y] >= h {
				visibleSides -= 1
				break
			}
		}

		for j = y; j >= 0; j-- {
			if j != y && treeMap[x][j] >= h {
				visibleSides -= 1
				break
			}
		}
		for j = y; j <= len(treeMap[x])-1; j++ {
			if j != y && treeMap[x][j] >= h {
				visibleSides -= 1
				break
			}
		}
		return visibleSides > 0
	}

	var visibleTrees int
	var x, y int
	for x = 0; x < len(treeMap); x++ {
		for y = 0; y < len(treeMap[x]); y++ {
			if checkEdgeVisibility(treeMap, x, y) || checkInteriorVisibility(treeMap, x, y, treeMap[x][y]) {
				visibleTrees += 1
			}
		}
	}

	return visibleTrees
}

func getVisibilityScore(treeMap [][]int) int {

	checkVisibilityScore := func(treeMap [][]int, x, y, h int) int {
		var left, right, down, up, total int
		var i, j int
		for i = x; i >= 0; i-- {
			if i == x {
				continue
			}
			left += 1
			if treeMap[i][y] >= h {
				break
			}
		}
		for i = x; i <= len(treeMap)-1; i++ {
			if i == x {
				continue
			}
			right += 1
			if treeMap[i][y] >= h {
				break
			}
		}
		for j = y; j >= 0; j-- {
			if j == y {
				continue
			}
			down += 1
			if treeMap[x][j] >= h {
				break
			}
		}
		for j = y; j <= len(treeMap[x])-1; j++ {
			if j == y {
				continue
			}
			up += 1
			if treeMap[x][j] >= h {
				break
			}
		}
		total = left * right * down * up
		return total
	}

	maxScore := 0
	var x, y int
	for x = 0; x < len(treeMap); x++ {
		for y = 0; y < len(treeMap[x]); y++ {
			itemScore := checkVisibilityScore(treeMap, x, y, treeMap[x][y])
			if itemScore > maxScore {
				maxScore = itemScore
			}
		}
	}

	return maxScore
}

func (d Days) Day8p1(source io.Reader) int {
	var trees [][]int

	lines := bufio.NewScanner(source)
	for lines.Scan() {
		if lines.Text() != "" {
			trees = append(trees, fillArray(lines.Text()))
		}
	}

	result := getVisibleTrees(trees)
	fmt.Printf("(Part 1) Visible trees: %d\n", result)
	return result
}

func (d Days) Day8p2(source io.Reader) int {
	var trees [][]int

	lines := bufio.NewScanner(source)
	for lines.Scan() {
		if lines.Text() != "" {
			trees = append(trees, fillArray(lines.Text()))
		}
	}

	result := getVisibilityScore(trees)
	fmt.Printf("(Part 2) Highest visibility score: %d\n", result)
	return result
}
