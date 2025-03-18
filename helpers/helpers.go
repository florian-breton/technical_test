package helpers

import (
	"os"
	"strconv"
	"strings"
)

func parse(input string) ([][]int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return nil, nil
	}
	var matrix [][]int
	for _, line := range lines {
		var row []int
		for _, char := range strings.TrimSpace(line) {
			n, err := strconv.Atoi(string(char))
			if err != nil {
				row = append(row, -1)
			} else {
				row = append(row, n)
			}
		}
		if len(row) > 0 {
			matrix = append(matrix, row)
		}
	}
	return matrix, nil
}

func PrepareFile(input string) ([][]int, error) {
	dat, err := os.ReadFile(input)
	if err != nil {
		return nil, err
	}
	return parse(string(dat))
}
