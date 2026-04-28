package differenceofsquares

func SquareOfSum(n int) int {
	a := 0
	for i := 0; i <= n; i++ {
		a += i
	}
	return a * a
}

func SumOfSquares(n int) int {
	a := 0
	for i := 0; i <= n; i++ {

		a += i * i
	}
	return a
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
