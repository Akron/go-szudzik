package szudzik

import "math"

// Pairing computes the pairing function following Szudzik (2016).
// It requires x and y to be non-negative.
func Pairing(x, y uint64) uint64 {
	if x < y {
		return y*y + x
	}
	return x*x + x + y
}

// Unpairing computes the inverse of the pairing function following Szudzik (2016).
// It requires z to be non-negative.
func Unpairing(z uint64) (x, y uint64) {
	q := math.Floor(math.Sqrt(float64(z)))
	m := float64(z) - q*q
	if m < q {
		return uint64(m), uint64(q)
	}
	return uint64(q), uint64(m - q)
}
