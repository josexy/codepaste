package router

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/josexy/gocodepaste/api/v1"
	"github.com/josexy/gocodepaste/api/v1/paste"
	"github.com/josexy/gocodepaste/api/v1/user"
	_ "github.com/josexy/gocodepaste/docs"
	"github.com/josexy/gocodepaste/global"
	"github.com/josexy/gocodepaste/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {

	if global.AppConfig.Server.Mode == "release" {
		gin.DefaultWriter = ioutil.Discard
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(middleware.RateLimit(), middleware.Cors(), middleware.Logger())
	middleware.NewJWTAuthMiddleWare(&middleware.AllUserAuthorizator{})

	// Swagger接口文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// page not found
	r.NoRoute(api.NotFoundPage)

	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)

		v1.POST("user/register", user.Register)
		v1.POST("user/login", middleware.AuthMiddleware.LoginHandler)

		auth := v1.Group("/auth")
		auth.GET("/refresh_token", middleware.AuthMiddleware.RefreshHandler)

		auth.Use(middleware.AuthMiddleware.MiddlewareFunc())
		{
			u := auth.Group("/user")
			{
				u.GET("", user.Get)
				u.PUT("", user.Update)
				u.DELETE("", middleware.AuthMiddleware.LogoutHandler)
			}

			p := auth.Group("/pastes")
			{
				p.GET("", paste.List)
				p.GET("/:key", paste.Get)
				p.POST("", paste.Create)
				p.PUT("/:id", paste.Update)
				p.DELETE("/:id", paste.Delete)
			}
		}
		p := v1.Group("/paste")
		{
			p.GET("/:key", api.Get)
			p.POST("", api.Create)
		}
	}
	return r
}
