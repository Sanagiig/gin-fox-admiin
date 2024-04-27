package initialize

import (
	"gin-one/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.New()
	pubRouter := Router.Group("")
	privateRouter := Router.Group("")

	systemRouter := router.RouterGroupApp.System
	{
		systemRouter.BaseRouter.InitBaseRouter(pubRouter)
		systemRouter.UserRouter.InitUserRouter(pubRouter, privateRouter)
		systemRouter.RoleRouter.InitRoleRouter(pubRouter, privateRouter)
		systemRouter.AuthorityRouter.InitAuthorityRouter(pubRouter, privateRouter)
		systemRouter.FTPRouter.InitApiRouter(pubRouter, privateRouter)
	}

	return Router
}
