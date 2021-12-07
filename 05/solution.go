package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var gridSize = 1000
var filePath = "input.txt"

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	grid := initGrid(gridSize)

	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		line := parseLine(text)
		if isVertical(line) {
			drawVerticalLine(grid, line)
		} else if isHorizontal(line) {
			drawHorizontalLine(grid, line)
		}
	}

	overlaps := countOverlaps(grid)
	fmt.Println("Part 1: ", overlaps)
}

func solvePart2() {
	grid := initGrid(gridSize)

	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		line := parseLine(text)
		if isVertical(line) {
			drawVerticalLine(grid, line)
		} else if isHorizontal(line) {
			drawHorizontalLine(grid, line)
		} else {
			drawDiagonalLine(grid, line)
		}
	}

	overlaps := countOverlaps(grid)
	fmt.Println("Part 2: ", overlaps)
}

func initGrid(size int) [][]int {
	grid := make([][]int, size)

	for i := 0; i < size; i++ {
		grid[i] = make([]int, size)
	}

	return grid
}

func parseLine(line string) Line {
	var lineStart, lineEnd Point
	coords := strings.Split(line, " -> ")
	start := coords[0]
	end := coords[1]
	lineStart.x, _ = strconv.Atoi(strings.Split(start, ",")[0])
	lineStart.y, _ = strconv.Atoi(strings.Split(start, ",")[1])
	lineEnd.x, _ = strconv.Atoi(strings.Split(end, ",")[0])
	lineEnd.y, _ = strconv.Atoi(strings.Split(end, ",")[1])
	return Line{lineStart, lineEnd}
}

func isVertical(line Line) bool {
	return line.start.x == line.end.x
}

func isHorizontal(line Line) bool {
	return line.start.y == line.end.y
}

func drawHorizontalLine(grid [][]int, line Line) {
	if line.start.x < line.end.x {
		for i := line.start.x; i <= line.end.x; i++ {
			grid[i][line.start.y]++
		}
	} else {
		for i := line.end.x; i <= line.start.x; i++ {
			grid[i][line.start.y]++
		}
	}
}

func drawVerticalLine(grid [][]int, line Line) {
	if line.start.y < line.end.y {
		for i := line.start.y; i <= line.end.y; i++ {
			grid[line.start.x][i]++
		}
	} else {
		for i := line.end.y; i <= line.start.y; i++ {
			grid[line.start.x][i]++
		}
	}
}

func drawDiagonalLine(grid [][]int, line Line) {

	// line heading up and to the right
	if line.start.x < line.end.x && line.start.y < line.end.y {
		for i := 0; i <= (line.end.x - line.start.x); i++ {
			grid[line.start.x+i][line.start.y+i]++
		}
	}

	// line heading up and to the left
	if line.start.x > line.end.x && line.start.y < line.end.y {
		for i := 0; i <= (line.start.x - line.end.x); i++ {
			grid[line.start.x-i][line.start.y+i]++
		}
	}

	// line heading down and to the right
	if line.start.x < line.end.x && line.start.y > line.end.y {
		for i := 0; i <= (line.end.x - line.start.x); i++ {
			grid[line.start.x+i][line.start.y-i]++
		}
	}

	// line heading down and to the left
	if line.start.x > line.end.x && line.start.y > line.end.y {
		for i := 0; i <= (line.start.x - line.end.x); i++ {
			grid[line.start.x-i][line.start.y-i]++
		}
	}
}

func countOverlaps(grid [][]int) int {
	count := 0
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if grid[i][j] > 1 {
				count++
			}
		}
	}
	return count
}
