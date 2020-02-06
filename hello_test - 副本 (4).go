
package main
import "testing"

func TestPrime(t *testing.T) {
	var primeTests = []struct {
		input    int
		expected bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, false},
	}
	for _, tt := range primeTests {
		actual := isPrime(tt.input)
		if actual != tt.expected {
			t.Errorf("%d %v %v", tt.input, actual, tt.expected)
		}
	}
}
