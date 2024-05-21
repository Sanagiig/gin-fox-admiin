package initialize

import (
	"gin-one/middlewares"
	"gin-one/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	Router.Use(middlewares.Cors())

	pubRouter := Router.Group("")
	privateRouter := Router.Group("")

	privateRouter.Use(middlewares.JWTAuth())
	systemRouter := router.RouterGroupApp.System
	{
		systemRouter.BaseRouter.InitBaseRouter(pubRouter)
		systemRouter.UserRouter.InitUserRouter(pubRouter, privateRouter)
		systemRouter.RoleRouter.InitRoleRouter(pubRouter, privateRouter)
		systemRouter.AuthorityRouter.InitAuthorityRouter(pubRouter, privateRouter)
		systemRouter.AuthorizeRouter.InitAuthorizeRouter(pubRouter, privateRouter)
		systemRouter.DataDicRouter.InitDataDicRouter(pubRouter, privateRouter)
		systemRouter.FTPRouter.InitApiRouter(pubRouter, privateRouter)
	}

	return Router
}
