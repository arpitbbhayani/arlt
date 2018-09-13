package arlt

// Setting represents possible settings for Rate Limiter
type Setting struct {
	RedisAddress  string
	RedisPassword string
	RedisDB       int
}

// DefaultSetting returns default setting for RateLimiter
func DefaultSetting() *Setting {
	return &Setting{
		RedisAddress:  "localhost:6379",
		RedisPassword: "",
		RedisDB:       0,
	}
}
