package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var filePath = "input.txt"

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	neighbours := loadInput()
	paths := findPaths(neighbours, false)
	fmt.Print("Part 1: ", len(paths))
}

func solvePart2() {
	neighbours := loadInput()
	paths := findPaths(neighbours, true)
	fmt.Print("Part 2: ", len(paths))
}

func loadInput() (neighbours map[string][]string) {
	neighbours = make(map[string][]string)
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		start := strings.Split(line, "-")[0]
		end := strings.Split(line, "-")[1]
		neighbours[start] = append(neighbours[start], end)
		neighbours[end] = append(neighbours[end], start)
	}

	return
}

func findPaths(neighbours map[string][]string, part2 bool) (paths []string) {
	paths = make([]string, 0)
	takeStep("start", neighbours, &paths, part2)
	return
}

func takeStep(path string, neighbours map[string][]string, paths *[]string, part2 bool) {
	pathSteps := strings.Split(path, ",")
	lastStep := pathSteps[len(pathSteps)-1]

	if lastStep == "end" {
		*paths = append(*paths, path)
		return
	}

	if !part2 && pathHasTwiceVisitedSameSmallCave(path) {
		return
	}

	if part2 && partHasVisitedMoreThanOneSmallCaveTwice(path) {
		return
	}

	for _, neighbour := range neighbours[lastStep] {
		if neighbour != "start" {
			newPath := path + "," + neighbour
			takeStep(newPath, neighbours, paths, part2)
		}
	}
}

func pathHasTwiceVisitedSameSmallCave(path string) bool {
	smallCavesVisited := make(map[string]int)
	pathSteps := strings.Split(path, ",")
	for _, step := range pathSteps {
		if isSmallCave(step) {
			smallCavesVisited[step]++
			if smallCavesVisited[step] > 1 {
				return true
			}
		}
	}
	return false
}

func partHasVisitedMoreThanOneSmallCaveTwice(path string) bool {
	smallCavesVisited := make(map[string]int)
	pathSteps := strings.Split(path, ",")
	numCavesVisitedTwice := 0

	for _, step := range pathSteps {
		if isSmallCave(step) {
			smallCavesVisited[step]++
		}
	}

	for _, v := range smallCavesVisited {
		if v > 2 {
			return true
		}
		if v == 2 {
			numCavesVisitedTwice++
		}
	}

	return numCavesVisitedTwice > 1
}

func isSmallCave(cave string) bool {
	for _, c := range cave {
		if !unicode.IsLower(c) {
			return false
		}
	}
	return true
}
