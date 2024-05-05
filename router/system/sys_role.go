package system

import (
	v1 "gin-one/api/v1"
	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (r *RoleRouter) InitRoleRouter(pubEngine *gin.RouterGroup, priEngine *gin.RouterGroup) {
	roleApi := v1.ApiGroupApp.SystemApiGroup.RoleApi
	pub := pubEngine.Group("role")
	{
		pub.POST("/createRole", roleApi.CreateRole)
		pub.PATCH("/updateRole", roleApi.UpdateRole)
		pub.DELETE("/deleteRoleById", roleApi.DeleteRoleById)
		pub.DELETE("/deleteRoleList", roleApi.DeleteRoleList)
		pub.GET("/getRoleById", roleApi.GetRoleByID)
		pub.GET("/getRolesByUserId", roleApi.GetRolesByUserID)
		pub.GET("/getRolePagination", roleApi.GetRolePagination)
	}

	pri := pubEngine.Group("role")
	{

	}
	_ = pri
}
