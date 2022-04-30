package paste

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/josexy/gocodepaste/global"
	"github.com/josexy/gocodepaste/model"
	"github.com/josexy/gocodepaste/pkg/codes"
	"github.com/josexy/gocodepaste/pkg/util"
	"github.com/josexy/gocodepaste/serializer"
	"gorm.io/gorm"
)

const maxPerPageSize = 10

type CreateService struct {
	Title        string `json:"title" binding:"max=25"`
	Lang         string `json:"lang" binding:"max=10"`
	Content      string `json:"content" binding:"required,min=1"`
	Password     string `json:"password" binding:"max=20"`
	Private      bool   `json:"private"`
	ExpireSecond int64  `json:"expire_second"`
}

type DeleteService struct {
}

type GetService struct {
}

type ListService struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

type UpdateService struct {
	Title    string `json:"title" binding:"max=25"`
	Lang     string `json:"lang" binding:"max=10"`
	Content  string `json:"content" binding:"required,min=1"`
	Password string `json:"password" binding:"max=20"`
	Private  bool   `json:"private"`
}

func (service *UpdateService) Update(userId, id int) serializer.Response {
	paste := model.Paste{
		Title:    strings.TrimSpace(service.Title),
		Lang:     strings.TrimSpace(service.Lang),
		Content:  service.Content,
		Password: strings.TrimSpace(service.Password),
		Private:  service.Private,
	}
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		var p model.Paste
		if err := tx.Where("user_id = ? and id = ?", userId, id).First(&p).Error; err != nil {
			return err
		}
		global.CheckAndDel(userId, p.Private, p.Key, p.Password)
		// 更新多列，即使是零值也要更新
		if err := tx.Select("title", "lang", "content", "password", "private").
			Where("user_id = ? and id = ?", userId, id).Updates(&paste).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return serializer.BuildResponse(codes.ErrorPasteNotExist)
		} else {
			return serializer.BuildResponse(codes.Error)
		}
	}
	return serializer.BuildResponse(codes.Success)
}

func (service *ListService) List(id int) serializer.ResponsePagination {
	// 分页
	if service.Page <= 0 {
		service.Page = 1
	}
	if service.PageSize <= 0 || service.PageSize > maxPerPageSize {
		service.PageSize = maxPerPageSize
	}

	var pastes []model.Paste
	var count int64
	// 页内项数
	size := service.PageSize

	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// example: select * from pastes where user_id=2 limit 3 offset 6;
		txx := tx.Where("user_id = ?", id).
			Limit(service.PageSize).Offset((service.Page - 1) * service.PageSize).
			Order("created_at DESC").Find(&pastes)
		if txx.Error != nil {
			return txx.Error
		}
		size = int(txx.RowsAffected)
		return tx.Model(model.Paste{}).Where("user_id = ?", id).Count(&count).Error
	})

	if err != nil {
		return serializer.BuildResponsePagination(codes.Error)
	}
	return serializer.BuildResponseWithDataPagination(
		codes.Success,
		serializer.BuildPastesList(pastes),
		service.Page,
		size,
		int(count),
	)
}

func (service *GetService) Get(userId int, key, password string) serializer.Response {
	password = strings.TrimSpace(password)

	var pastex serializer.Paste

	data := global.CheckAndGet(userId, key, password)
	if data != nil {
		err := json.Unmarshal(data, &pastex)
		if err != nil {
			return serializer.BuildResponse(codes.Error)
		}
	} else {
		var paste model.Paste

		now := time.Now()
		err := global.DB.Transaction(func(tx *gorm.DB) error {
			// 获取数据
			if err := tx.Where("`key` = ?", key).First(&paste).Error; err != nil {
				return err
			}

			// 查询成功
			isDeletePaste := false
			// 判断是永久、阅后即焚还是过期
			switch paste.ExpireCount {
			case 0:
				// 在访问后立即删除
				isDeletePaste = true
			case 1:
				// 过期删除
				if paste.IsExpired(now) {
					isDeletePaste = true
				}
			}
			// 验证密码是否正确、是否私有
			if int(paste.UserID) != userId {
				// 私有paste，只有已经登录的用户才可以创建私有的paste
				if paste.UserID != 1 && paste.Private {
					isDeletePaste = false
				}
				// 需要密码
				if paste.Password != "" {
					if password == "" {
						isDeletePaste = false
					} else if paste.Password != password {
						isDeletePaste = false
					}
				}
			}
			if isDeletePaste {
				return tx.Where("id = ?", paste.ID).Delete(&paste).Error
			}
			return nil
		})

		if err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				return serializer.BuildResponse(codes.ErrorPasteNotExist)
			default:
				return serializer.BuildResponse(codes.Error)
			}
		}
		if paste.ExpireCount == 1 && paste.IsExpired(now) {
			return serializer.BuildResponse(codes.ErrorPasteNotExist)
		}

		// 已登录用户访问其他用户（包括公开用户id=1）的paste
		if int(paste.UserID) != userId {
			// 私有paste，只有已经登录的用户才可以创建私有的paste
			if paste.UserID != 1 && paste.Private {
				return serializer.BuildResponse(codes.ErrorPasteAccessDenied)
			}
			// 需要密码
			if paste.Password != "" {
				if password == "" {
					return serializer.BuildResponse(codes.ErrorPasteNeedPassword)
				} else if paste.Password != password {
					return serializer.BuildResponse(codes.ErrorPastePasswordIncorrect)
				}
			}
		}

		pastex = serializer.BuildPaste(paste)
		// 阅后即焚，无需缓存
		if paste.ExpireCount == 0 {
			return serializer.BuildResponseWithData(codes.Success, pastex)
		}
		data, err = json.Marshal(pastex)
		if err != nil {
			return serializer.BuildResponse(codes.Error)
		}

		var remaining time.Duration
		switch paste.ExpireCount {
		case -1:
			remaining = time.Hour
		case 1:
			remaining = paste.RemainingTime(now)
		}
		global.CheckAndSet(userId, paste.Private, key, paste.Password, data, remaining)
	}
	return serializer.BuildResponseWithData(codes.Success, pastex)
}

func (service *DeleteService) Delete(userId, id int) serializer.Response {
	var paste model.Paste
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ? and id = ?", userId, id).First(&paste).Error; err != nil {
			return err
		}
		global.CheckAndDel(userId, paste.Private, paste.Key, paste.Password)
		if err := tx.Where("user_id = ? and id = ?", userId, id).Delete(&paste).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return serializer.BuildResponse(codes.ErrorPasteNotExist)
		} else {
			return serializer.BuildResponse(codes.Error)
		}
	}
	return serializer.BuildResponse(codes.Success)
}

func (service *CreateService) Create(userId int) serializer.Response {
	paste := model.Paste{
		Key:      util.GetUniqueKey(),
		Title:    strings.TrimSpace(service.Title),
		Lang:     strings.TrimSpace(service.Lang),
		Content:  service.Content,
		Password: strings.TrimSpace(service.Password),
		Private:  service.Private,
		UserID:   uint(userId),
	}

	if service.ExpireSecond < 0 {
		// 永久
		paste.ExpireCount = -1
	} else if service.ExpireSecond == 0 {
		// 阅后即焚
		paste.ExpireCount = 0
	} else {
		// 过期时间
		paste.ExpireCount = 1
		paste.ExpiredAt = time.Now().Add(time.Second * time.Duration(service.ExpireSecond))
	}
	if err := global.DB.Create(&paste).Error; err != nil {
		return serializer.BuildResponse(codes.Error)
	}
	paste.Content = ""
	return serializer.BuildResponseWithData(
		codes.Success,
		serializer.BuildPaste(paste),
	)
}
