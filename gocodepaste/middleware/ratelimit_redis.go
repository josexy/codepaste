package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/josexy/gocodepaste/global"
	"github.com/josexy/gocodepaste/limiter"
	"github.com/josexy/gocodepaste/pkg/codes"
)

// RateLimit 分布式限流
func RateLimit() gin.HandlerFunc {
	limit := limiter.NewTokenBucketLimiter(global.Redis.(*redis.Client), 60, 5)
	return func(ctx *gin.Context) {
		key := fmt.Sprintf("limit:%s", ctx.ClientIP())
		if !limit.Allow(context.Background(), key) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": codes.Error,
				"msg":  "api rate limit exceeded",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
