package szudzik

import "testing"

func TestPairing(t *testing.T) {
	for x := uint64(0); x < 10; x++ {
		for y := uint64(0); y < 10; y++ {
			z := Pairing(x, y)
			x2, y2 := Unpairing(z)
			if x != x2 || y != y2 {
				t.Errorf("Pairing(%d, %d) = %d, Unpairing(%d) = (%d, %d)", x, y, z, z, x2, y2)
			}
		}
	}
}
