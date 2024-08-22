package gcd

import (
	"golang.org/x/exp/constraints"
)

// GCD returns the greatest common divisor of integers a and b.
// GCD(0, 0) returns 0.
func GCD[T constraints.Integer](a, b T) T {
	return euclidean(a, b)
}

// euclidean returns the greater common divisor of integers a and b.
// The GCD is calculated using the Euclidean algorithm, see https://en.wikipedia.org/wiki/Euclidean_algorithm
func euclidean[T constraints.Integer](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	
	return a
}

// euclidean returns the greater common divisor of integers a and b.
// The GCD is calculated using the binary algorithm, see https://en.wikipedia.org/wiki/Binary_GCD_algorithm
func binary[T constraints.Integer](a, b T) T {
	if a == 0 && b == 0 {
		return 0
	}
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	var d int
	for a % 2 == 0 && b % 2 == 0 {
		// While both a and b are even, divide each by 2.
		// Count the number of times we do this.
		a /= 2
		b /= 2
		d++
	}
	
	divideBy2UntilOdd := func(v T) T {
		for v % 2 == 0 {
			v /= 2
		}
		return v
	}
	
	// Ensure both a and b are odd.
	a = divideBy2UntilOdd(a)
	b = divideBy2UntilOdd(b)
	
	for a != b {
		if a > b {
			a -= b
			a = divideBy2UntilOdd(a)
		} else {
			b -= a
			b = divideBy2UntilOdd(b)
		}
	}
	
	return (1 << d) * a
}