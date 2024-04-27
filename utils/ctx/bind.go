package ctx

import (
	"gin-one/message"
	"gin-one/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// MustBindWithCtx
//
//	@Description: 校验 ctx 参数，并处理错误
//	@param c
//	@param data
//	@return bool
func MustBindWithCtx(c *gin.Context, data any) bool {
	if err := c.ShouldBind(data); err != nil {
		response.FailWithMessage(message.REQ_DATA_ERR, err.Error(), c)
		return false
	}
	return true
}

func MusQueryWithCtx(c *gin.Context, data any) bool {
	if err := c.ShouldBindQuery(data); err != nil {
		response.FailWithMessage(message.REQ_DATA_ERR, err.Error(), c)
		return false
	}
	return true
}

func MustBindWithCopy(c *gin.Context, reqData any, data any) bool {
	if !MustBindWithCtx(c, reqData) {
		return false
	}

	err := copier.Copy(data, reqData)
	if err != nil {
		response.FailWithMessage(message.DATA_STRUCT_ERR, err.Error(), c)
		return false
	}
	return true
}
