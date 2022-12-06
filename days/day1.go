package days

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Days struct{}

func (d Days) Day1p1(source io.Reader) int {

	calories := map[int]int{}
	var e1, e2, e3 int

	e := 1
	fs := bufio.NewScanner(source)
	for fs.Scan() {
		line := fs.Text()
		if line != "" {
			w, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			calories[e] += w
		} else {
			switch true {
			case calories[e] > e1:
				e1 = calories[e]
			case calories[e] > e2:
				e2 = calories[e]
			case calories[e] > e3:
				e3 = calories[e]
			}
			e += 1
		}
	}

	fmt.Printf("(Part 1) Top elf (sum of calories): %d\n", e1)
	return e1
}

func (d Days) Day1p2(source io.Reader) int {

	calories := map[int]int{}
	var e1, e2, e3 int

	e := 1
	fs := bufio.NewScanner(source)
	for fs.Scan() {
		line := fs.Text()
		if line != "" {
			w, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			calories[e] += w
		} else {
			switch true {
			case calories[e] > e1:
				e1 = calories[e]
			case calories[e] > e2:
				e2 = calories[e]
			case calories[e] > e3:
				e3 = calories[e]
			}
			e += 1
		}
	}

	fmt.Printf("(Part 2) Top 3 elves (sum of calories): %d\n", e1+e2+e3)
	return e1 + e2 + e3
}
