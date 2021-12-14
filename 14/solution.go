package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var filePath = "input.txt"

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	result := solve(10)
	fmt.Println("Part 1: ", result)
}

func solvePart2() {
	result := solve(40)
	fmt.Println("Part 2: ", result)
}

func solve(numSteps int) (result int) {
	polymer, lastChar := loadPolymerTemplate()
	rules := loadInsertionRules()
	polymer = expandPolymer(polymer, rules, numSteps)
	mostCommonCount, leastCommonCount := getMostLeastCommonCounts(polymer, lastChar)
	return mostCommonCount - leastCommonCount
}

func loadPolymerTemplate() (pairs map[string]int, lastChar string) {
	pairs = make(map[string]int)
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	template := scanner.Text()
	lastChar = template[len(template)-1:]

	for i := 0; i < len(template)-1; i++ {
		pair := string(template[i]) + string(template[i+1])
		pairs[pair]++
	}

	return
}

func loadInsertionRules() (rules map[string]string) {
	rules = make(map[string]string)
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "->") {
			targetPair := strings.Split(line, " -> ")[0]
			newElement := strings.Split(line, " -> ")[1]
			rules[targetPair] = newElement
		}
	}

	return
}

func expandPolymer(polymer map[string]int, rules map[string]string, numSteps int) map[string]int {
	for i := 0; i < numSteps; i++ {
		polymer = step(polymer, rules)
	}
	return polymer
}

func step(polymerPairs map[string]int, rules map[string]string) (newPolymerPairs map[string]int) {
	newPolymerPairs = make(map[string]int)

	for pair := range polymerPairs {
		for _, newPair := range getNewPairs(pair, rules) {
			newPolymerPairs[newPair] += polymerPairs[pair]
		}
	}

	return
}

func getNewPairs(pair string, rules map[string]string) (newPairs []string) {
	newPairs = make([]string, 0)
	newPairs = append(newPairs, string(pair[0])+string(rules[pair]))
	newPairs = append(newPairs, string(rules[pair])+string(pair[1]))
	return
}

func getMostLeastCommonCounts(pairs map[string]int, lastChar string) (mostCommonCount, leastCommonCount int) {
	letters := make(map[string]int)
	leastCommonCount = math.MaxInt64

	for pair := range pairs {
		letters[pair[0:1]] += pairs[pair]
	}

	letters[lastChar]++

	for _, v := range letters {
		if v > mostCommonCount {
			mostCommonCount = v
		}
		if v < leastCommonCount {
			leastCommonCount = v
		}
	}

	return
}
