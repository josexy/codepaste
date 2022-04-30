package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	// AppConfig 全局配置对象
	AppConfig = new(Config)
	// DB 全局MySQL数据库对象
	DB *gorm.DB
	// Redis redis全局对象
	Redis redis.UniversalClient
	// Logger 日志
	Logger *logrus.Logger
)
