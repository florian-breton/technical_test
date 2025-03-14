package technicalTest

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordonates struct {
	x int
	y int
}

var maxCoordonates Coordonates

func FindWay(input string) int {
	matrix := prepareFile(input)

	returnValue := proceed(matrix)
	return returnValue
}

func prepareFile(input string) [][]int {
	dat, err := os.ReadFile(input)
	check(err)
	file := string(dat)
	tabElement := parse(file)
	return tabElement
}

func proceed(matrix [][]int) int {
	allPathPossible := getAllPathPossible(matrix)
	allUniquePath := getuniquePath(allPathPossible)
	fmt.Println(allUniquePath)

	return len(allPathPossible)
}

// remove duplicate path
func getuniquePath(allPathPossible [][2]Coordonates) [][2]Coordonates {
	var uniquePath [][2]Coordonates
	for _, path := range allPathPossible {
		if !isIn(uniquePath, path) {
			uniquePath = append(uniquePath, path)
		}
	}
	return uniquePath
}

// check if path is already in the list
func isIn(uniquePath [][2]Coordonates, path [2]Coordonates) bool {
	for _, unique := range uniquePath {
		if unique[0] == path[0] && unique[1] == path[1] {
			return true
		}
	}
	return false
}

// get all path possible
func getAllPathPossible(matrix [][]int) [][2]Coordonates {
	var pathPossible [][2]Coordonates
	for x, line := range matrix {
		for y, cell := range line {
			if cell == 0 {
				initialPosition := Coordonates{x, y}
				pathPossible = FindNextStep(matrix, initialPosition, initialPosition, initialPosition, pathPossible)
			}
		}
	}

	return pathPossible
}

// FindNextStep find next step based on the current position, previous position, initial position and the matrix
// return all path possible
func FindNextStep(matrix [][]int, position Coordonates, previousPosition Coordonates, initialPosition Coordonates, pathPossible [][2]Coordonates) [][2]Coordonates {

	if position.x < 0 || position.y < 0 || position.x >= maxCoordonates.x || position.y >= maxCoordonates.y {
		return pathPossible
	}
	if matrix[position.x][position.y] == 9 && matrix[previousPosition.x][previousPosition.y] == 8 {
		path := [2]Coordonates{initialPosition, position}
		pathPossible = append(pathPossible, path)
		return pathPossible
	}
	if position == previousPosition || matrix[position.x][position.y] == matrix[previousPosition.x][previousPosition.y]+1 {
		pathPossible = FindNextStep(matrix, Coordonates{x: position.x + 1, y: position.y}, position, initialPosition, pathPossible)
		pathPossible = FindNextStep(matrix, Coordonates{x: position.x - 1, y: position.y}, position, initialPosition, pathPossible)
		pathPossible = FindNextStep(matrix, Coordonates{x: position.x, y: position.y + 1}, position, initialPosition, pathPossible)
		pathPossible = FindNextStep(matrix, Coordonates{x: position.x, y: position.y - 1}, position, initialPosition, pathPossible)
	}
	return pathPossible
}

func parse(input string) [][]int {
	line := Explode(input, "\n")
	var matrix [][]int

	for _, item := range line {
		lineTab := Explode(item, "")
		var matrixLine []int
		for _, cell := range lineTab {
			matrixLine = append(matrixLine, StringToInt(cell))
		}
		matrix = append(matrix, matrixLine)
	}
	maxCoordonates = Coordonates{len(matrix), len(matrix[0])}
	return matrix
}

func Explode(text, delimiter string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func StringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return n
}

func main() {
	//FindWay("FindWay/InputTest.txt")
	FindWay("InputTest2.txt")
}
