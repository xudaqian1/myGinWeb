package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"myGinWeb/pkg/setting"
	"time"
)

// Limiter 定义属性
type Limiter struct {
	// Redis client connection.
	rc *redis.Client
}

// NewLimiter 根据redisURL创建新的limiter并返回
func NewLimiter(config setting.RedisConfig) (*Limiter, error) {
	rc := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password, // no password set
		DB:       0,               // use default DB
	})
	if err := rc.Ping().Err(); err != nil {
		return nil, err
	}

	return &Limiter{rc: rc}, nil
}

// Allow 通过redis的value判断第几次访问并返回是否允许访问
func (l *Limiter) Allow(key string, events int64, per time.Duration) bool {
	curr := l.rc.LLen(key).Val()
	if curr >= events {
		return false
	}

	if v := l.rc.Exists(key).Val(); v == 0 {
		pipe := l.rc.TxPipeline()
		pipe.RPush(key, key)
		//设置过期时间
		pipe.Expire(key, per)
		_, _ = pipe.Exec()
	} else {
		l.rc.RPushX(key, key)
	}

	return true
}
