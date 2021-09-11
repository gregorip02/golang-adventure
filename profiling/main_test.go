package profiling

import "testing"

type FibonacciSuite struct {
	n int
	r int
}

var suite []FibonacciSuite = []FibonacciSuite{
	{n: 8, r: 21},
	{n: 2, r: 1},
	{n: 0, r: 0},
	{n: 20, r: 6765},
	{n: 1, r: 1},
}

func TestFibonacciFastest(t *testing.T) {
	for _, s := range suite {
		r := FibonacciFastest(s.n)
		if r != s.r {
			t.Fatalf("Unexpected result, got %d expected %d", r, s.r)
		}
	}
}

func TestFibonacciSlowest(t *testing.T) {
	for _, s := range suite {
		r := FibonacciSlowest(s.n)
		if r != s.r {
			t.Fatalf("Unexpected result, got %d expected %d", r, s.r)
		}
	}
}
