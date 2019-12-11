package pkg1_test

import (
	"testing"

	"github.com/dhermes/coverage-ordering/incorrect/pkg1"
)

func TestMultiply(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{0, 1, 0},
		{5, -4, -20},
		{8, 10, 80},
		{-7, -7, 49},
	}

	for _, tc := range tests {
		got := pkg1.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
