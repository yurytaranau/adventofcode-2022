package days

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

var shapeScores = map[string]int{
	"A": 1, // Rock
	"B": 2, // Paper
	"C": 3, // Scissors
}
var shapeMap = map[string]string{
	"X": "A",
	"Y": "B",
	"Z": "C",
}
var roundScores = map[string]int{
	"AA": 3, // Rock vs. Rock
	"AB": 6, // Rock vs. Paper
	"AC": 0, // Rock vs. Scissors
	"BA": 0, // Paper vs. Rock
	"BB": 3, // Paper vs. Paper
	"BC": 6, // Paper vs. Scissors
	"CA": 6, // Scissors vs. Rock
	"CB": 0, // Scissors vs. Paper
	"CC": 3, // Scissors vs. Scissors
}

func (d Days) Day2p1(source io.Reader) int {
	// Part 1
	buf := &bytes.Buffer{}
	tee := io.TeeReader(source, buf)

	score := 0
	fs := bufio.NewScanner(tee)
	for fs.Scan() {
		round := fs.Text()
		split := strings.Split(round, " ")
		opponent, shape := split[0], split[1]
		score += shapeScores[shapeMap[shape]] + roundScores[opponent+shapeMap[shape]]
	}
	fmt.Printf("Total score (first strategy): %d\n", score)
	return score
}

func (d Days) Day2p2(source io.Reader) int {
	newStrategy := map[string]int{
		"X": 0, // Need to lose
		"Y": 3, // Draw
		"Z": 6, // Need to win
	}

	myResponse := func(opponent string, strategy int) string {
		for k, v := range roundScores {
			split := strings.Split(k, "")
			x, y := split[0], split[1]
			if x == opponent && v == strategy {
				return y
			}
		}
		return ""
	}

	score := 0
	fs := bufio.NewScanner(source)
	for fs.Scan() {
		round := fs.Text()
		split := strings.Split(round, " ")
		opponent, strategy := split[0], split[1]
		myShape := myResponse(opponent, newStrategy[strategy])
		if myShape == "" {
			panic("something went bad")
		}
		score += shapeScores[myShape] + roundScores[opponent+myShape]
	}
	fmt.Printf("Total score (new strategy): %d\n", score)
	return score
}
