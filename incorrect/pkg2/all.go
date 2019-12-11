package pkg2

import (
	"github.com/dhermes/coverage-ordering/incorrect/pkg1"
)

// Fibonacci does what it says on the tin
func Fibonacci(n int) int {
	if n < 2 {
		return 1
	}
	return pkg1.Add(Fibonacci(n-1), Fibonacci(n-2))
}
