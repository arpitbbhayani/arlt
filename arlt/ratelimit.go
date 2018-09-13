package arlt

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// Arlt represents an instance of ratelimiter
type Arlt struct {
	Setting     *Setting
	RedisClient *redis.Client
}

// NewArlt returns reference to new instance of ratelimiter
func NewArlt(setting *Setting) (*Arlt, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     setting.RedisAddress,
		Password: setting.RedisPassword,
		DB:       setting.RedisDB,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &Arlt{
		Setting:     setting,
		RedisClient: redisClient,
	}, nil

}

// DidLimitExceed returns if the limit has been exceeded for the given key
// and configuration.
func (a *Arlt) DidLimitExceed(key Key, configuration Configuration) (bool, error) {
	currentTime := time.Now()
	windowEndTime := currentTime
	windowStartTime := time.Unix(currentTime.Unix()-int64(configuration.WindowdurationInSeconds), 0)
	windowEndTimeString := strconv.FormatInt(windowEndTime.UnixNano(), 10)
	windowStartTimeString := strconv.FormatInt(windowStartTime.UnixNano(), 10)
	windowDiscardTimeString := strconv.FormatInt(windowStartTime.UnixNano()-1, 10)

	redisClient := a.RedisClient
	kh := key.Hash()

	err := redisClient.Watch(func(tx *redis.Tx) error {
		pipe := tx.TxPipeline()
		pipe.Expire(kh, time.Duration(configuration.WindowdurationInSeconds)*time.Second)
		pipe.ZRemRangeByScore(kh, "0", windowDiscardTimeString)
		cardinalityResponse := pipe.ZCount(kh, windowStartTimeString, windowEndTimeString)
		pipe.Exec()

		if cardinalityResponse.Val() >= int64(configuration.MaxTicksPerWindow) {
			return RateLimitExceededError(cardinalityResponse.Val())
		}

		tx.ZAdd(kh, redis.Z{
			Member: windowEndTimeString,
			Score:  float64(windowEndTime.UnixNano()),
		})

		return nil
	}, kh)

	_, ok := err.(RateLimitExceededError)
	if ok {
		return true, nil
	}

	return false, err
}
