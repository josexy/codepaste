package global

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const luaDelScript = `
if redis.call("GET", KEYS[1]) == ARGV[1] then
	return redis.call("DEL", KEYS[1])
else
	return 0
end
`

type Locker interface {
	Get() bool
	Block(seconds int64) bool
	Release() bool
	ForceRelease()
}

// Redis分布式锁
type lock struct {
	ctx     context.Context
	name    string
	owner   string
	seconds int64
}

// Get 通过 setnx 获取分布式锁
func (l *lock) Get() bool {
	return Redis.SetNX(l.ctx, l.name, l.owner, time.Duration(l.seconds)*time.Second).Val()
}

// Block 阻塞一段时间后再次获得锁
func (l *lock) Block(seconds int64) bool {
	start := time.Now().Unix()
	for {
		if l.Get() {
			return true
		} else {
			// 每次阻塞1秒，直到seconds
			time.Sleep(time.Second * time.Duration(1))
			if time.Now().Unix()-seconds >= start {
				return false
			}
		}
	}
}

func (l *lock) Release() bool {
	// 通过LUA脚本释放锁
	luaScript := redis.NewScript(luaDelScript)
	result := luaScript.Run(l.ctx, Redis, []string{l.name}, l.owner).Val().(int64)
	return result != 0
}

func (l *lock) ForceRelease() {
	Redis.Del(l.ctx, l.name).Val()
}

func Lock(name string, seconds int64) Locker {
	return &lock{
		ctx:     context.Background(),
		name:    name,
		owner:   uuid.New().String(),
		seconds: seconds,
	}
}
