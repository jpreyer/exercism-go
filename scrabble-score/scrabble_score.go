package scrabble

import "strings"

var values = map[rune]int{
	'A': 1,
	'E': 1,
	'I': 1,
	'O': 1,
	'U': 1,
	'L': 1,
	'N': 1,
	'R': 1,
	'S': 1,
	'T': 1,
	'D': 2,
	'G': 2,
	'B': 3,
	'C': 3,
	'M': 3,
	'P': 3,
	'F': 4,
	'H': 4,
	'V': 4,
	'W': 4,
	'Y': 4,
	'K': 5,
	'J': 8,
	'X': 8,
	'Q': 10,
	'Z': 10,
}

//Score takes a word as input and returns the Scrabble score for the word.
func Score(w string) int {
	var wordScore int

	//convert word to upper case and Iterate over letters in work
	w = strings.ToUpper(w)
	//Convert w to an array of runes
	r := []rune(w)
	//Iterate over rune array to compile score
	for _, v := range r {
		wordScore += values[v]
	}

	return wordScore
}
