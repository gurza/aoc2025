package day10

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		line     string
		lights   int
		buttons  []int
		joltages []int
	}{
		{
			line:     "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
			lights:   0b0110,
			buttons:  []int{8, 10, 4, 12, 5, 3},
			joltages: []int{3, 5, 4, 7},
		},
		{
			line:     "[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
			lights:   0b01000,
			buttons:  []int{29, 12, 17, 7, 30},
			joltages: []int{7, 5, 12, 7, 2},
		},
		{
			line:     "[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
			lights:   0b101110,
			buttons:  []int{31, 25, 55, 6},
			joltages: []int{10, 11, 11, 5, 10, 5},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := parse(tt.line)

			assert.Equal(t, tt.lights, got.lights, "lights mismatch")
			assert.Equal(t, tt.buttons, got.buttons, "buttons mismatch")
			assert.Equal(t, tt.joltages, got.joltages, "joltages mismatch")
		})
	}
}
