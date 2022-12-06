package days

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func shift(ss []string, pos int) []string {
	return ss[pos+1:]
}

func position(ss []string, s string) int {
	for i, item := range ss {
		if item == s {
			return i
		}
	}
	return -1
}

func uniqueChain(ss []string, length int) int {
	marker := []string{}
	for i, char := range ss {
		pos := position(marker, char)
		if pos >= 0 {
			marker = shift(marker, pos)
		}
		marker = append(marker, char)
		if len(marker) == length {
			return i + 1
		}
	}
	return -1
}

func (d Days) Day6p1(source io.Reader) (sop int) {
	fs := bufio.NewScanner(source)
	for fs.Scan() {
		sop = uniqueChain(strings.Split(fs.Text(), ""), 4) // start-of-packet iterations count
		fmt.Printf("(Part 1) start-of-packet marker detected after %d iterations\n", sop)
	}
	return
}

func (d Days) Day6p2(source io.Reader) (som int) {
	fs := bufio.NewScanner(source)
	for fs.Scan() {
		som = uniqueChain(strings.Split(fs.Text(), ""), 14) // start-of-message iterations
		fmt.Printf("(Part 1) start-of-message marker detected after %d iterations\n", som)
	}
	return
}
