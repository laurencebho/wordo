package main

import (
	"log"
	"math"
	"os"
	"strings"
)

type Candidate struct {
	word     string
	patterns [3][3][3][3][3]float64
	score    float64
}

func getWords() []string {
	data, err := os.ReadFile("words.txt")

	if err != nil {
		log.Fatal(err)
	}

	data_str := string(data)
	words := strings.Fields(data_str)
	return words
}

func checkMatch(candidate_word string, target_word string) [5]int {
	var ret [5]int
	for i, c := range candidate_word {
		if string(c) == string(target_word[i]) {
			ret[i] = 2
		} else if strings.Contains(target_word, string(c)) {
			ret[i] = 1
		} else {
			ret[i] = 0
		}
	}
	return ret
}

//calculate the score (uncertainty) for a candidate word
func calcScore(candidate *Candidate) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					for m := 0; m < 3; m++ {
						var prob float64 = candidate.patterns[i][j][k][l][m] / 12972
						if prob > 0 {
							candidate.score -= prob * math.Log2(prob)
						}
					}
				}
			}
		}
	}

}
