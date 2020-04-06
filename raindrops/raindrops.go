package raindrops

import "strconv"

// Convert takes an integer as input and exports a string as detailed in the Raindrops Exercism assignment
func Convert(a int) string {
	var sound string
	var divisibleBy357 bool = false
	if a%3 == 0 {
		sound = sound + "Pling"
		divisibleBy357 = true
	}
	if a%5 == 0 {
		sound = sound + "Plang"
		divisibleBy357 = true
	}
	if a%7 == 0 {
		sound = sound + "Plong"
		divisibleBy357 = true
	}
	if divisibleBy357 {
		return sound
	}
	return strconv.Itoa(a)
}
