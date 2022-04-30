package limiter

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// LeakyBucketLimiter 漏桶限流器
type LeakyBucketLimiter struct {
	maxLevel        int           // 最高水位
	currentVelocity int           // 水流出恒速度/秒
	client          *redis.Client // Redis客户端
	script          *redis.Script
}

func NewLeakyBucketLimiter(client *redis.Client, peakLevel, currentVelocity int) *LeakyBucketLimiter {
	return &LeakyBucketLimiter{
		maxLevel:        peakLevel,
		currentVelocity: currentVelocity,
		client:          client,
		script:          redis.NewScript(readLuaScript("limiter/lua/leaky_bucket_limiter.lua")),
	}
}

func (l *LeakyBucketLimiter) Allow(ctx context.Context, resource string) bool {
	// 当前时间
	now := time.Now().Unix()
	success, err := l.script.Run(ctx, l.client, []string{resource}, l.maxLevel, l.currentVelocity, now).Bool()
	if err != nil {
		return false
	}
	// 若到达窗口请求上限，请求失败
	if !success {
		return true
	}
	return false
}
