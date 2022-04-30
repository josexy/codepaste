package api

import (
	"github.com/gin-gonic/gin"
	"github.com/josexy/gocodepaste/api/v1/common"
	"github.com/josexy/gocodepaste/pkg/codes"
	"github.com/josexy/gocodepaste/serializer"
)

func NotFoundPage(c *gin.Context) {
	common.ResponseJson(c, serializer.BuildResponse(codes.PageNotFound))
}
