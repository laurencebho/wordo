package main

import (
	"fmt"
)

func main() {
	words := getWords()
	for _, w := range words {
		fmt.Println(w)
	}

	for _, word := range words {
		c1 := Candidate{word: word}
		var patterns [3][3][3][3][3]float64

		for _, w := range words {
			match := checkMatch(c1.word, w)
			patterns[match[0]][match[1]][match[2]][match[3]][match[4]] += 1
		}

		c1.patterns = patterns

		calcScore(&c1)

		fmt.Println(c1.word, "   ", c1.score)
	}
}
