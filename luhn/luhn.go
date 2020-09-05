package luhn

import (
	"strings"
	"unicode"
)

//Valid takes a string as input and returns a boolean value in determining if the string is valid per the Luhn formula
//http://en.wikipedia.org/wiki/Luhn_algorithm
func Valid(s string) bool {

	total := 0

	s = strings.ReplaceAll(s, " ", "")

	if len(s) < 2 {
		return false
	}

	//
	//if len(s) is odd, start at 1m so SecondDigit is false
	//if len(s) is even, start at zero so SecondDigit is true
	SecondDigit := len(s)%2 == 0

	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}

		digit := int(r - '0')

		if SecondDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		total += digit
		SecondDigit = !SecondDigit
	}

	return total%10 == 0
}
