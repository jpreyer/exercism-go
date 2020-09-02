package luhn

import (
	"strings"
	"unicode"
)

//Slice_Sum takes a slice of int  as input and returns the sum of the slice
func Slice_Sum(slice []int) int {
	total := 0
	for _, v := range slice {
		total += v
	}
	return total
}

//Compute_Digit_Values takes a slice of int at input and returns a slice of integers that have been
//processed by the Lutn algorithm.
//The first step of the Luhn algorithm is to double every second digit, starting from the right.
//If doubling the number results in a number greater than 9 then subtract 9 from the product.
//Note: The digits we are computiung the checksum for much already be reversed!
func Compute_Digit_Values(digits []int) []int {
	i := len(digits) % 2
	for i <= len(digits)-1 {
		digits[i] *= 2
		if digits[i] > 9 {
			digits[i] -= 9
		}
		i += 2
	}
	return digits
}

//Valid takes a string as input and returns a boolean value in determining if the string is valid per the Luhn formula
//http://en.wikipedia.org/wiki/Luhn_algorithm
func Valid(s string) bool {
	//var rev_str []string
	x := []int{}

	s = strings.ReplaceAll(s, " ", "")

	if len(s) < 2 {
		return false
	}

	for _, v := range s {
		// check to see if we have a digit [0-9]
		if unicode.IsDigit(v) {
			x = append(x, int(v-'0'))
		} else {
			return false
		}
	}
	x = Compute_Digit_Values(x)
	sum := Slice_Sum(x)

	if sum%10 == 0 {
		return true
	} else {
		return false
	}

}
