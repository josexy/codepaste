package global

import (
	"fmt"
	"os"
)

func readEnv() {
	mode := os.Getenv("APP_MODE")
	if mode == "" {
		mode = AppConfig.Server.Mode
	}

	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			AppConfig.MySQL.UserName, AppConfig.MySQL.Password,
			AppConfig.MySQL.Hostname, AppConfig.MySQL.Port, AppConfig.MySQL.DB)
	}

	redisAddress := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PW")
	if redisAddress == "" {
		redisAddress = fmt.Sprintf("%s:%d", AppConfig.Redis.Hostname, AppConfig.Redis.Port)
	}
	if redisPassword == "" {
		redisPassword = AppConfig.Redis.Password
	}

	AppConfig.Server.Mode = mode
	AppConfig.MySQL.DSN = dsn
	AppConfig.Redis.Addr = redisAddress
	AppConfig.Redis.Password = redisPassword

	fmt.Println(AppConfig.MySQL.DSN)
}
