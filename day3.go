package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (d days) Day3() {

	file, err := os.Open("sources/day3")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)

	// Part 1
	findBadItems := func(p1, p2 []string) []string {
		contains := func(ss []string, s string) bool {
			for _, i := range ss {
				if i == s {
					return true
				}
			}
			return false
		}
		res := []string{}
		for _, i1 := range p1 {
			for _, i2 := range p2 {
				if i1 == i2 {
					if !contains(res, i1) {
						res = append(res, i1)
					}

				}
			}
		}
		return res
	}

	findScore := func(s string) int {
		r := []rune(s)
		if len(r) > 1 {
			panic("something went wrong")
		}
		if r[0] >= 97 && r[0] <= 122 {
			return int(r[0]) - 97 + 1
		} else if r[0] >= 65 && r[0] <= 90 {
			return int(r[0]) - 65 + 27
		}
		return 0
	}

	score := 0
	for fs.Scan() {
		line := fs.Text()
		split := strings.Split(line, "")
		part1, part2 := split[0:len(split)/2], split[len(split)/2:]
		badItems := findBadItems(part1, part2)
		for _, item := range badItems {
			score += findScore(item)
		}
	}
	fmt.Printf("Score (part 1): %d\n", score)

	// Part 2
	file, err = os.Open("sources/day3")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fs = bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)

	getGroupBadge := func(group []string) string {
		badges := map[string]int{}
		for _, g := range group {
			split := strings.Split(g, "")
			uniqueItems := map[string]bool{}
			for _, s := range split {
				_, ok := uniqueItems[s]
				if !ok {
					badges[s] += 1
				}
				uniqueItems[s] = true
			}
		}
		for badge, count := range badges {
			if count == 3 {
				return badge
			}
		}
		return ""
	}

	score = 0
	group := make([]string, 0, 3)
	for fs.Scan() {
		group = append(group, fs.Text())
		if len(group) == 3 {
			badge := getGroupBadge(group)
			score += findScore(badge)
			group = make([]string, 0, 3)
		}
	}
	fmt.Printf("Score (part 2): %d\n", score)
}
