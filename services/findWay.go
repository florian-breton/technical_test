package services

import (
	"fmt"
	"technical_test/models"
	"technical_test/repository"
)

var maxCoordonates models.Coordonates

func FindWay(input string) int {
	matrix := repository.PrepareFile(input)
	maxCoordonates = models.Coordonates{len(matrix), len(matrix[0])}

	returnValue := proceed(matrix)
	return returnValue
}

func proceed(matrix [][]int) int {
	allPathPossible := getAllPathPossible(matrix)
	allUniquePath := getuniquePath(allPathPossible)
	fmt.Println(allUniquePath)

	return len(allPathPossible)
}

// remove duplicate path
func getuniquePath(allPathPossible [][2]models.Coordonates) [][2]models.Coordonates {
	var uniquePath [][2]models.Coordonates
	for _, path := range allPathPossible {
		if !isIn(uniquePath, path) {
			uniquePath = append(uniquePath, path)
		}
	}
	return uniquePath
}

// check if path is already in the list
func isIn(uniquePath [][2]models.Coordonates, path [2]models.Coordonates) bool {
	for _, unique := range uniquePath {
		if unique[0] == path[0] && unique[1] == path[1] {
			return true
		}
	}
	return false
}

// get all path possible
func getAllPathPossible(matrix [][]int) [][2]models.Coordonates {
	var pathPossible [][2]models.Coordonates
	for x, line := range matrix {
		for y, cell := range line {
			if cell == 0 {
				initialPosition := models.Coordonates{x, y}
				go findNextStep(matrix, initialPosition, initialPosition, initialPosition, pathPossible)
			}
		}
	}

	return pathPossible
}

// FindNextStep find next step based on the current position, previous position, initial position and the matrix
// return all path possible
func findNextStep(matrix [][]int, position models.Coordonates, previousPosition models.Coordonates, initialPosition models.Coordonates, pathPossible [][2]models.Coordonates) [][2]models.Coordonates {

	if position.x < 0 || position.y < 0 || position.x >= maxCoordonates.x || position.y >= maxCoordonates.y {
		return pathPossible
	}
	if matrix[position.x][position.y] == 9 && matrix[previousPosition.x][previousPosition.y] == 8 {
		path := [2]models.Coordonates{initialPosition, position}
		pathPossible = append(pathPossible, path)
		return pathPossible
	}
	if position == previousPosition || matrix[position.x][position.y] == matrix[previousPosition.x][previousPosition.y]+1 {
		pathPossible = findNextStep(matrix, models.Coordonates{position.X + 1, position.Y}, position, initialPosition, pathPossible)
		pathPossible = findNextStep(matrix, models.Coordonates{position.X - 1, y: position.y}, position, initialPosition, pathPossible)
		pathPossible = findNextStep(matrix, models.Coordonates{position.X, position.Y + 1}, position, initialPosition, pathPossible)
		pathPossible = findNextStep(matrix, models.Coordonates{position.Y, position.Y - 1}, position, initialPosition, pathPossible)
	}

	return pathPossible
}
