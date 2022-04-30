package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ratelimit2 "github.com/juju/ratelimit"
	ratelimit1 "go.uber.org/ratelimit"
)

func RateLimit1() gin.HandlerFunc {
	// 每秒 100000 个请求
	rl := ratelimit1.New(100000)
	return func(c *gin.Context) {
		if rl.Take().Sub(time.Now()) > 0 {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}

func RateLimit2(fillInterval time.Duration, cap int64) gin.HandlerFunc {
	rl := ratelimit2.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		if rl.TakeAvailable(1) == 1 {
			c.Next()
			return
		}
		c.String(http.StatusOK, "rate limit...")
		c.Abort()
	}
}
