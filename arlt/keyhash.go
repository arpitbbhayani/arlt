package arlt

// Key represents key against which rate limit will
// be configured and applied
type Key string

// Hash returns a hashed version of key k
func (k *Key) Hash() string {
	return string(*k)
}
