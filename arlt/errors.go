package arlt

import (
	"fmt"
)

// RateLimitExceededError represents error that will be raised
// when rate limit is exceeded.
type RateLimitExceededError int64

func (r RateLimitExceededError) Error() string {
	return fmt.Sprintf("rate limit exceeded: %d", r)
}
