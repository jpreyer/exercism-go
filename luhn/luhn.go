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
	l := len(s)

	if l < 2 {
		return false
	}

	for index, v := range s {
		if unicode.IsDigit(v) {
			d := int(v - '0')
			if (l%2 == 1 && index%2 == 1) || (l%2 == 0 && index%2 == 0) {
				total += ComputeDigitValue(d)
			} else {
				total += d
			}
		} else {
			return false
		}
	}

	return total%10 == 0
}
