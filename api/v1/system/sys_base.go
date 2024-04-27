package system

import (
	"gin-one/message"
	"gin-one/model/common/response"
	"gin-one/model/system/request"
	"gin-one/service/system"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SysBaseApi struct{}

func (api *SysBaseApi) GetCaptcha(c *gin.Context) {
	ip := c.ClientIP()
	res, err := baseService.GetCaptcha(ip)
	if err != nil {
		response.FailWithMessage(message.QUERY_ERR, err.Error(), c)
		return
	}

	response.OkWithDetailed(res, message.QUERY_OK, c)
}

func (api *SysBaseApi) CheckCaptcha(c *gin.Context) {
	var data request.CheckCaptchaReq
	var ip = c.ClientIP()
	var err error
	if err = c.ShouldBind(&data); err != nil {
		response.FailWithMessage(message.OPER_ERR, err.Error(), c)
		return
	}

	msgCode, res, err := baseService.CheckCaptchaCode(ip, data.ID, data.CaptchaCode)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}
	response.OkWithDetailed(res, msgCode, c)
}

func (api *SysBaseApi) InitDB(c *gin.Context) {
	system.ServiceGroupApp.InitDBService.InitData()
	c.String(http.StatusOK, "OK")
}
