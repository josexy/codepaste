package common

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/josexy/gocodepaste/pkg/codes"
	"github.com/josexy/gocodepaste/serializer"
)

const identityKey = "data"

func ResponseJsonError(c *gin.Context, err error) {
	if _, ok := err.(validator.ValidationErrors); ok {
		ResponseJson(c, serializer.BuildResponse(codes.ErrorValidation))
	} else {
		ResponseJson(c, serializer.BuildResponse(codes.ErrorParameter))
	}
}

func ResponseJson(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}

func ResponseError(c *gin.Context, code int, message string) {
	c.JSON(code, serializer.BuildError(code, message))
}

func GetUserId(c *gin.Context) int {
	data := jwt.ExtractClaims(c)[identityKey]
	id := int(data.(float64))
	return id
}
