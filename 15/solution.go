package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var filePath = "input.txt"

type Point struct{ x, y int }

type Todo struct {
	point Point
	risk  int
}

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	risks, width, height := loadRiskMap()
	leastRisk := getLeastRisk(risks, width, height)
	fmt.Println("Part 1:", leastRisk)
}

func solvePart2() {
	risks, width, height := loadRiskMap()
	fullMap, width, height := createFullRiskMap(risks, width, height)
	leastRisk := getLeastRisk(fullMap, width, height)
	fmt.Println("Part 2: ", leastRisk)
}

func loadRiskMap() (risks map[Point]int, width int, height int) {
	risks = make(map[Point]int)
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)
	rowNum := 0

	for scanner.Scan() {
		line := scanner.Text()
		for colNum, c := range line {
			risk, _ := strconv.Atoi(string(c))
			risks[Point{colNum, rowNum}] = risk
			if colNum > width {
				width = colNum + 1
			}
		}
		rowNum++
	}

	return risks, width, rowNum
}

func getLeastRisk(risks map[Point]int, width int, height int) int {
	pathRisk := make(map[Point]int)
	pathRisk[Point{0, 0}] = 0
	todoList := make([]Todo, 0)
	todoList = append(todoList, Todo{Point{0, 0}, 0})

	for {
		if len(todoList) == 0 {
			break
		}
		todo := popLowestRiskTodo(&todoList)
		for _, neighbour := range getNeighbours(todo.point) {
			if _, ok := risks[neighbour]; ok {
				neighbourRisk := getNeighbourRisk(pathRisk, neighbour)
				if neighbourRisk > (todo.risk + risks[neighbour]) {
					if isTodoInList(todoList, Todo{neighbour, neighbourRisk}) {
						removeTodo(&todoList, Todo{neighbour, neighbourRisk})
					}
					pathRisk[neighbour] = todo.risk + risks[neighbour]
					todoList = append(todoList, Todo{neighbour, pathRisk[neighbour]})
				}
			}
		}
	}

	return pathRisk[Point{width - 1, height - 1}]
}

func popLowestRiskTodo(todoList *[]Todo) Todo {
	lowestRisk := (*todoList)[0]
	for _, todo := range *todoList {
		if todo.risk < lowestRisk.risk {
			lowestRisk = todo
		}
	}
	removeTodo(todoList, lowestRisk)
	return lowestRisk
}

func removeTodo(todoList *[]Todo, todoToRemove Todo) {
	for i, todo := range *todoList {
		if todo.point == todoToRemove.point && todo.risk == todoToRemove.risk {
			*todoList = append((*todoList)[:i], (*todoList)[i+1:]...)
		}
	}
}

func getNeighbours(point Point) []Point {
	return []Point{
		{point.x - 1, point.y},
		{point.x + 1, point.y},
		{point.x, point.y - 1},
		{point.x, point.y + 1},
	}
}

func getNeighbourRisk(pathRisks map[Point]int, point Point) int {
	if risk, ok := pathRisks[point]; ok {
		return risk
	}
	return math.MaxInt64
}

func isTodoInList(todoList []Todo, todo Todo) bool {
	for _, todoInList := range todoList {
		if todoInList.point == todo.point && todoInList.risk == todo.risk {
			return true
		}
	}
	return false
}

func createFullRiskMap(risks map[Point]int, width int, height int) (fullRisks map[Point]int, fullWidth int, fullHeight int) {
	fullRiskMap := make(map[Point]int)

	for i := 0; i <= 4; i++ {
		for j := 0; j <= 4; j++ {
			for x := 0; x < width; x++ {
				for y := 0; y < height; y++ {
					point := Point{(x + (i * width)), (y + (j * height))}
					risk := risks[Point{x, y}] + i + j
					if risk > 9 {
						risk = risk - 9
					}
					fullRiskMap[point] = risk
					if point.x > fullWidth {
						fullWidth = point.x + 1
					}
					if point.y > fullHeight {
						fullHeight = point.y + 1
					}
				}
			}
		}
	}

	return fullRiskMap, fullWidth, fullHeight
}
