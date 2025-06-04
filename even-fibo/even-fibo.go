package evenfibo

//https://projecteuler.net/problem=2

func EvenFibonacci(limit int) int {
	if limit < 2 {
		return 0
	}

	a, b := 1, 2
	sum := 0

	for b <= limit {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}

	return sum
}
