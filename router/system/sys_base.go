package system

import (
	v1 "gin-one/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (router *BaseRouter) InitBaseRouter(pubEngin *gin.RouterGroup) {
	pub := pubEngin.Group("/base")
	{
		pub.GET("getCaptcha", v1.ApiGroupApp.SystemApiGroup.SysBaseApi.GetCaptcha)
		pub.PUT("checkCaptcha", v1.ApiGroupApp.SystemApiGroup.SysBaseApi.CheckCaptcha)
		pub.POST("initDb", v1.ApiGroupApp.SystemApiGroup.SysBaseApi.InitDB)
	}
}
