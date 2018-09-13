package arlt

import (
	"testing"
)

func TestKeyHash(t *testing.T) {
	positives := []string{
		"random-key",
		"random-key-2",
	}

	for i := range positives {
		k := Key(positives[i])
		if len([]byte(k.Hash())) != 4 {
			t.Error("crc not of length 4")
		}
	}
}
