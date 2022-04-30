package serializer

import (
	"time"

	"github.com/josexy/gocodepaste/pkg/codes"
)

type Response struct {
	Code int         `json:"code"`           // 错误码
	Msg  string      `json:"msg"`            // 错误码可读消息
	Data interface{} `json:"data,omitempty"` // 返回此次请求响应的数据
}

type ResponsePagination struct {
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	Data     interface{} `json:"data,omitempty"`
	Page     int         `json:"page,omitempty"`
	PageSize int         `json:"page_size,omitempty"`
	Total    int         `json:"total,omitempty"`
}

type Token struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

func BuildResponse(code int) Response {
	return Response{
		Code: code,
		Msg:  codes.GetCodeMessage(code),
	}
}

func BuildResponsePagination(code int) ResponsePagination {
	return ResponsePagination{
		Code: code,
		Msg:  codes.GetCodeMessage(code),
	}
}

func BuildResponseWithData(code int, data interface{}) Response {
	return Response{
		Code: code,
		Msg:  codes.GetCodeMessage(code),
		Data: data,
	}
}

func BuildResponseWithDataPagination(code int, data interface{}, page, pageSize, total int) ResponsePagination {
	return ResponsePagination{
		Code:     code,
		Msg:      codes.GetCodeMessage(code),
		Data:     data,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}
}

func BuildError(code int, message string) Response {
	return Response{
		Code: code,
		Msg:  message,
	}
}
