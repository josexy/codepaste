package api

import (
	"github.com/gin-gonic/gin"
	"github.com/josexy/gocodepaste/api/v1/common"
	"github.com/josexy/gocodepaste/serializer"
)

// Ping
// @Summary ping 测试
// @Schemes
// @Description ping 测试
// @Tags PING
// @Produce json
// @Success 200 {object} serializer.Response
// @Router /ping [get]
func Ping(c *gin.Context) {
	common.ResponseJson(c, serializer.Response{
		Code: 0,
		Msg:  "Pong",
	})
}
