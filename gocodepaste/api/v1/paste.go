package api

import (
	"github.com/gin-gonic/gin"
	"github.com/josexy/gocodepaste/api/v1/common"
	"github.com/josexy/gocodepaste/service/paste"
)

// Get
// @Summary 获取便利贴
// @Schemes
// @Description 获取便利贴
// @Tags PUBLIC
// @Produce json
// @Success 200 {object} serializer.Response "成功"
// @Param key path string true "key"
// @Param password query string false "密码"
// @Router /paste/{key} [get]
func Get(c *gin.Context) {
	var service paste.GetService
	common.ResponseJson(c, service.Get(c.Param("key"), c.Query("password")))
}

// Create
// @Summary 创建便利贴
// @Schemes
// @Description 创建便利贴
// @Tags PUBLIC
// @Produce json
// @Param data body paste.CommonCreateService true "信息"
// @Success 200 {object} serializer.Response "成功"
// @Router /paste [post]
func Create(c *gin.Context) {
	var service paste.CommonCreateService
	if err := c.ShouldBind(&service); err == nil {
		common.ResponseJson(c, service.Create())
	} else {
		common.ResponseJsonError(c, err)
	}
}
