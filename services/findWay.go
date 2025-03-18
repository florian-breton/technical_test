package services

import "technical_test/helpers"

type Coordinates struct {
	x int
	y int
}

func FindWay(input string) (int, error) {
	matrix, err := helpers.PrepareFile(input)
	if err != nil {
		return 0, err
	}

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0, nil
	}
	maxCoords := Coordinates{x: len(matrix), y: len(matrix[0])}

	totalScore := 0
	for x, row := range matrix {
		for y, cell := range row {
			if cell == 0 { // Trailhead
				score := calculateTrailheadScore(matrix, Coordinates{x, y}, maxCoords)
				totalScore += score
			}
		}
	}
	return totalScore, nil
}

func calculateTrailheadScore(matrix [][]int, start Coordinates, maxCoords Coordinates) int {
	visited := make(map[Coordinates]struct{})
	reachableNines := make(map[Coordinates]struct{})
	findReachableNines(matrix, start, -1, visited, reachableNines, maxCoords)
	return len(reachableNines)
}

func findReachableNines(matrix [][]int, pos Coordinates, prevVal int, visited, reachableNines map[Coordinates]struct{}, maxCoords Coordinates) {
	// Boundary check
	if pos.x < 0 || pos.y < 0 || pos.x >= maxCoords.x || pos.y >= maxCoords.y {
		return
	}

	currentVal := matrix[pos.x][pos.y]

	// If we've visited this position with a valid sequence, no need to revisit
	if _, ok := visited[pos]; ok {
		return
	}

	// Check if this is a valid step: either the start (prevVal = -1) or increment by 1
	if prevVal == -1 && currentVal != 0 { // Must start at 0
		return
	}
	if prevVal != -1 && currentVal != prevVal+1 { // Must increase by exactly 1
		return
	}

	visited[pos] = struct{}{}

	// If we reach a 9, record it
	if currentVal == 9 {
		reachableNines[pos] = struct{}{}
	}

	directions := []Coordinates{
		{x: pos.x + 1, y: pos.y},
		{x: pos.x - 1, y: pos.y},
		{x: pos.x, y: pos.y + 1},
		{x: pos.x, y: pos.y - 1},
	}

	for _, next := range directions {
		findReachableNines(matrix, next, currentVal, visited, reachableNines, maxCoords)
	}
}
