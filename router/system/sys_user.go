package system

import (
	v1 "gin-one/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (router *UserRouter) InitUserRouter(pubEngine *gin.RouterGroup, priEngine *gin.RouterGroup) {
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	pub := pubEngine.Group("user")
	{
		pub.POST("/register", userApi.Register)
		pub.POST("/createUser", userApi.CreateUser)
		pub.PATCH("/updateUser", userApi.UpdateUser)
		pub.DELETE("/deleteUserById", userApi.DeleteUserById)
		pub.DELETE("/deleteUserList", userApi.DeleteUserList)
		pub.GET("/getUser", userApi.GetUser)
		pub.GET("/getUserById", userApi.GetUserByID)
		pub.GET("/getUserPagination", userApi.GetUserPagination)
		pub.POST("/login", userApi.Login)
	}
	pri := priEngine.Group("user")
	{
		pri.POST("createUser1", userApi.CreateUser)
	}
}
