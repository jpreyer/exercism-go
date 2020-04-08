package collatzconjecture

import "errors"

//CollatzConjecture takes an integer as an input, runs in through the Collatz Conjecture algorithm, and returns the number of iterations to get to 1 as well as an error message.
func CollatzConjecture(n int) (int, error) {
	var iteration int = 0

	if n <= 0 {
		return 0, errors.New("false")
	}
	for n != 1 {
		switch n % 2 {
		case 0:
			n = n / 2
		case 1:
			n = 3*n + 1
		}
		iteration++
	}
	return iteration, nil
}
