package luhn

import (
	"strings"
	"unicode"
)

//ComputeDigitValue takes a slice of int at input and returns a slice of integers that have been
//processed by the Lutn algorithm.
//The first step of the Luhn algorithm is to double every second digit, starting from the right.
//If doubling the number results in a number greater than 9 then subtract 9 from the product.
//Note: The digits we are computiung the checksum for much already be reversed!
//We then sum up the digit values and return the sum
func ComputeDigitValue(digit int) int {
	digit *= 2
	if digit > 9 {
		digit -= 9
	}
	return digit
}

//Valid takes a string as input and returns a boolean value in determining if the string is valid per the Luhn formula
//http://en.wikipedia.org/wiki/Luhn_algorithm
func Valid(s string) bool {
	total := 0
	s = strings.ReplaceAll(s, " ", "")

	if len(s) < 2 {
		return false
	}

	//Is the digit we're working on one that should be given to ComputeDigitValue?  At the start, no it is not
	SecondDigit := false
	//Convert input string to an array of Runes
	RuneString := []rune(s)

	for j := len(s) - 1; j >= 0; j-- {

		if !unicode.IsDigit(RuneString[j]) {
			return false
		}

		d := int(RuneString[j] - '0')

		if SecondDigit {
			total += ComputeDigitValue(d)
		} else {
			total += d
		}
		//moving to next digit, so flip if this is now SecondDigit
		SecondDigit = !SecondDigit

	}

	return total%10 == 0
}
