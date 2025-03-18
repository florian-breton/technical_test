package helpers

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Explode(text, delimiter string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
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

func Parse(input string) [][]int {
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

	return matrix
}
