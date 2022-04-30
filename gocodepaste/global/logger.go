package global

import (
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

func InitLogger(config *ServerConfig) {
	Logger = logrus.New()
	// 输出位置
	Logger.Out = setOutputFile(config.LogFileDir)
	// 等级
	Logger.SetLevel(logrus.InfoLevel)
	// 格式
	Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		DisableColors:   true,
	})
}

func setOutputFile(dir string) *os.File {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return os.Stdout
		}
	}

	now := time.Now()
	file := now.Format("2006-01-02") + ".log"
	p := path.Join(dir, file)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		if _, err := os.Create(p); err != nil {
			return os.Stdout
		}
	}
	fp, err := os.OpenFile(p, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return os.Stdout
	}
	return fp
}
