package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
)

var filePath = "input.txt"

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	totalFuel := 0
	numbers := loadNumbers(filePath)
	median, _ := stats.Median(numbers)

	for _, number := range numbers {
		totalFuel += calculateFuelPart1(int(median), int(number))
	}

	println("Part 1: ", totalFuel)
}

func solvePart2() {
	totalFuel := math.Inf(1)
	numbers := loadNumbers(filePath)
	min, _ := stats.Min(numbers)
	max, _ := stats.Max(numbers)

	for i := min; i <= max; i++ {
		fuel := 0.0
		for _, number := range numbers {
			fuel += calculateFuelPart2(int(i), int(number))
		}
		if fuel < totalFuel {
			totalFuel = fuel
		}
	}

	println("Part 2: ", int(totalFuel))
}

func loadNumbers(inputFile string) []float64 {
	numbers := []float64{}
	file, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		for _, value := range strings.Split(line, ",") {
			thisNumber, _ := strconv.ParseFloat(value, 64)
			numbers = append(numbers, thisNumber)
		}
		break
	}

	return numbers
}

func calculateFuelPart1(target int, current int) int {
	if target < current {
		return (current - target)
	} else {
		return (target - current)
	}
}

func calculateFuelPart2(target int, current int) float64 {
	distance := math.Abs(float64(target - current))
	fuel := (distance * (distance + 1)) / 2
	return fuel
}
