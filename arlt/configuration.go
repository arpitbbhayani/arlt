package arlt

import (
	"encoding/binary"
)

// Configuration represents one ratelimit configuration
type Configuration struct {
	MaxTicksPerWindow       uint32
	WindowdurationInSeconds uint16
}

func configurationToByteSlice(c *Configuration) []byte {
	var b [6]byte
	// MaxTicksPerWindow, 4 bytes, big endian
	binary.BigEndian.PutUint32(b[0:4], uint32(c.MaxTicksPerWindow))
	// WindowdurationInSeconds, 4 bytes, big endian
	binary.BigEndian.PutUint16(b[4:6], uint16(c.WindowdurationInSeconds))
	return b[:]
}

func configurationFromByteSclice(b []byte) *Configuration {
	return &Configuration{
		MaxTicksPerWindow:       binary.BigEndian.Uint32(b[:4]),
		WindowdurationInSeconds: binary.BigEndian.Uint16(b[4:6]),
	}
}

// Serialize returns a string representation of configuration
// details that will be persisted
func (c Configuration) Serialize() []byte {
	return configurationToByteSlice(&c)
}

func (c Configuration) String() string {
	return string(c.Serialize())
}

// ConfigurationFromBytes returns a configuration object given a bytes slice
func ConfigurationFromBytes(b []byte) *Configuration {
	return configurationFromByteSclice(b)
}
