package pkg1_test

import (
	"testing"

	"github.com/dhermes/coverage-ordering/correct/pkg1"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{7, 21},
		{8, 34},
	}

	for _, tc := range tests {
		got := pkg1.Fibonacci(tc.n)
		if tc.want != got {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
