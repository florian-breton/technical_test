package technicalTest

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
		{"Day 10 Part 1", args{"InputTest.txt"}, 81},
		{"Day 10 Part 1", args{"InputTest2.txt"}, 2},
		{"Day 10 Part 2", args{"Input.txt"}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindWay(tt.args.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
