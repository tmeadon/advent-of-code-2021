package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var filePath = "input.txt"

type Point struct {
	row int
	col int
}

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	heightMap := loadInput()
	lowPoints, _ := findLowPoints(heightMap)
	risk := calculateRiskLevel(lowPoints)
	fmt.Println("Part 1: ", risk)
}

func solvePart2() {
	heightMap := loadInput()
	_, lowPoints := findLowPoints(heightMap)
	basinSizes := make([]int, len(lowPoints))

	for _, lowPoint := range lowPoints {
		basinSize := discoverBasinSize(lowPoint, heightMap)
		basinSizes = append(basinSizes, basinSize)
	}

	sort.Ints(basinSizes)
	l := len(basinSizes)
	productOfTopThreeBasins := basinSizes[l-1] * basinSizes[l-2] * basinSizes[l-3]
	fmt.Println("Part 2: ", productOfTopThreeBasins)
}

func loadInput() [][]int {
	input := make([][]int, 0)
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineNums := make([]int, 0)
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			lineNums = append(lineNums, num)
		}
		input = append(input, lineNums)
	}

	return input
}

func findLowPoints(heightMap [][]int) ([]int, []Point) {
	lowPoints := make([]int, 0)
	lowPointCoords := make([]Point, 0)

	for i, row := range heightMap {
		for j, col := range row {
			if isLowPoint(heightMap, i, j) {
				lowPoints = append(lowPoints, col)
				lowPointCoords = append(lowPointCoords, Point{row: i, col: j})
			}
		}
	}

	return lowPoints, lowPointCoords
}

func isLowPoint(heightMap [][]int, i, j int) bool {
	numRows := len(heightMap)
	numColumns := len(heightMap[0])

	// if we're at the top left corner
	if i == 0 && j == 0 {
		return heightMap[i][j] < heightMap[i+1][j] && heightMap[i][j] < heightMap[i][j+1]
	}

	// if we're at the top right corner
	if i == 0 && j == numColumns-1 {
		return heightMap[i][j] < heightMap[i+1][j] && heightMap[i][j] < heightMap[i][j-1]
	}

	// if we're at the bottom left corner
	if i == numRows-1 && j == 0 {
		return heightMap[i][j] < heightMap[i-1][j] && heightMap[i][j] < heightMap[i][j+1]
	}

	// if we're at the bottom right corner
	if i == numRows-1 && j == numColumns-1 {
		return heightMap[i][j] < heightMap[i-1][j] && heightMap[i][j] < heightMap[i][j-1]
	}

	// if we're at the top row
	if i == 0 {
		return heightMap[i][j] < heightMap[i+1][j] && heightMap[i][j] < heightMap[i][j-1] && heightMap[i][j] < heightMap[i][j+1]
	}

	// if we're at the bottom row
	if i == numRows-1 {
		return heightMap[i][j] < heightMap[i-1][j] && heightMap[i][j] < heightMap[i][j-1] && heightMap[i][j] < heightMap[i][j+1]
	}

	// if we're at the left most column
	if j == 0 {
		return heightMap[i][j] < heightMap[i-1][j] && heightMap[i][j] < heightMap[i+1][j] && heightMap[i][j] < heightMap[i][j+1]
	}

	// if we're at the right most column
	if j == numColumns-1 {
		return heightMap[i][j] < heightMap[i-1][j] && heightMap[i][j] < heightMap[i+1][j] && heightMap[i][j] < heightMap[i][j-1]
	}

	// if we're in the middle
	return heightMap[i][j] < heightMap[i-1][j] && heightMap[i][j] < heightMap[i+1][j] &&
		heightMap[i][j] < heightMap[i][j-1] && heightMap[i][j] < heightMap[i][j+1]
}

func calculateRiskLevel(lowPoints []int) int {
	riskLevel := 0
	for _, point := range lowPoints {
		riskLevel += point + 1
	}
	return riskLevel
}

func discoverBasinSize(lowPoint Point, heightMap [][]int) int {
	basin := make([]Point, 0)
	basin = expandBasin(basin, heightMap, lowPoint)
	return len(basin)
}

func expandBasin(basin []Point, heightMap [][]int, start Point) []Point {
	if !basinContainsPoint(start, basin) {
		basin = append(basin, start)

		if start.row > 0 {
			if heightMap[start.row-1][start.col] != 9 {
				basin = expandBasin(basin, heightMap, Point{start.row - 1, start.col})
			}
		}
		if start.row < len(heightMap)-1 {
			if heightMap[start.row+1][start.col] != 9 {
				basin = expandBasin(basin, heightMap, Point{start.row + 1, start.col})
			}
		}
		if start.col > 0 {
			if heightMap[start.row][start.col-1] != 9 {
				basin = expandBasin(basin, heightMap, Point{start.row, start.col - 1})
			}
		}
		if start.col < len(heightMap[0])-1 {
			if heightMap[start.row][start.col+1] != 9 {
				basin = expandBasin(basin, heightMap, Point{start.row, start.col + 1})
			}
		}
	}

	return basin
}

func basinContainsPoint(point Point, basin []Point) bool {
	for _, p := range basin {
		if p == point {
			return true
		}
	}
	return false
}
