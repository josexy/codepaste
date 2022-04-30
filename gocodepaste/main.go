package main

import (
	"flag"

	"github.com/josexy/gocodepaste/global"
	"github.com/josexy/gocodepaste/router"
)

var configPath string

func InitParse() {
	flag.StringVar(&configPath, "c", "conf/config.yaml", "config file path")
}

// @title Swagger 接口文档
// @version 1.0
// @description 使用 Go-Swagger 自动生成的接口文档
// @termsOfService http://swagger.io/terms/

// @contact.name josexy
// @contact.url http://github.com/josexy
// @contact.email josephxrays@gmail.com

// @license.name MIT
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:10086
// @BasePath /api/v1
func main() {

	InitParse()
	flag.Parse()

	global.InitConfig(configPath)
	svr := NewServer(global.AppConfig.Server.Address, router.NewRouter())
	svr.Run()
}
