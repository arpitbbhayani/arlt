package arlt

import (
	"testing"
)

func areSlicesEqual(a []byte, b []byte) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestConfiguration_deserialize(t *testing.T) {
	positives := [][]uint32{
		[]uint32{1, 1},
		[]uint32{100, 100},
		[]uint32{3000, 3000},
		[]uint32{30000, 30000},
		[]uint32{65535, 65535},
	}

	for i := range positives {
		x := Configuration{
			MaxTicksPerWindow:       positives[i][0],
			WindowdurationInSeconds: uint16(positives[i][1]),
		}
		y := ConfigurationFromBytes(x.Serialize())
		if x.MaxTicksPerWindow != y.MaxTicksPerWindow ||
			x.WindowdurationInSeconds != y.WindowdurationInSeconds {
			t.Errorf("deserialization failed for %v, %v", x, y)
		}
	}
}
