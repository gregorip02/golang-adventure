package sum

import (
	"testing"
)

type SumSuite struct {
	a int
	b int
	r int
}

func TestSum(t *testing.T) {
	tests := []SumSuite {
		{ 10, 10, 20 },
		{ 34, -1, 33 },
		{ -1, -1, -2 },
		{ -1, -10, -11 },
	}

	for _, test := range tests {
		result := Sum(test.a, test.b)
		if result != test.r {
			t.Fatalf("Unexpected error, got %d expected %d", result, test.r)
		}
	}
}
