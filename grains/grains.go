package grains

import "errors"

func Square(n int) (uint64, error) {
	SquareTotal := uint64(1)
	if n > 64 || n <= 0 {
		return uint64(0), errors.New("false")
	}
	for i := 2; i <= n; i++ {
		SquareTotal *= 2
	}
	return SquareTotal, nil
}

func Total() uint64 {
	sum := uint64(0)
	for i := 1; i <= 64; i++ {
		tmp, _ := Square(i)
		sum += tmp
	}

	return sum
}
