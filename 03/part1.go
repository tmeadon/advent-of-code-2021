package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := getInputSlice("input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(input []string) {
	maxInputLength := getInputValuesMaxLength(input)
	gammaBinary := buildGammaBinaryString(input, maxInputLength)
	epsilonBinary := convertGammaBinaryToEpsilonBinary(gammaBinary)
	gammaDecimal := convertBinaryToDecimal(gammaBinary)
	epsilonDecimal := convertBinaryToDecimal(epsilonBinary)
	fmt.Println("Power consumption: ", (gammaDecimal * epsilonDecimal))
}

func getInputSlice(filePath string) []string {
	inputFileBytes, _ := ioutil.ReadFile(filePath)
	input := strings.Split(string(inputFileBytes), "\r\n")
	return input
}

func getInputValuesMaxLength(input []string) int {
	maxLength := 0
	for _, line := range input {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}
	return maxLength
}

func buildGammaBinaryString(input []string, maxInputLength int) string {
	var sb strings.Builder

	for i := 0; i < maxInputLength; i++ {
		mostCommon := getMostCommonBitValue(input, i)
		sb.WriteString(fmt.Sprint(mostCommon))
	}

	return sb.String()
}

func getMostCommonBitValue(input []string, bitIndex int) int {
	bitValues := getAllBitValues(input, bitIndex)
	return getMostFrequentElement(bitValues)
}

func getAllBitValues(input []string, bitIndex int) []int {
	bitValues := []int{}
	for _, line := range input {
		bitValues = append(bitValues, getBitValue(line, bitIndex))
	}
	return bitValues
}

func getBitValue(input string, bitIndex int) int {
	value, _ := strconv.Atoi(string(input[bitIndex]))
	return value
}

func getMostFrequentElement(input []int) int {
	elements := map[int]int{}
	var maxCount int
	var frequency int

	for _, value := range input {
		elements[value]++
		if elements[value] > maxCount {
			maxCount = elements[value]
			frequency = value
		}
	}

	return frequency
}

func convertGammaBinaryToEpsilonBinary(input string) string {
	var sb strings.Builder

	for i := 0; i < len(input); i++ {
		thisChar := input[i]
		if thisChar == '0' {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
	}

	return sb.String()
}

func convertBinaryToDecimal(binary string) int64 {
	decimal, _ := strconv.ParseInt(binary, 2, 64)
	return decimal
}

func solvePart2(input []string) {
	oxygenRating := getOxygenRatingDecimal(input)
	co2Rating := getCo2RatingDecimal(input)
	fmt.Println("Oxygen rating: ", oxygenRating)
	fmt.Println("CO2 rating: ", co2Rating)
	fmt.Println("Life support rating: ", (oxygenRating * co2Rating))
}

func getOxygenRatingDecimal(input []string) int64 {
	oxygenRatingBinary := getOxygenRatingBinary(input)
	return convertBinaryToDecimal(oxygenRatingBinary)
}

func getOxygenRatingBinary(input []string) string {
	rating := filterValues(input, true)
	return rating
}

func getCo2RatingDecimal(input []string) int64 {
	co2RatingBinary := getCo2RatingBinary(input)
	return convertBinaryToDecimal(co2RatingBinary)
}

func getCo2RatingBinary(input []string) string {
	rating := filterValues(input, false)
	return rating
}

func filterValues(input []string, useMostCommon bool) string {
	bitIndex := 0

	for {
		mostCommonList := []string{}
		leastCommonList := []string{}
		allBitValues := getAllBitValues(input, bitIndex)
		numZeros := countNumber(allBitValues, 0)
		numOnes := countNumber(allBitValues, 1)
		var mostCommon int
		var leastCommon int

		if numOnes >= numZeros {
			mostCommon = 1
			leastCommon = 0
		} else {
			mostCommon = 0
			leastCommon = 1
		}

		for _, number := range input {
			bitValue := getBitValue(number, bitIndex)
			if bitValue == mostCommon {
				mostCommonList = append(mostCommonList, number)
			}
			if bitValue == leastCommon {
				leastCommonList = append(leastCommonList, number)
			}
		}

		if useMostCommon {
			input = mostCommonList
		} else {
			input = leastCommonList
		}

		if len(input) == 1 {
			break
		}

		bitIndex++
	}

	return input[0]
}

func countNumber(bitValues []int, number int) int {
	count := 0
	for _, value := range bitValues {
		if value == number {
			count++
		}
	}
	return count
}
