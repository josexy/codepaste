package global

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Hostname   string   `yaml:"hostname"`
	Port       int      `yaml:"port"`
	Password   string   `yaml:"password"`
	DB         int      `yaml:"db"`
	Cluster    bool     `yaml:"cluster_enable"`
	CLusterIPs []string `yaml:"cluster_ips"`
	Addr       string
}

func InitRedis(config *RedisConfig) {
	if config.Cluster {
		Redis = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        config.CLusterIPs,
			Password:     config.Password,
			DialTimeout:  10 * time.Second,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		})
	} else {
		Redis = redis.NewClient(&redis.Options{
			Addr:         config.Addr,
			Password:     config.Password,
			DB:           config.DB,
			DialTimeout:  10 * time.Second,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		})
	}
	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

// pastes:keys = {key1:1, key2:2, ...}
// public:{key}
// protect:{key}:{password}
// private:{key}:{id}
const (
	keySet  = "pastes:keys"
	public  = "public"
	protect = "protect"
	private = "private"
)

func getPasteKey(typ, key, arg string) string {
	if typ == "public" {
		return fmt.Sprintf("%s:%s", typ, key)
	}
	return fmt.Sprintf("%s:%s:%s", typ, key, arg)
}

func ZAdd(key string, t time.Duration) {
	now := time.Now().Add(t)
	Redis.ZAdd(context.Background(), keySet, &redis.Z{Score: float64(now.Unix()), Member: key})
}

func ZExist(key string) bool {
	return Redis.ZScore(context.Background(), keySet, key).Err() == nil
}

func ZDel(key string) {
	Redis.ZRem(context.Background(), keySet, key)
}

func ZDelExpireKey() {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	Redis.ZRemRangeByScore(context.Background(), keySet, "-inf", fmt.Sprintf("(%s", now))
}

func Del(typ, key, arg string) {
	Redis.Del(context.Background(), getPasteKey(typ, key, arg))
}

func Get(typ, key, arg string) []byte {
	res, err := Redis.Get(context.Background(), getPasteKey(typ, key, arg)).Bytes()
	if err != nil {
		return nil
	}
	return res
}

func Set(typ, key, arg string, value interface{}, second time.Duration) {
	Redis.SetEX(context.Background(), getPasteKey(typ, key, arg), value, second)
}

func CheckAndDel(id int, pri bool, key, password string) {
	typ, arg := public, ""
	if pri {
		typ = private
		arg = strconv.Itoa(id)
	} else if password != "" {
		typ = protect
		arg = password
	}
	ZDel(key)
	Del(typ, key, arg)
}

func CheckAndSet(id int, pri bool, key, password string, value interface{}, second time.Duration) {
	typ, arg := public, ""
	if pri {
		typ = private
		arg = strconv.Itoa(id)
	} else if password != "" {
		typ = protect
		arg = password
	}
	ZAdd(key, second)
	Set(typ, key, arg, value, second)
}

func CheckAndGet(id int, key, password string) (data []byte) {
	// 删除过期的keys
	ZDelExpireKey()

	// key不存在
	if !ZExist(key) {
		return
	}
	// public，不需要密码
	if data = Get(public, key, ""); data != nil {
		return
	}
	// protect，需要密码
	if password != "" {
		if data = Get(protect, key, password); data != nil {
			return
		}
	}
	// private
	if id > 0 {
		return Get(private, key, strconv.Itoa(id))
	}
	return
}
