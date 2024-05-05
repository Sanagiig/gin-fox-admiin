package system

import (
	v1 "gin-one/api/v1"
	"github.com/gin-gonic/gin"
)

type DataDicRouter struct{}

func (r *DataDicRouter) InitDataDicRouter(pubEngine *gin.RouterGroup, priEngine *gin.RouterGroup) {
	roleApi := v1.ApiGroupApp.SystemApiGroup.DataDicApi
	pub := pubEngine.Group("dataDic")
	{
		pub.POST("/createDataDic", roleApi.CreateDataDic)
		pub.PATCH("/updateDataDic", roleApi.UpdateDataDic)
		pub.DELETE("/deleteDataDicById", roleApi.DeleteDataDicById)
		pub.DELETE("/deleteDataDicList", roleApi.DeleteDataDicList)
		pub.GET("/getDataDicById", roleApi.GetDataDicByID)
		pub.GET("/getDataDicListByParentIds", roleApi.GetDataDicListByParentIDs)
		pub.GET("/getDataDicPagination", roleApi.GetDataDicPagination)
		pub.GET("/getDataDicListByCode", roleApi.GetDataDicListByCode)
	}

	pri := pubEngine.Group("role")
	{

	}
	_ = pri
}
