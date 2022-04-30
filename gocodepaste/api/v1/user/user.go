package user

import (
	"fmt"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/josexy/gocodepaste/api/v1/common"
	"github.com/josexy/gocodepaste/service/user"
)

// Register
// @Summary 用户注册
// @Schemes
// @Description 用户注册
// @Tags USER
// @Accept json
// @Produce json
// @Param data body user.RegisterService true "用户信息"
// @Success 200 {object} serializer.Response "成功"
// @Router /user/register [post]
func Register(c *gin.Context) {
	var service user.RegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		common.ResponseJson(c, res)
	} else {
		common.ResponseJsonError(c, err)
	}
}

// Login
// @Summary 用户登陆
// @Schemes
// @Description 用户登陆
// @Tags USER
// @Accept json
// @Produce json
// @Param data body user.LoginService true "用户信息"
// @Success 200 {object} serializer.Response "成功"
// @Failure 401 {object} serializer.Response "失败"
// @Router /user/login [post]
func Login(c *gin.Context) (interface{}, error) {
	var service user.LoginService
	// 绑定请求参数
	if err := c.ShouldBind(&service); err == nil {
		if u := service.Login(); u != nil {
			return u, nil
		} else {
			return nil, jwt.ErrFailedAuthentication
		}
	} else {
		return nil, jwt.ErrMissingLoginValues
	}
}

// Update
// @Summary 修改用户信息
// @Schemes
// @Description 修改用户信息
// @Tags USER
// @Accept json
// @Produce json
// @Param data body user.UpdateService true "用户信息"
// @Success 200 {object} serializer.Response "成功"
// @Security BearerAuth
// @Param Authorization header string true "JWT"
// @Router /auth/user [put]
func Update(c *gin.Context) {
	var service user.UpdateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(common.GetUserId(c))
		common.ResponseJson(c, res)
	} else {
		fmt.Println(err)
		common.ResponseJsonError(c, err)
	}
}

// Get
// @Summary 获取用户信息
// @Schemes
// @Description 获取用户信息
// @Tags USER
// @Produce json
// @Success 200 {object} serializer.Response "成功"
// @Security BearerAuth
// @Param Authorization header string true "JWT"
// @Router /auth/user [get]
func Get(c *gin.Context) {
	var service user.GetService
	common.ResponseJson(c, service.Get(common.GetUserId(c)))
}
