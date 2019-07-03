package quick_sort_test_test

import (
	"testing"
)

func Loop(n uint64) (result uint64) {
	result = 1
	var i uint64 = 1
	for ; i < n; i++ {
		result *= i
	}
	return result
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func TestLoop(t *testing.T) {
	t.Log("Loop:", Loop(32))
}

func TestFactorial(t *testing.T) {
	t.Log("Factorial:", Factorial(32))
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop(40)
	}
}




