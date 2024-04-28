package middlewares

import (
	"gin-one/service/system"
	"github.com/gin-gonic/gin"
)

var casbinService = system.CasbinServiceApp

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//claims, err := ctx.GetClaims(c)
		//if err != nil {
		//	zap.L().Error("Claims 解释错误", zap.Error(err))
		//}
		//
		//path := c.Request.URL.Path
		//obj := strings.Trim(path, global.Config.System.RouterPrefix)
		//act := c.Request.Method

	}
}
