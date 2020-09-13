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
	sum := uint64(0)
	for i := 1; i <= 64; i++ {
		tmp, _ := Square(i)
		sum += tmp
	}

	return sum
}
