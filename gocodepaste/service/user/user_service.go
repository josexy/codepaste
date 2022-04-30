package user

import (
	"strings"

	"github.com/josexy/gocodepaste/global"
	"github.com/josexy/gocodepaste/model"
	"github.com/josexy/gocodepaste/pkg/codes"
	"github.com/josexy/gocodepaste/serializer"
	"gorm.io/gorm"
)

type GetService struct {
}

type LoginService struct {
	UserName string `json:"username" binding:"required,min=5,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type RegisterService struct {
	Nickname string `json:"nickname" binding:"required,min=3,max=20"`
	UserName string `json:"username" binding:"required,min=5,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type UpdateService struct {
	Nickname string `json:"nickname" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

func (service *UpdateService) check() (serializer.Response, bool) {
	if len(service.Nickname) > 0 {
		// 昵称已经存在
		var count int64
		global.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
		if count > 0 {
			return serializer.BuildResponse(codes.ErrorUserNicknameAlreadyExisted), false
		}
	}
	return serializer.Response{}, true
}

func (service *UpdateService) Update(id int) serializer.Response {
	service.Nickname = strings.TrimSpace(service.Nickname)
	service.Password = strings.TrimSpace(service.Password)

	if resp, ok := service.check(); !ok {
		return resp
	}

	user := model.User{
		Nickname: service.Nickname,
	}

	// 加密密码
	if user.EncryptPassword(service.Password) != nil {
		return serializer.BuildResponse(codes.Error)
	}

	// 修改信息
	tx := global.DB.Where("id = ?", id).Updates(&user)
	if tx.Error != nil {
		return serializer.BuildResponse(codes.Error)
	}
	if tx.RowsAffected <= 0 {
		return serializer.BuildResponse(codes.ErrorUserNotExist)
	}
	return serializer.BuildResponse(codes.Success)
}

func (service *RegisterService) check() (serializer.Response, bool) {
	// 用户名已经存在
	var count int64
	global.DB.Model(&model.User{}).Where("username = ?", service.UserName).Count(&count)
	if count > 0 {
		return serializer.BuildResponse(codes.ErrorUserAlreadyExisted), false
	}

	// 昵称已经存在
	count = 0
	global.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return serializer.BuildResponse(codes.ErrorUserNicknameAlreadyExisted), false
	}

	return serializer.Response{}, true
}

func (service *RegisterService) Register() serializer.Response {
	service.Nickname = strings.TrimSpace(service.Nickname)
	service.UserName = strings.TrimSpace(service.UserName)
	service.Password = strings.TrimSpace(service.Password)

	if resp, ok := service.check(); !ok {
		return resp
	}

	user := model.User{
		Nickname: service.Nickname,
		Username: service.UserName,
	}

	// 加密密码
	if user.EncryptPassword(service.Password) != nil {
		return serializer.BuildResponse(codes.Error)
	}

	// 创建用户
	if err := global.DB.Create(&user).Error; err != nil {
		return serializer.BuildResponse(codes.Error)
	}

	return serializer.BuildResponseWithData(
		codes.Success,
		serializer.BuildUser(user),
	)
}

func (service *LoginService) Login() *serializer.User {
	service.UserName = strings.TrimSpace(service.UserName)
	service.Password = strings.TrimSpace(service.Password)

	var user model.User

	if err := global.DB.Where("username = ?", service.UserName).First(&user).Error; err != nil {
		// 不存在记录
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		// 内部错误
		return nil
	}

	// 验证密码
	if !user.CheckPassword(service.Password) {
		return nil
	}
	u := serializer.BuildUser(user)
	return &u
}

func (service *GetService) Get(id int) serializer.Response {
	var user model.User

	if err := global.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return serializer.BuildResponse(codes.ErrorUserNotExist)
	}
	return serializer.BuildResponseWithData(codes.Success, serializer.BuildUser(user))
}
