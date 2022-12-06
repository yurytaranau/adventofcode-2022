package days

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay6_Part1(t *testing.T) {
	d := Days{}
	tests := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    7,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
		"nppdvjthqldpwncqszvftbrmjlhg":      6,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
	}
	for input, exp := range tests {
		res := d.Day6p1(strings.NewReader(input))
		assert.Equal(t, exp, res, input)
	}
}

func TestDay6_Part2(t *testing.T) {
	d := Days{}
	tests := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
		"nppdvjthqldpwncqszvftbrmjlhg":      23,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
	}
	for input, exp := range tests {
		res := d.Day6p2(strings.NewReader(input))
		assert.Equal(t, exp, res, input)
	}
}
