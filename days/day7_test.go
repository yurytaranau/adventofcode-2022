package days

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay7_Part1(t *testing.T) {
	d := Days{}
	input := `
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`
	res := d.Day7p1(strings.NewReader(input))
	assert.Equal(t, 95437, res, input)
}

func TestDay7_Part2(t *testing.T) {
	d := Days{}
	input := `
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`
	res := d.Day7p2(strings.NewReader(input))
	assert.Equal(t, 24933642, res, input)
}
