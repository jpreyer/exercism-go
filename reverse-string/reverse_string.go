package reverse

//import "unicode/utf8"

//Reverse takes a string as input and returns the strings reversed as output
func Reverse(s string) string {

	b := []rune(s)
	RevString := make([]rune, len(b), (len(b)))

	for index, r := range b {
		RevString[len(b)-index-1] = r
	}

	return string(RevString)
}
