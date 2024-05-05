package response

import (
	"gin-one/global"
	"gin-one/message"
	"gin-one/model/common/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

const (
	SUCCESS = 200
	NO_AUTH = 401
	ERROR   = 500
)

type Response struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
	ErrMsg string      `json:"errMsg"`
}

type PageResponse struct {
	Page     int
	PageSize int
	Total    int64
	Records  interface{} `json:"records"`
}

func Result(code int, data interface{}, msg string, errMsg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
		errMsg,
	})
}

func PageWitheResult(code int, data PageResponse, msg string, errMsg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
		errMsg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, global.Msg.Msg(message.OPER_OK), "", c)
}

func OkWithMessage(msgCode int, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, global.Msg.Msg(msgCode), "", c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, global.Msg.Msg(message.QUERY_OK), "", c)
}

func OkWithDetailed(data interface{}, msgCode int, c *gin.Context) {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Slice {
		if val.Len() == 0 && val.IsNil() {
			data = make(map[string]interface{}, 0)
		}
	}

	code := SUCCESS
	if !global.Msg.IsOkCode(msgCode) {
		code = ERROR
	}
	Result(code, data, global.Msg.Msg(msgCode), "", c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, global.Msg.Msg(message.OPER_FAILED), "", c)
}

func FailWithMessage(msgCode int, errMsg string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, global.Msg.Msg(msgCode), errMsg, c)
}

func FailWithDetailed(data interface{}, msgCode int, errMsg string, c *gin.Context) {
	Result(ERROR, data, global.Msg.Msg(msgCode), errMsg, c)
}

func FailWithPage(pageInfo request.PageInfo, msgCode int, errMsg string, c *gin.Context) {
	PageWitheResult(ERROR, PageResponse{
		pageInfo.Page,
		pageInfo.PageSize,
		0,
		make([]interface{}, 0),
	}, global.Msg.Msg(msgCode), errMsg, c)
}

func WrapPageData(pageInfo request.PageInfo, count int64, data interface{}) PageResponse {
	val := reflect.ValueOf(data)
	if val.IsNil() {
		data = make([]interface{}, 0)
	}
	return PageResponse{
		pageInfo.Page,
		pageInfo.PageSize,
		count,
		data,
	}
}

func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		NO_AUTH,
		nil,
		message,
		"",
	})
}
