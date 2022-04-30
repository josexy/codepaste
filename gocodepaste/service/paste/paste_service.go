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

const maxContentSize = 2000

type CommonCreateService struct {
	Title        string `json:"title" binding:"max=25"`
	Lang         string `json:"lang" binding:"max=10"`
	Content      string `json:"content" binding:"required,min=1"`
	Password     string `json:"password" binding:"max=20"`
	ExpireSecond int64  `json:"expire_second"`
}

type GetService struct {
}

func (service *GetService) Get(key, password string) serializer.Response {
	key = strings.TrimSpace(key)
	password = strings.TrimSpace(password)

	var paste model.Paste
	var pastex serializer.Paste

	data := global.CheckAndGet(-1, key, password)
	if data != nil {
		err := json.Unmarshal(data, &pastex)
		if err != nil {
			return serializer.BuildResponse(codes.Error)
		}
	} else {
		now := time.Now()
		err := global.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Where("`key` = ?", key).First(&paste).Error; err != nil {
				return err
			}

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

			if paste.Private {
				isDeletePaste = false
			}
			if paste.Password != "" {
				if password == "" {
					isDeletePaste = false
				} else if paste.Password != password {
					isDeletePaste = false
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

		// 私有paste
		if paste.Private {
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
		pastex = serializer.BuildPaste(paste)
		data, err = json.Marshal(pastex)
		if err != nil {
			return serializer.BuildResponse(codes.Error)
		}
		var remaining time.Duration
		switch paste.ExpireCount {
		case -1:
			remaining = time.Second * 30
		case 1:
			remaining = paste.RemainingTime(now)
		}
		global.CheckAndSet(-1, false, key, paste.Password, data, remaining)
	}
	return serializer.BuildResponseWithData(codes.Success, pastex)
}

func (service *CommonCreateService) Create() serializer.Response {
	// 未登录用户限制大小
	if len(service.Content) > maxContentSize {
		return serializer.BuildResponse(codes.ErrorPasteContentLimit)
	}

	paste := model.Paste{
		Key:      util.GetUniqueKey(),
		Title:    strings.TrimSpace(service.Title),
		Content:  service.Content,
		Lang:     strings.TrimSpace(service.Lang),
		Password: strings.TrimSpace(service.Password),
		UserID:   1, // 该paste属于公共用户
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
