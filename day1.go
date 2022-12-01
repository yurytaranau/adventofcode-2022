package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func (d days) Day1() {

	readFile, err := os.Open("sources/day1")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fs := bufio.NewScanner(readFile)
	fs.Split(bufio.ScanLines)

	calories := map[int]int{}
	var e1, e2, e3 int

	e := 1
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

	fmt.Printf("Top elf (sum of calories): %d, Top 3 elves (sum of calories): %d\n", e1, e1+e2+e3)
}
