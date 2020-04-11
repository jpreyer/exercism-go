package raindrops

import "strconv"

// Convert takes an integer as input and exports a string as detailed in the Raindrops Exercism assignment
func Convert(a int) string {
	var sound string

	if a%3 == 0 {
		sound += "Pling"
	}
	if a%5 == 0 {
		sound += "Plang"
	}
	if a%7 == 0 {
		sound += "Plong"
	}
	if sound != "" {
		return sound
	}
	return strconv.Itoa(a)
}
