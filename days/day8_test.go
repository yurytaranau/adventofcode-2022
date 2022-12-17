package days

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay8_Part1(t *testing.T) {
	d := Days{}
	input := `
30373
25512
65332
33549
35390
`
	res := d.Day8p1(strings.NewReader(input))
	assert.Equal(t, 21, res, input)
}

func TestDay8_Part2(t *testing.T) {
	d := Days{}
	input := `
30373
25512
65332
33549
35390
`
	res := d.Day8p2(strings.NewReader(input))
	assert.Equal(t, 8, res, input)
}
