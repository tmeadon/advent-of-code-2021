package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputFile = "input.txt"

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	result := runSimulation(80)
	fmt.Println("Part 1: ", result)
}

func solvePart2() {
	result := runSimulation(256)
	fmt.Println("Part 2: ", result)
}

func runSimulation(numDays int) (numFish int) {
	numbers := loadFishTimers()

	for i := 0; i < numDays; i++ {
		numbers = tick(numbers)
	}

	return countFish(numbers)
}

func loadFishTimers() map[int]int {
	numbers := newNumberMap()
	file, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		for _, value := range strings.Split(line, ",") {
			thisNumber, _ := strconv.Atoi(value)
			numbers[thisNumber]++
		}
		break
	}

	return numbers
}

func newNumberMap() map[int]int {
	numbers := make(map[int]int)

	for i := 0; i <= 8; i++ {
		numbers[i] = 0
	}

	return numbers
}

func tick(numbers map[int]int) map[int]int {
	newNumbers := newNumberMap()

	for i := (len(numbers) - 1); i >= 0; i-- {
		if i > 0 {
			newNumbers[i-1] = numbers[i]
		} else {
			newNumbers[6] += numbers[i]
			newNumbers[8] += numbers[i]
		}
	}

	return newNumbers
}

func countFish(numbers map[int]int) int {
	fishCount := 0

	for i := 0; i <= len(numbers); i++ {
		fishCount += numbers[i]
	}

	return fishCount
}
