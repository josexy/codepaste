package global

import (
	"io/ioutil"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type JwtConfig struct {
	JwtSecret                  string `yaml:"jwt_secret"`
	JwtAccessTokenExpiredTime  int64  `yaml:"jwt_access_token_expired"`
	JwtRefreshTokenExpiredTime int64  `yaml:"jwt_refresh_token_expired"`
}

type ServerConfig struct {
	Address    string `yaml:"address"`
	Mode       string `yaml:"mode"`
	LogFileDir string `yaml:"logfile_dir"`
}

// Config 全局配置
type Config struct {
	Server *ServerConfig `yaml:"server"`
	Jwt    *JwtConfig    `yaml:"jwt"`
	MySQL  *MySQLConfig  `yaml:"mysql"`
	Redis  *RedisConfig  `yaml:"redis"`
}

// InitConfig 初始化配置项
func InitConfig(configPath string) {

	// 加载环境变量
	godotenv.Load()

	config, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(config, AppConfig); err != nil {
		panic(err)
	}

	readEnv()

	InitLogger(AppConfig.Server)
	InitMySQL(AppConfig.MySQL)
	InitRedis(AppConfig.Redis)
}
