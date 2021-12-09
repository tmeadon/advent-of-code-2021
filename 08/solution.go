package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var filePath = "input.txt"

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	_, outputValues := readInputFile()
	identifiableDigits := countIdentifiableDigits(outputValues)
	fmt.Println("Part 1: ", identifiableDigits)
}

func solvePart2() {
	sum := 0
	signalPatterns, outputValues := readInputFile()

	for i := 0; i < len(signalPatterns); i++ {
		resolvedDigits := resolveDigits(signalPatterns[i])
		output := decodeOutput(outputValues[i], &resolvedDigits)
		sum += output
	}

	fmt.Println("Part 2: ", sum)
}

func readInputFile() (signalPatterns [][]string, outputValues [][]string) {
	signalPatterns = make([][]string, 0)
	outputValues = make([][]string, 0)
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		signals := strings.Split(line, " | ")[0]
		outputs := strings.Split(line, " | ")[1]
		signalPatterns = append(signalPatterns, strings.Fields(signals))
		outputValues = append(outputValues, strings.Fields(outputs))
	}

	return signalPatterns, outputValues
}

func countIdentifiableDigits(outputValues [][]string) int {
	count := 0

	for i := 0; i < len(outputValues); i++ {
		for j := 0; j < len(outputValues[i]); j++ {
			if digitIsIdentifiable(outputValues[i][j]) {
				count++
			}
		}
	}

	return count
}

func digitIsIdentifiable(digit string) bool {
	l := len(digit)
	return (l == 2 || l == 3 || l == 4 || l == 7)
}

func resolveDigits(signalPattern []string) map[int]string {
	resolvedDigits := make(map[int]string)
	resolveIdentifiableDigits(signalPattern, &resolvedDigits)
	findNumberThree(signalPattern, &resolvedDigits)
	findNumberTwo(signalPattern, &resolvedDigits)
	findNumberFive(signalPattern, &resolvedDigits)
	findNumberSix(signalPattern, &resolvedDigits)
	findNumberZero(signalPattern, &resolvedDigits)
	findNumberNine(signalPattern, &resolvedDigits)
	return resolvedDigits
}

// we already know how to find the numbers 1 (length 2), 4 (length 4), 7 (length 3) and 8 (length 7)
func resolveIdentifiableDigits(signalPattern []string, resolvedDigits *map[int]string) {
	for _, pattern := range signalPattern {
		l := len(pattern)
		switch l {
		case 2:
			(*resolvedDigits)[1] = pattern
		case 3:
			(*resolvedDigits)[7] = pattern
		case 4:
			(*resolvedDigits)[4] = pattern
		case 7:
			(*resolvedDigits)[8] = pattern
		}
	}
}

// number 3 is the only 5 char segment that contains both chars that make up number 1
func findNumberThree(signalPattern []string, resolvedDigits *map[int]string) {
	numberOnePattern := (*resolvedDigits)[1]

	for _, pattern := range signalPattern {
		if len(pattern) == 5 {
			match := true
			for _, char := range numberOnePattern {
				if !strings.Contains(pattern, string(char)) {
					match = false
					break
				}
			}

			if match {
				(*resolvedDigits)[3] = pattern
				break
			}
		}
	}
}

// number 2 is the 5 char segment that shares exactly two chars with number 4
func findNumberTwo(signalPattern []string, resolvedDigits *map[int]string) {
	numberFourPattern := (*resolvedDigits)[4]

	for _, pattern := range signalPattern {
		if len(pattern) == 5 {
			charsSharedWithFour := 0

			for _, char := range numberFourPattern {
				if strings.Contains(pattern, string(char)) {
					charsSharedWithFour++
				}
			}

			if charsSharedWithFour == 2 {
				(*resolvedDigits)[2] = pattern
				break
			}
		}
	}
}

// number 5 is the remaining 5 char segment
func findNumberFive(signalPattern []string, resolvedDigits *map[int]string) {
	for _, pattern := range signalPattern {
		if len(pattern) == 5 {
			if !isPatternResolved(pattern, resolvedDigits) {
				(*resolvedDigits)[5] = pattern
				break
			}
		}
	}
}

// number 6 is the only 6 char segment that doesn't contain all of the segments in number 1
func findNumberSix(signalPattern []string, resolvedDigits *map[int]string) {
	numberOnePattern := (*resolvedDigits)[1]

	for _, pattern := range signalPattern {
		if len(pattern) == 6 {
			match := false
			for _, char := range numberOnePattern {
				if !strings.Contains(pattern, string(char)) {
					match = true
					break
				}
			}

			if match {
				(*resolvedDigits)[6] = pattern
				break
			}
		}
	}
}

// number 0 is the only remaining 6 char segment that doesn't contain all of the segments in number 4
func findNumberZero(signalPattern []string, resolvedDigits *map[int]string) {
	numberFourPattern := (*resolvedDigits)[4]

	for _, pattern := range signalPattern {
		if len(pattern) == 6 {
			if pattern != (*resolvedDigits)[6] {
				match := false
				for _, char := range numberFourPattern {
					if !strings.Contains(pattern, string(char)) {
						match = true
						break
					}
				}

				if match {
					(*resolvedDigits)[0] = pattern
					break
				}
			}
		}
	}
}

// number 9 is the remaining 6 char segment
func findNumberNine(signalPattern []string, resolvedDigits *map[int]string) {
	for _, pattern := range signalPattern {
		if len(pattern) == 6 {
			if !isPatternResolved(pattern, resolvedDigits) {
				(*resolvedDigits)[9] = pattern
				break
			}
		}
	}
}

// checks if a pattern is already in the resolvedDigits map
func isPatternResolved(pattern string, resolvedDigits *map[int]string) bool {
	for _, digit := range *resolvedDigits {
		if digit == pattern {
			return true
		}
	}
	return false
}

func decodeOutput(outputValues []string, resolvedDigits *map[int]string) int {
	decoded := ""

	for _, value := range outputValues {
		for k, v := range *resolvedDigits {
			if areSignalsEqual(value, v) {
				decoded += strconv.Itoa(k)
			}
		}
	}

	number, _ := strconv.Atoi(decoded)

	return number
}

// checks if two strings have the same characters despite the order
func areSignalsEqual(signal1 string, signal2 string) bool {
	if len(signal1) != len(signal2) {
		return false
	}

	for i := 0; i < len(signal1); i++ {
		if !strings.Contains(signal2, string(signal1[i])) {
			return false
		}
	}

	return true
}
