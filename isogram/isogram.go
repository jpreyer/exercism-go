package isogram

import (
	"regexp"
	"strings"
)

//function will determine if input string is an isogram and return true if it is and false if not
func IsIsogram(w string) bool {

	var status bool
	var track_usage = map[rune]int{}

	// Let's remove spaces and dashes, we don't care about multiple occurances of them
	w = strings.ToLower(w)
	re := regexp.MustCompile(`[-]`)
	w = re.ReplaceAllString(w, "")
	re1 := regexp.MustCompile(`[ ]+`)
	w = re1.ReplaceAllString(w, "")

	status = true
	// Clever way to test if a key is in a map, review!!!
	//From https://www.golangprograms.com/how-to-check-if-a-map-contains-a-key-in-go.html
	for _, v := range w {
		_, ok := track_usage[v]
		if !ok {
			track_usage[v] += 1
		} else {
			status = false
		}
	}

	return status

}
