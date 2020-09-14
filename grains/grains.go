package grains

import "errors"

//Square takes an integer n an input and outputs the number of grains on that square: 2^(n-1)
func Square(n int) (uint64, error) {
	if n > 64 || n <= 0 {
		return uint64(0), errors.New("false")
	}

	return 1 << (n - 1), nil
}

//Total returns the total grains on a 8x8 chess board
func Total() uint64 {
	// 2 + 2^1 + 2^2 + 2^3 + ...... + 2^64 is a geometric series!
	// We can find the sum of this by calculating [1 * (1 - 2^64)] / [1 - 2] = [1 - 2^64] / -1 = 2^64 -1

	return 1<<64 - 1

}
