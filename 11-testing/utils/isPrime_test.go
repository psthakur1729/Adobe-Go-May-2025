package utils

import (
	"fmt"
	"testing"
)

func TestIsPrime_13(t *testing.T) {
	t.Skip()
	// Arrange
	no := 13
	expected := true

	// Act
	actual := IsPrime(no)

	// Assert
	if actual != expected {
		t.Errorf("Expected : %v, Actual : %v\n", expected, actual)
	}
}

func TestIsPrime(t *testing.T) {
	t.Skip()
	var test_data = []struct {
		no       int
		expected bool
	}{
		{no: 13, expected: true},
		{no: 17, expected: true},
		{no: 19, expected: true},
		{no: 23, expected: true},
	}
	for _, td := range test_data {
		t.Run(fmt.Sprintf("isPrime(%d)", td.no), func(t *testing.T) {
			t.Parallel()
			// arrange
			no := td.no
			expected := td.expected

			// act
			actual := IsPrime(no)

			// assert
			if actual != expected {
				t.Errorf("Expected : %v, Actual : %v", expected, actual)
			}
		})
	}

}

// Micro benchmarking
func Benchmark_IsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(97)
	}
}

func GeneratePrimes() []int {
	var primes []int = make([]int, 0, 25)
	for no := 2; no <= 100; no++ {
		if IsPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
}

func Benchmark_GeneratePrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePrimes()
	}
}
