package repository

import (
	"os"
	"technical_test/helpers"
)

func PrepareFile(input string) [][]int {
	if dat, err := os.ReadFile(input); err != nil {
		panic(err)
	} else {
		file := string(dat)
		tabElement := helpers.Parse(file)

		return tabElement
	}
}
