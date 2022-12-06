package days

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func (d Days) Day4p1(source io.Reader) int {
	fullOverlap := func(s string) bool {
		getRange := func(r string) (int, int) {
			split := strings.Split(r, "-")
			start, err := strconv.Atoi(split[0])
			if err != nil {
				panic("something went wrong")
			}
			end, err := strconv.Atoi(split[1])
			if err != nil {
				panic("something went wrong")
			}
			return start, end
		}
		split := strings.Split(s, ",")
		r1, r2 := split[0], split[1]

		startR1, endR1 := getRange(r1)
		startR2, endR2 := getRange(r2)

		if startR1 >= startR2 && endR1 <= endR2 {
			// (startR1, endR1) is fully covered by (startR2, endR2)
			return true
		} else if startR2 >= startR1 && endR2 <= endR1 {
			// (startR2, endR2) is fully covered by (startR1, endR1)
			return true
		}
		return false
	}

	count := 0
	fs := bufio.NewScanner(source)
	for fs.Scan() {
		if fullOverlap(fs.Text()) {
			count += 1
		}
	}
	fmt.Printf("Count (part 1): %d\n", count)
	return count
}

func (d Days) Day4p2(source io.Reader) int {
	partialOverlap := func(s string) bool {
		getRange := func(r string) (int, int) {
			split := strings.Split(r, "-")
			start, err := strconv.Atoi(split[0])
			if err != nil {
				panic("something went wrong")
			}
			end, err := strconv.Atoi(split[1])
			if err != nil {
				panic("something went wrong")
			}
			return start, end
		}
		split := strings.Split(s, ",")
		r1, r2 := split[0], split[1]

		startR1, endR1 := getRange(r1)
		startR2, endR2 := getRange(r2)

		if endR1 >= startR2 && endR1 <= endR2 {
			// (startR1, endR1) is partially covered by (startR2, endR2)
			return true
		} else if endR2 >= startR1 && endR2 <= endR1 {
			// (startR2, endR2) is partially covered by (startR1, endR1)
			return true
		}
		return false
	}

	count := 0
	fs := bufio.NewScanner(source)
	for fs.Scan() {
		if partialOverlap(fs.Text()) {
			count += 1
		}
	}
	fmt.Printf("Count (part 2): %d\n", count)
	return count
}
