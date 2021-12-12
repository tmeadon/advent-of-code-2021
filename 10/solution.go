package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var filePath = "input.txt"
var charPairs = map[string]string{
	"}": "{",
	"]": "[",
	")": "(",
	">": "<",
}
var syntaxErrorPoints = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}
var syntaxCompletionPoints = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

// define a stack and some functions to keep track of the number of open and closed characters
type CharStack struct {
	chars []string
	count int
}

func NewCharStack() *CharStack {
	return &CharStack{
		chars: make([]string, 0),
		count: 0,
	}
}

func (cs *CharStack) Push(char string) {
	cs.chars = append(cs.chars[:cs.count], char)
	cs.count++
}

func (cs *CharStack) Pop() string {
	if cs.count == 0 {
		return ""
	}
	cs.count--
	return cs.chars[cs.count]
}

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	input := loadInput()
	errorScore := 0

	for _, line := range input {
		illegalChar, found := findIllegalCharacter(line)
		if found {
			score := syntaxErrorPoints[illegalChar]
			errorScore += score
		}
	}

	fmt.Println("Part 1: ", errorScore)
}

func solvePart2() {
	input := loadInput()
	completionScores := make([]int, 0)

	for _, line := range input {
		_, found := findIllegalCharacter(line)
		if !found {
			completion := completeMissingLines(line)
			score := calculateCompletionScore(completion)
			completionScores = append(completionScores, score)
		}
	}

	sort.Ints(completionScores)
	midIndex := len(completionScores) / 2
	fmt.Println("Part 2: ", completionScores[midIndex])
}

func loadInput() []string {
	input := make([]string, 0)
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func findIllegalCharacter(line string) (char string, found bool) {
	charStack := NewCharStack()

	for _, char := range line {
		if isClosingChar(string(char)) {
			if charStack.Pop() != charPairs[string(char)] {
				return string(char), true
			}
		} else {
			charStack.Push(string(char))
		}
	}

	return "", false
}

func isClosingChar(char string) bool {
	for closingChar := range charPairs {
		if char == closingChar {
			return true
		}
	}
	return false
}

func completeMissingLines(line string) []string {
	charStack := NewCharStack()
	completion := make([]string, 0)

	for _, char := range line {
		if isClosingChar(string(char)) {
			charStack.Pop()
		} else {
			charStack.Push(string(char))
		}
	}

	numIncomplete := charStack.count

	for i := 0; i <= numIncomplete; i++ {
		char := charStack.Pop()
		if char == "" {
			break
		}
		completion = append(completion, charPairReverseLookup(char))
	}

	return completion
}

func calculateCompletionScore(completion []string) int {
	score := 0
	for _, char := range completion {
		score = score * 5
		score += syntaxCompletionPoints[char]
	}
	return score
}

func charPairReverseLookup(char string) string {
	for key, value := range charPairs {
		if value == char {
			return key
		}
	}
	return ""
}
