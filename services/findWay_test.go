package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWay(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{"1", args{"../tests/InputTest.txt"}, 36},
		{"2", args{"../tests/InputTest2.txt"}, 2},
		{"3", args{"../tests/Input.txt"}, 786},
		{"4", args{"../tests/inputTest3.txt"}, 566},
		{"5", args{"../tests/inputTest4.txt"}, 2},
		{"5", args{"../tests/inputTest5.txt"}, 4},
		{"5", args{"../tests/inputTest6.txt"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FindWay(tt.args.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
