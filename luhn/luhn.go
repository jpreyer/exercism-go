package luhn

import (
	"strings"
	"unicode"
)

//ComputeDigitValue takes a slice of int at input and returns a slice of integers that have been
//processed by the Lutn algorithm.
//The first step of the Luhn algorithm is to double every second digit, starting from the right.
//If doubling the number results in a number greater than 9 then subtract 9 from the product.
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
	var SecondDigit bool
	total := 0

	s = strings.ReplaceAll(s, " ", "")

	if len(s) < 2 {
		return false
	}

	//
	//if len(s) is odd, start at 1
	//if len(s) is even, start at zero
	if len(s)%2 == 0 {
		SecondDigit = true
	} else {
		SecondDigit = false
	}

	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}

		d := int(r - '0')

		if SecondDigit {
			d = ComputeDigitValue(d)
		}
		total += d
		SecondDigit = !SecondDigit
	}

	return total%10 == 0
}
