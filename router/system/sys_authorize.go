package system

import (
	v1 "gin-one/api/v1"
	"github.com/gin-gonic/gin"
)

type AuthorizeRouter struct {
}

func (r *AuthorityRouter) InitAuthorizeRouter(pubEngine *gin.RouterGroup, priEngine *gin.RouterGroup) {
	AuthorizeApi := v1.ApiGroupApp.SystemApiGroup.AuthorizeApi
	pub := pubEngine.Group("authority")
	{

	}

	pri := pubEngine.Group("authority")
	{
		pri.PATCH("/updateUserRoles", AuthorizeApi.UpdateUserRoles)
		pri.PATCH("/updateRoleAuthorities", AuthorizeApi.UpdateRoleAuthorities)
	}
	_ = pub
	_ = pri
}
