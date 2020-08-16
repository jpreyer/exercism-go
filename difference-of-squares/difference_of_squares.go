package diffsquares

//SquareofSum returns the square of the sum of the first n integers.  We calculate
//the sum using the formula [n*(n+1)] / 2
//Go's Pow function uses float64, which I guess we could use, but
//this is just integer arithmetic!!!
func SquareOfSum(n int) int {
	sq_sum := ((n * (n + 1)) / 2)
	return sq_sum * sq_sum
}

//SumofSquares calculates the sum of the squares of the first n integers by using the formula
// [n * (n+1) * (2*n + 1)]/6
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

//Differnece calculates the difference of the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
