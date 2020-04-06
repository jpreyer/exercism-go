package hamming

import (
	"errors"
)

// Distance will Calculate hamming distance
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return 0, errors.New("false")
	}
	code := 0
	for i := 0; i < len(ar); i++ {
		if ar[i] != br[i] {
			code++
		}
	}
	return code, nil
}
