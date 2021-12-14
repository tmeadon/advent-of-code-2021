package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var filePath = "input.txt"

type Position struct {
	row int
	col int
}

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	squids := loadInput()
	numFlashed := 0

	for i := 0; i < 100; i++ {
		numFlashed += step(&squids)
	}

	fmt.Println("Part 1: ", numFlashed)
}

func solvePart2() {
	squids := loadInput()
	numSquids := len(squids) * len(squids[0])
	firstSimultaneousFlash := 0

	for i := 0; i < 1000; i++ {
		numFlashed := step(&squids)
		if numFlashed == numSquids {
			firstSimultaneousFlash = i + 1
			break
		}
	}

	fmt.Println("Part 2: ", firstSimultaneousFlash)
}

func loadInput() [][]int {
	input := make([][]int, 0)
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0)
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
		}
		input = append(input, row)
	}

	return input
}

func step(squids *[][]int) (numFlashed int) {
	flashedSquids := make([]Position, 0)

	// increase the energy of all squids
	for i := 0; i < len(*squids); i++ {
		for j := 0; j < len((*squids)[i]); j++ {
			(*squids)[i][j]++
		}
	}

	// keep flashing squid until they are all done
	for {
		numFlashed = len(flashedSquids)
		flashSquids(squids, &flashedSquids)
		if len(flashedSquids) == numFlashed {
			break
		}
		numFlashed = len(flashedSquids)
	}

	// finally set all squids with energy < 9 to 0
	for i := 0; i < len(*squids); i++ {
		for j := 0; j < len((*squids)[i]); j++ {
			if (*squids)[i][j] > 9 {
				(*squids)[i][j] = 0
			}
		}
	}

	return numFlashed
}

func flashSquids(squids *[][]int, flashedSquids *[]Position) {
	for i := 0; i < len(*squids); i++ {
		for j := 0; j < len((*squids)[i]); j++ {
			if (*squids)[i][j] == 10 && !squidHasFlashed(i, j, flashedSquids) {
				flashSquid(squids, i, j, flashedSquids)
			}
		}
	}
}

func squidHasFlashed(row int, col int, flashedSquids *[]Position) bool {
	for _, pos := range *flashedSquids {
		if pos.row == row && pos.col == col {
			return true
		}
	}
	return false
}

func flashSquid(squids *[][]int, row int, col int, flashedSquids *[]Position) {
	*flashedSquids = append(*flashedSquids, Position{row, col})

	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i >= 0 && i < len(*squids) && j >= 0 && j < len((*squids)[i]) {
				if i != row || j != col {
					if (*squids)[i][j] == 10 && !squidHasFlashed(i, j, flashedSquids) {
						flashSquid(squids, i, j, flashedSquids)
					}
					(*squids)[i][j]++
				}
			}
		}
	}
}

// func step(squids *[][]int) (numFlashed int) {
// 	// increase the energy of all squids
// 	for i := 0; i < len(*squids); i++ {
// 		for j := 0; j < len((*squids)[i]); j++ {
// 			(*squids)[i][j]++
// 		}
// 	}

// 	// flash all squids with energy > 9
// 	for i := 0; i < len(*squids); i++ {
// 		for j := 0; j < len((*squids)[i]); j++ {
// 			if (*squids)[i][j] == 10 {
// 				numFlashed += flashSquid(squids, i, j)
// 			}
// 		}
// 	}

// 	// finally set all squids with energy < 9 to 0
// 	for i := 0; i < len(*squids); i++ {
// 		for j := 0; j < len((*squids)[i]); j++ {
// 			if (*squids)[i][j] > 9 {
// 				(*squids)[i][j] = 0
// 			}
// 		}
// 	}

// 	return numFlashed
// }

// func flashSquid(squids *[][]int, row int, col int) (numFlashed int) {
// 	numFlashed++

// 	for i := row - 1; i <= row+1; i++ {
// 		for j := col - 1; j <= col+1; j++ {
// 			if i >= 0 && i < len(*squids) && j >= 0 && j < len((*squids)[i]) {
// 				(*squids)[i][j]++
// 				if (*squids)[i][j] == 10 {
// 					numFlashed += flashSquid(squids, i, j)
// 				}
// 			}
// 		}
// 	}

// 	return numFlashed
// }

// func flashSquid(squids *[][]int, row int, col int, flashedSquids *[]Position) {
// 	*flashedSquids = append(*flashedSquids, Position{row, col})

// 	// scan the table and increase the energy of all adjacent squids
// 	for i := row - 1; i <= row+1; i++ {
// 		for j := col - 1; j <= col+1; j++ {
// 			if i >= 0 && i < len(*squids) && j >= 0 && j < len((*squids)[i]) {
// 				// if (*squids)[i][j] != 0 && !squidHasFlashed(i, j, flashedSquids) {
// 				// 	increaseSquidEnergy(squids, i, j, flashedSquids)
// 				// 	if (*squids)[i][j] == 0 {
// 				// 		flashSquid(squids, i, j, flashedSquids)
// 				// 	}
// 				// }
// 				increaseSquidEnergy(squids, i, j)
// 				if (*squids)[i][j] == 0 {
// 					flashSquid(squids, i, j, flashedSquids)
// 				}
// 			}
// 		}
// 	}

// 	numFlashes++
// }

// func step(squids *[][]int) {
// 	flashedSquids := make([]Position, 0)

// 	for i := 0; i < len(*squids); i++ {
// 		for j := 0; j < len((*squids)[i]); j++ {
// 			increaseSquidEnergy(squids, i, j, &flashedSquids)
// 		}
// 	}

// 	// reset the energy for squids that flashed
// 	for _, pos := range flashedSquids {
// 		(*squids)[pos.row][pos.col] = 0
// 	}
// }

// func increaseSquidEnergy(squids *[][]int, i int, j int, flashedSquids *[]Position) {
// 	if (*squids)[i][j] == 9 {
// 		flashSquid(squids, i, j, flashedSquids)
// 	} else {
// 		(*squids)[i][j]++
// 	}
// }

// func flashSquid(squids *[][]int, row int, col int, flashedSquids *[]Position) {
// 	// set this squid to 0
// 	(*squids)[row][col] = 0

// 	// scan the table and increase the energy of all adjacent squids
// 	for i := row - 1; i <= row+1; i++ {
// 		for j := col - 1; j <= col+1; j++ {
// 			if i >= 0 && i < len(*squids) && j >= 0 && j < len((*squids)[i]) {
// 				if (*squids)[i][j] != 0 && !squidHasFlashed(i, j, flashedSquids) {
// 					increaseSquidEnergy(squids, i, j, flashedSquids)
// 				}
// 			}
// 		}
// 	}

// 	*flashedSquids = append(*flashedSquids, Position{row, col})
// 	numFlashes++
// }

// func squidHasFlashed(row int, col int, flashedSquids *[]Position) bool {
// 	for _, pos := range *flashedSquids {
// 		if pos.row == row && pos.col == col {
// 			return true
// 		}
// 	}

// 	return false
// }

// // if we're in the top left corner
// if row == 0 && col == 0 {
// 	increaseSquidEnergy(squids, row, col+1)
// 	increaseSquidEnergy(squids, row+1, col)
// 	increaseSquidEnergy(squids, row+1, col+1)
// 	return
// }

// // if we're in the top right corner
// if row == 0 && col == len((*squids)[row])-1 {
// 	increaseSquidEnergy(squids, row, col-1)
// 	increaseSquidEnergy(squids, row+1, col)
// 	increaseSquidEnergy(squids, row+1, col-1)
// 	return
// }

// // if we're in the bottom left corner
// if row == len((*squids))-1 && col == 0 {
// 	increaseSquidEnergy(squids, row, col+1)
// 	increaseSquidEnergy(squids, row-1, col)
// 	increaseSquidEnergy(squids, row-1, col+1)
// 	return
// }

// // if we're in the bottom right corner
// if row == len((*squids))-1 && col == len((*squids)[row])-1 {
// 	increaseSquidEnergy(squids, row, col-1)
// 	increaseSquidEnergy(squids, row-1, col)
// 	increaseSquidEnergy(squids, row-1, col-1)
// 	return
// }

// // if we're in the top row
// if row == 0 {
// 	increaseSquidEnergy(squids, row, col-1)
// 	increaseSquidEnergy(squids, row, col+1)
// 	increaseSquidEnergy(squids, row+1, col)
// 	increaseSquidEnergy(squids, row+1, col-1)
// 	increaseSquidEnergy(squids, row+1, col+1)
// 	return
// }

// // if we're in the bottom row
// if row == len((*squids))-1 {
// 	increaseSquidEnergy(squids, row, col-1)
// 	increaseSquidEnergy(squids, row, col+1)
// 	increaseSquidEnergy(squids, row-1, col)
// 	increaseSquidEnergy(squids, row-1, col-1)
// 	increaseSquidEnergy(squids, row-1, col+1)
// 	return
// }

// // if we're in the left most column
// if col == 0 {
// 	increaseSquidEnergy(squids, row-1, col)
// 	increaseSquidEnergy(squids, row+1, col)
// 	increaseSquidEnergy(squids, row, col+1)

// 	return
// }

// // if we're in the right most column
// if col == len((*squids)[row])-1 {
// 	increaseSquidEnergy(squids, row-1, col)
// 	increaseSquidEnergy(squids, row+1, col)
// 	increaseSquidEnergy(squids, row, col-1)
// 	return
// }

// // if we're in the middle

// }
