// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var ans string
	re := regexp.MustCompile(`[-_,]`)
	s = re.ReplaceAllString(s, " ")
	re1 := regexp.MustCompile(`[ ]+`)
	s = re1.ReplaceAllString(s, " ")

	phrase_split := strings.Split(strings.Title(s), " ")
	for i := 0; i < len(phrase_split); i++ {
		ans = ans + string(phrase_split[i][0])
	}
	return ans
}
