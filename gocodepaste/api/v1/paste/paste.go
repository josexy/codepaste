package paste

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/josexy/gocodepaste/api/v1/common"
	"github.com/josexy/gocodepaste/service/auth/paste"
)

// List
// @Summary 获取所有便利贴
// @Schemes
// @Description 获取所有便利贴
// @Tags PASTES
// @Produce json
// @Success 200 {object} serializer.Response "成功"
// @Security BearerAuth
// @Param Authorization header string true "JWT"
// @Param page query string false "page"
// @Param page_size query string false "page_size"
// @Router /auth/pastes [get]
func List(c *gin.Context) {
	var service paste.ListService
	if err := c.ShouldBind(&service); err == nil {
		common.ResponseJson(c, service.List(common.GetUserId(c)))
	} else {
		common.ResponseJsonError(c, err)
	}
}

// Get
// @Summary 获取便利贴
// @Schemes
// @Description 获取便利贴
// @Tags PASTES
// @Produce json
// @Success 200 {object} serializer.Response "成功"
// @Param key path string true "key"
// @Param password query string false "密码"
// @Security BearerAuth
// @Param Authorization header string true "JWT"
// @Router /auth/pastes/:key [get]
func Get(c *gin.Context) {
	var service paste.GetService
	common.ResponseJson(c, service.Get(common.GetUserId(c), c.Param("key"), c.Query("password")))
}

// Create
// @Summary 创建便利贴
// @Schemes
// @Description 创建便利贴
// @Tags PASTES
// @Accept json
// @Produce json
// @Param data body paste.CreateService true "信息"
// @Success 200 {object} serializer.Response "成功"
// @Security BearerAuth
// @Param Authorization header string true "JWT"
// @Router /auth/pastes [post]
func Create(c *gin.Context) {
	var service paste.CreateService
	if err := c.ShouldBind(&service); err == nil {
		common.ResponseJson(c, service.Create(common.GetUserId(c)))
	} else {
		common.ResponseJsonError(c, err)
	}
}

// Delete
// @Summary 删除便利贴
// @Schemes
// @Description 删除便利贴
// @Tags PASTES
// @Accept json
// @Produce json
// @Success 200 {object} serializer.Response "成功"
// @Param id path string true "id"
// @Security BearerAuth
// @Param Authorization header string true "JWT"
// @Router /auth/pastes/:id [delete]
func Delete(c *gin.Context) {
	var service paste.DeleteService
	if id, err := strconv.Atoi(c.Param("id")); err != nil {
		common.ResponseJsonError(c, err)
	} else {
		common.ResponseJson(c, service.Delete(common.GetUserId(c), id))
	}
}

// Update
// @Summary 更新便利贴
// @Schemes
// @Description 更新便利贴
// @Tags PASTES
// @Accept json
// @Produce json
// @Param data body paste.UpdateService true "信息"
// @Success 200 {object} serializer.Response "成功"
// @Param id path string true "id"
// @Security BearerAuth
// @Param Authorization header string true "JWT"
// @Router /auth/pastes/:id [put]
func Update(c *gin.Context) {
	var service paste.UpdateService
	if err := c.ShouldBind(&service); err == nil {
		if id, err := strconv.Atoi(c.Param("id")); err != nil {
			common.ResponseJsonError(c, err)
		} else {
			common.ResponseJson(c, service.Update(common.GetUserId(c), id))
		}
	} else {
		common.ResponseJsonError(c, err)
	}
}
