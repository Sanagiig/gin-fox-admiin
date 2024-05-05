package system

import (
	v1 "gin-one/api/v1"
	"github.com/gin-gonic/gin"
)

type AuthorityRouter struct {
}

func (r *AuthorityRouter) InitAuthorityRouter(pubEngine *gin.RouterGroup, priEngine *gin.RouterGroup) {
	authorityApi := v1.ApiGroupApp.SystemApiGroup.AuthorityApi
	pub := pubEngine.Group("authority")
	{
		pub.POST("/createAuthority", authorityApi.CreateAuthority)
		pub.PATCH("/updateAuthority", authorityApi.UpdateAuthority)
		pub.DELETE("/deleteAuthorityById", authorityApi.DeleteAuthorityById)
		pub.DELETE("/deleteAuthorityList", authorityApi.DeleteAuthorityList)
		pub.GET("/getAuthorityById", authorityApi.GetAuthorityByID)
		pub.GET("/getAuthorityPagination", authorityApi.GetAuthorityPagination)
		pub.GET("/getAuthorityTree", authorityApi.GetAuthorityTree)
		pub.GET("/getAuthorityChildren", authorityApi.GetAuthorityChildren)
	}

	pri := pubEngine.Group("authority")
	{

	}
	_ = pri
}
