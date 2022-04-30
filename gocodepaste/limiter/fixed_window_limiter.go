package limiter

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

// FixedWindowLimiter 固定窗口限流器
type FixedWindowLimiter struct {
	limit  int           // 窗口请求上限数量
	window int           // 窗口时间大小，也就是窗口过期时间，单位为毫秒
	client *redis.Client // Redis客户端
	script *redis.Script // lua脚本
}

func NewFixedWindowLimiter(client *redis.Client, limit int, window time.Duration) (*FixedWindowLimiter, error) {
	// redis过期时间精度最大到毫秒，因此窗口必须能被毫秒整除
	if window%time.Millisecond != 0 {
		return nil, errors.New("the window uint must not be less than millisecond")
	}
	return &FixedWindowLimiter{
		limit:  limit,
		window: int(window / time.Millisecond),
		client: client,
		script: redis.NewScript(readLuaScript("limiter/lua/fixed_window_limiter.lua")),
	}, nil
}

func (l *FixedWindowLimiter) Allow(ctx context.Context, resource string) bool {
	success, err := l.script.Run(ctx, l.client, []string{resource}, l.window, l.limit).Bool()
	if err != nil {
		return false
	}
	// 若到达窗口请求上限，请求失败
	if !success {
		return false
	}
	return true
}
