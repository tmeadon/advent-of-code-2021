package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Board struct {
	cells [5][5]BoardCell
	won   bool
}

type BoardCell struct {
	number int
	marked bool
}

func main() {
	chosenNumbers := loadChosenNumbers("input.txt")
	boards := loadBoards("input.txt")
	firstWinnerFound := false
	var lastWinner *Board
	var lastWinningNumber int
	numWinners := 0

	for _, number := range chosenNumbers {
		drawNumber(number, &boards)
		newWinners := findWinners(&boards)

		if len(newWinners) > 0 {
			if !firstWinnerFound {
				part1Result := calculateResult(newWinners[0], number)
				fmt.Println("Part 1 result: ", part1Result)
				firstWinnerFound = true
			}
			markWinners(&boards, newWinners)
			lastWinner = newWinners[len(newWinners)-1]
			lastWinningNumber = number
			numWinners += len(newWinners)
		}

		if numWinners == len(boards) {
			break
		}
	}

	part2Result := calculateResult(lastWinner, lastWinningNumber)
	fmt.Println("Part 2 result: ", part2Result)
}

func loadChosenNumbers(filePath string) []int {
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)
	var numbersList string

	for scanner.Scan() {
		line := scanner.Text()
		numbersList = numbersList + line
		if line == "" {
			break
		}
	}

	var numbers []int

	for _, number := range strings.Split(numbersList, ",") {
		thisNum, _ := strconv.Atoi(number)
		numbers = append(numbers, thisNum)
	}

	return numbers
}

func loadBoards(filePath string) []Board {
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)
	var boards []Board
	thisBoard := Board{won: false}
	boardLineRegex := `(?:\s*\d+\s*){5}`
	boardCurrentLineNum := 0
	readingBoard := false

	for scanner.Scan() {
		thisLine := scanner.Text()
		matched, _ := regexp.Match(boardLineRegex, []byte(thisLine))

		if matched {
			readingBoard = true
			readBoardLine(&thisLine, &thisBoard.cells, &boardCurrentLineNum)
			boardCurrentLineNum++
		} else {
			finishReadingBoard(&boards, &thisBoard, &boardCurrentLineNum, &readingBoard)
		}
	}

	finishReadingBoard(&boards, &thisBoard, &boardCurrentLineNum, &readingBoard)
	return boards
}

func readBoardLine(thisLine *string, thisBoard *[5][5]BoardCell, boardCurrentLineNum *int) {
	for i, number := range strings.Fields(*thisLine) {
		thisNum, _ := strconv.Atoi(number)
		thisBoard[*boardCurrentLineNum][i] = BoardCell{thisNum, false}
	}
}

func finishReadingBoard(boards *[]Board, thisBoard *Board, boardCurrentLineNum *int, readingBoard *bool) {
	if *readingBoard {
		*boards = append(*boards, *thisBoard)
		*thisBoard = Board{won: false}
		*boardCurrentLineNum = 0
		*readingBoard = false
	}
}

func drawNumber(number int, boards *[]Board) {
	for b := 0; b < len(*boards); b++ {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if (*boards)[b].cells[i][j].number == number {
					(*boards)[b].cells[i][j].marked = true
				}
			}
		}
	}
}

func findWinners(boards *[]Board) []*Board {
	winners := []*Board{}
	for b := 0; b < len(*boards); b++ {
		if boardHasWinningSet(&((*boards)[b])) && !((*boards)[b].won) {
			winners = append(winners, &((*boards)[b]))
		}
	}
	return winners
}

func boardHasWinningSet(board *Board) bool {
	for i := 0; i < 5; i++ {
		winningRow := true
		winningColumn := true

		for j := 0; j < 5; j++ {
			if !board.cells[i][j].marked {
				winningRow = false
			}
			if !board.cells[j][i].marked {
				winningColumn = false
			}
			if j == 4 && (winningRow || winningColumn) {
				return true
			}
		}
	}

	return false
}

func removeBoard(boards [][5][5]BoardCell, boardToRemove [5][5]BoardCell) [][5][5]BoardCell {
	for i, board := range boards {
		if board == boardToRemove {
			boards = append(boards[:i], boards[i+1:]...)
		}
	}
	return boards
}

func calculateResult(board *Board, lastNumber int) int {
	sumOfUnmarked := 0

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !board.cells[i][j].marked {
				sumOfUnmarked += board.cells[i][j].number
			}
		}
	}

	return (sumOfUnmarked * lastNumber)
}

func markWinners(boards *[]Board, winners []*Board) {
	for i := 0; i < len(winners); i++ {
		for j := 0; j < len(*boards); j++ {
			if (*boards)[j] == *winners[i] {
				(*boards)[j].won = true
			}
		}
	}
}
