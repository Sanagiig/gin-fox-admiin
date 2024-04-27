package system

import (
	v1 "gin-one/api/v1"
	"github.com/gin-gonic/gin"
)

type FTPRouter struct{}

func (f *FTPRouter) InitApiRouter(pubRouterGroup *gin.RouterGroup, privateRouterGroup *gin.RouterGroup) {
	pub := pubRouterGroup.Group("ftp")
	{
		pub.POST("upload", v1.ApiGroupApp.SystemApiGroup.FTPApi.Upload)
	}
}
