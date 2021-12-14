package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputFile = "input.txt"
var outputFile = "output.txt"

type dot struct {
	x int
	y int
}

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	dotMap := loadDotMap()
	folds := loadFoldInstructions()
	dotMap = fold(dotMap, folds[0])
	fmt.Println("Part 1: ", len(dotMap))
}

func solvePart2() {
	dotMap := loadDotMap()
	folds := loadFoldInstructions()

	for _, instruction := range folds {
		dotMap = fold(dotMap, instruction)
	}

	writeDotsToFile(dotMap)
}

func loadDotMap() map[dot]bool {
	dotMap := make(map[dot]bool)
	file, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(file)

	// read the dots and stop at the first blank line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		x, _ := strconv.Atoi(strings.Split(line, ",")[0])
		y, _ := strconv.Atoi(strings.Split(line, ",")[1])
		dotMap[dot{x, y}] = true
	}

	return dotMap
}

func loadFoldInstructions() []string {
	folds := make([]string, 0)
	file, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "fold along") {
			folds = append(folds, strings.Replace(line, "fold along ", "", -1))
		}
	}

	return folds
}

func fold(dotMap map[dot]bool, foldInstruction string) map[dot]bool {
	if strings.Contains(foldInstruction, "y=") {
		dotMap = foldUp(dotMap, getFoldLine(foldInstruction))
	} else if strings.Contains(foldInstruction, "x=") {
		dotMap = foldLeft(dotMap, getFoldLine(foldInstruction))
	}
	return dotMap
}

func getFoldLine(foldInstruction string) int {
	foldLine, _ := strconv.Atoi(strings.Split(foldInstruction, "=")[1])
	return foldLine
}

func foldUp(dotMap map[dot]bool, foldLine int) map[dot]bool {
	for thisDot := range dotMap {
		if thisDot.y > foldLine {
			newY := foldLine - (thisDot.y - foldLine)
			dotMap[dot{thisDot.x, newY}] = true
			delete(dotMap, thisDot)
		}
	}
	return dotMap
}

func foldLeft(dotMap map[dot]bool, foldLine int) map[dot]bool {
	for thisDot := range dotMap {
		if thisDot.x > foldLine {
			newX := foldLine - (thisDot.x - foldLine)
			dotMap[dot{newX, thisDot.y}] = true
			delete(dotMap, thisDot)
		}
	}
	return dotMap
}

func writeDotsToFile(dotMap map[dot]bool) {
	width, height := getPageSize(dotMap)
	dotMapString := buildDotMapString(dotMap, height, width)
	file, _ := os.Create(outputFile)
	defer file.Close()
	file.WriteString(dotMapString)
}

func getPageSize(dotMap map[dot]bool) (maxX, maxY int) {
	for thisDot := range dotMap {
		if thisDot.x > maxX {
			maxX = thisDot.x
		}
		if thisDot.y > maxY {
			maxY = thisDot.y
		}
	}
	return
}

func buildDotMapString(dotMap map[dot]bool, height, width int) string {
	var sb strings.Builder

	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			if dotMap[dot{x, y}] {
				sb.WriteString("#")
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
