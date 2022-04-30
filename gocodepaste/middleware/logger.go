package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/josexy/gocodepaste/global"
)

func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		method := c.Request.Method
		uri := c.Request.RequestURI
		code := c.Writer.Status()
		ip := c.ClientIP()

		global.Logger.Infof("| %3d | %13v | %15s | %s | %s |",
			code,
			latencyTime,
			ip,
			method,
			uri,
		)
	}
}
