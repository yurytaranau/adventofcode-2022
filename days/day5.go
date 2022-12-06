package days

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Reverse() stack {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

var parseCrates = func(line string, stacks []stack) []stack {
	parseCrateName := func(s []string) string {
		r := []rune(s[1])
		if r[0] >= 65 && r[0] <= 90 {
			return string(r[0])
		}
		return ""
	}

	crate := make([]string, 0, 3)
	i := 0
	crateCompleted := false
	for _, r := range strings.Split(line, "") {
		if crateCompleted {
			//skip empty spaces beween crates
			crateCompleted = false
			continue
		}
		crate = append(crate, r)
		if len(crate) == 3 {
			// analyze crate name
			crateName := parseCrateName(crate)
			if crateName != "" {
				stacks[i] = stacks[i].Push(crateName)
			}
			i += 1
			crate = make([]string, 0, 3)
			crateCompleted = true
		}

	}
	return stacks
}

func (d Days) Day5p1(source io.Reader) string {

	moveCrate := func(stacks []stack, from, to int) []stack {
		stack, item := stacks[from].Pop()
		stacks[from] = stack
		stacks[to] = stacks[to].Push(item)
		return stacks
	}

	stacks := make([]stack, 9)
	stacksCompleted := false

	fs := bufio.NewScanner(source)
	for fs.Scan() {
		// parse stacks of crates
		if fs.Text() != "" && !stacksCompleted {
			stacks = parseCrates(fs.Text(), stacks)
			continue
		} else if fs.Text() == "" {
			stacksCompleted = true
			// invert stacks
			for i, s := range stacks {
				stacks[i] = s.Reverse()
			}
		}
		// parse and execute movements
		re := regexp.MustCompile("[0-9]+")
		command := re.FindAllString(fs.Text(), -1)
		if len(command) > 0 {
			iterate, err := strconv.Atoi(command[0])
			if err != nil {
				panic("something went wrong")
			}
			for {
				if iterate == 0 {
					break
				}
				from, err := strconv.Atoi(command[1])
				if err != nil {
					panic("something went wrong")
				}
				to, err := strconv.Atoi(command[2])
				if err != nil {
					panic("something went wrong")
				}
				stacks = moveCrate(stacks, from-1, to-1)
				iterate--
			}
		}

	}
	// dump stack heads
	var heads string
	for _, s := range stacks {
		_, head := s.Pop()
		heads += head
	}
	fmt.Printf("(Part 1) Stacks head: %s\n", heads)
	return heads
}

func (d Days) Day5p2(source io.Reader) string {
	moveCrates := func(stacks []stack, count, from, to int) []stack {
		crates := stack{}
		// extract required crates
		i := count
		for {
			if i == 0 {
				break
			}
			stack, item := stacks[from].Pop()
			stacks[from] = stack
			crates = crates.Push(item)
			i--
		}
		// push to new stack
		i = count
		for {
			if i == 0 {
				break
			}
			stack, item := crates.Pop()
			crates = stack
			stacks[to] = stacks[to].Push(item)
			i--
		}
		return stacks
	}

	stacks := make([]stack, 9)
	stacksCompleted := false

	fs := bufio.NewScanner(source)
	for fs.Scan() {
		// parse stacks of crates
		if fs.Text() != "" && !stacksCompleted {
			stacks = parseCrates(fs.Text(), stacks)
			continue
		} else if fs.Text() == "" {
			stacksCompleted = true
			// invert stacks
			for i, s := range stacks {
				stacks[i] = s.Reverse()
			}
		}
		// parse and execute movements
		re := regexp.MustCompile("[0-9]+")
		command := re.FindAllString(fs.Text(), -1)
		if len(command) > 0 {
			count, err := strconv.Atoi(command[0])
			if err != nil {
				panic("something went wrong")
			}
			from, err := strconv.Atoi(command[1])
			if err != nil {
				panic("something went wrong")
			}
			to, err := strconv.Atoi(command[2])
			if err != nil {
				panic("something went wrong")
			}
			stacks = moveCrates(stacks, count, from-1, to-1)
		}

	}
	// dump stack heads
	heads := ""
	for _, s := range stacks {
		_, head := s.Pop()
		heads += head
	}
	fmt.Printf("(Part 2) Stacks head: %s\n", heads)
	return heads
}
