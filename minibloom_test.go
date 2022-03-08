package minibloom

import (
	"testing"
)

func TestCalc(t *testing.T) {
	n := 100000000
	p := 0.005
	size, hashCounts := Calculate(n, p)
	t.Log(size, hashCounts)
}
