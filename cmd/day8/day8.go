package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Coordinates struct {
	row int
	col int
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}

	data, err := os.ReadFile(filepath.Join(cwd, "inputs", "day8.txt"))
	if err != nil {
		os.Exit(2)
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	rows := len(lines) - 1
	cols := len(lines[1])
	fmt.Println("Rows: ", rows, " cols: ", cols)
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}
	for rowIdx, rowValue := range lines {
		for colIdx, value := range rowValue {
			asInt, err := strconv.Atoi(string(value))
			if err != nil {
				os.Exit(3)
			}
			grid[rowIdx][colIdx] = asInt
		}
	}

	totalVisible := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if isVisible(grid, i, j) {
				totalVisible++
			}
		}
	}

	maxScenic := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			ss := scenicScore(grid, i, j)
			if ss > maxScenic {
				maxScenic = ss
			}
		}
	}
	printGrid(grid)
	totalVisible = totalVisible + 2*(rows+cols) - 4
	fmt.Println("Total visible trees:", totalVisible)
	fmt.Println("Max scenic score: ", maxScenic)
}

func scenicScore(grid [][]int, row int, col int) int {
	currentElement := grid[row][col]
	scenicUp := 0
	scenicDown := 0
	scenicLeft := 0
	scenicRight := 0
	for i := row - 1; i >= 0; i-- {
		scenicUp++
		if grid[i][col] >= currentElement {
			break
		}
	}

	for i := row + 1; i < len(grid); i++ {
		scenicDown++
		if grid[i][col] >= currentElement {
			break
		}
	}

	for j := col - 1; j >= 0; j-- {
		scenicLeft++
		if grid[row][j] >= currentElement {
			break
		}
	}
	for j := col + 1; j < len(grid[row]); j++ {
		scenicRight++
		if grid[row][j] >= currentElement {
			break
		}
	}
	fmt.Println("element: ", currentElement, "scores: ", scenicUp, scenicDown, scenicLeft, scenicRight)

	return scenicUp * scenicDown * scenicLeft * scenicRight
}

func isVisible(grid [][]int, row int, col int) bool {
	return isVisibleX(grid, row, col) || isVisibleY(grid, row, col)
}

func isVisibleX(grid [][]int, row int, col int) bool {
	currentElem := grid[row][col]
	visibleLeft := true
	for _, val := range grid[row][0:col] {
		if currentElem <= val {
			visibleLeft = false
		}
	}
	visibleRight := true
	for _, val := range grid[row][col+1:] {
		if currentElem <= val {
			visibleRight = false
		}
	}
	return visibleLeft || visibleRight
}

func isVisibleY(grid [][]int, row int, col int) bool {
	currentElem := grid[row][col]
	visibleUp := true
	for i := 0; i < row; i++ {
		if currentElem <= grid[i][col] {
			visibleUp = false
		}
	}

	visibleDown := true
	for i := row + 1; i < len(grid); i++ {
		if currentElem <= grid[i][col] {
			visibleDown = false
		}
	}
	return visibleUp || visibleDown
}

func printGrid(grid [][]int) {
	for i := range grid {
		fmt.Printf(("[ "))
		for j := range grid[i] {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Printf("]\n")
	}
}

func indexOf(arr []int, val int) int {
	for pos, v := range arr {
		if v == val {
			return pos
		}
	}
	return -1
}
