package isogram

import (
	"strings"
)

//function will determine if input string is an isogram and return true if it is and false if not
func IsIsogram(w string) bool {

	track_usage := make(map[rune]bool)

	// Let's remove spaces and dashes, we don't care about multiple occurances of them
	w = strings.ToLower(w)
	w = strings.ReplaceAll(w, "-", "")
	w = strings.ReplaceAll(w, " ", "")

	// Clever way to test if a key is in a map, review!!!
	//From https://www.golangprograms.com/how-to-check-if-a-map-contains-a-key-in-go.html

	for _, v := range w {
		if _, ok := track_usage[v]; ok {
			return false
		} else {
			track_usage[v] = true
		}
	}

	return true

}
