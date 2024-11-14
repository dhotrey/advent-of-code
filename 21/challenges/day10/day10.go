package day10

import (
	"21/utils"
	"slices"

	"github.com/charmbracelet/log"
)

var brackets = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
	'>': '<',
}

var illegalCharacterScoreTable = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var autocompletePointsTable = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func Sol(mode string) {
	data, file := utils.GetInput(10, mode)
	defer file.Close()
	var totalCorruptionScore int
	autocompleteScores := []int{}
	for data.Scan() {
		line := data.Text()
		lineCorruptionScore := getSyntaxErrorScore(line)
		totalCorruptionScore += lineCorruptionScore

		if lineCorruptionScore == 0 {
			autocompleteScores = append(autocompleteScores, getAutocompleteScore(line))
		}
	}
	log.Infof("Total syntax error score is : %d", totalCorruptionScore)
	slices.Sort(autocompleteScores)
	log.Debug(autocompleteScores)
	log.Infof("Autocomplete comp winner is %d", autocompleteScores[len(autocompleteScores)/2])
}

func getAutocompleteScore(line string) int {
	var s stack

	for _, char := range line {
		switch char {
		case '(', '{', '[', '<':
			s.push(char)
		case ')', '}', ']', '>':
			_, err := s.pop()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	incomplete := string(s)
	var score int
	for i := len(incomplete) - 1; i >= 0; i-- {
		char := incomplete[i]
		score *= 5
		score += autocompletePointsTable[rune(char)]
	}
	log.Debug(incomplete, "score", score)
	return score
}

func getSyntaxErrorScore(line string) int {
	var score int
	var s stack
	log.Debugf("Calculating score for : %s", line)
	for _, char := range line {
		switch char {
		case '(', '{', '[', '<':
			s.push(char)
		case ')', '}', ']', '>':
			closingBracket := char
			openingBracket, err := s.pop()
			if err != nil {
				log.Fatal(err)
			}
			if openingBracket != brackets[closingBracket] {
				points := illegalCharacterScoreTable[closingBracket]
				log.Debugf("incorrect bracket sequence %s%s : %d points", string(openingBracket), string(closingBracket), points)
				score += points
			}
		}
	}
	return score
}
