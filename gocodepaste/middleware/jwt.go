package middleware

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/josexy/gocodepaste/api/v1/common"
	"github.com/josexy/gocodepaste/api/v1/user"
	"github.com/josexy/gocodepaste/global"
	"github.com/josexy/gocodepaste/pkg/codes"
	"github.com/josexy/gocodepaste/serializer"
)

var (
	realm          = "gocodepaste"
	identityKey    = "data"
	AuthMiddleware *jwt.GinJWTMiddleware
)

type IAuthorizator interface {
	HandleAuthorizator(data interface{}, c *gin.Context) bool
}

type AllUserAuthorizator struct {
}

func (*AllUserAuthorizator) HandleAuthorizator(data interface{}, c *gin.Context) bool {
	return true
}

func responseLogout(c *gin.Context, code int) {
	common.ResponseJson(c, serializer.BuildResponse(codes.Success))
}

func responseError(c *gin.Context, code int, message string) {
	common.ResponseError(c, code, message)
}

func responseToken(c *gin.Context, code int, token string, time time.Time) {
	common.ResponseJson(c, serializer.BuildResponseWithData(codes.Success,
		serializer.Token{
			Token:  token,
			Expire: time,
		}))
}

func NewJWTAuthMiddleWare(authorizator IAuthorizator) {
	AuthMiddleware = &jwt.GinJWTMiddleware{
		Realm:         realm,
		Timeout:       time.Minute * time.Duration(global.AppConfig.Jwt.JwtAccessTokenExpiredTime),
		MaxRefresh:    time.Minute * time.Duration(global.AppConfig.Jwt.JwtRefreshTokenExpiredTime),
		Key:           []byte(global.AppConfig.Jwt.JwtSecret),
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		IdentityKey:   identityKey,
		Authenticator: user.Login,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if data != nil {
				// 将用户ID保存到JWT的负载信息中，并返回给客户端
				if v, ok := data.(*serializer.User); ok {
					return jwt.MapClaims{
						identityKey: int(v.ID),
					}
				}
			}
			return jwt.MapClaims{}
		},
		LoginResponse:   responseToken,
		RefreshResponse: responseToken,
		LogoutResponse:  responseLogout,
		/*
			bool, for JSON booleans
			float64, for JSON numbers
			string, for JSON strings
			[]interface{}, for JSON arrays
			map[string]interface{}, for JSON objects
			nil for JSON null
		*/
		IdentityHandler: func(c *gin.Context) interface{} {
			data := jwt.ExtractClaims(c)[identityKey]
			id := int(data.(float64))
			return id
		},
		Authorizator: authorizator.HandleAuthorizator,
		Unauthorized: responseError,
	}
	if err := AuthMiddleware.MiddlewareInit(); err != nil {
		global.Logger.Error(err)
	}
}
