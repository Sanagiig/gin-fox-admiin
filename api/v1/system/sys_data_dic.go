package system

import (
	"gin-one/message"
	comReq "gin-one/model/common/request"
	"gin-one/model/common/response"
	"gin-one/model/system"
	"gin-one/model/system/request"
	"gin-one/service/common"
	"gin-one/utils/ctx"
	"github.com/gin-gonic/gin"
)

type DataDicApi struct{}

func (api *DataDicApi) CreateDataDic(c *gin.Context) {
	var data request.CreateDataDicReq
	var dataDicModel system.SysDataDic
	var err error

	if !ctx.MustBindWithCopy(c, &data, &dataDicModel) {
		return
	}

	msgCode, err := dataDicService.CreateDataDic(&dataDicModel)
	if err != nil {
		response.FailWithMessage(message.OPER_DB_ERR, err.Error(), c)
		return
	} else if msgCode != message.OPER_OK {
		response.FailWithMessage(msgCode, "", c)
		return
	}

	response.OkWithDetailed(dataDicModel, message.OPER_OK, c)
}

func (api *DataDicApi) UpdateDataDic(c *gin.Context) {
	var data request.UpdateDataDicReq
	var dataDicModel system.SysDataDic
	var err error

	if !ctx.MustBindWithCopy(c, &data, &dataDicModel) {
		return
	}

	err = dataDicService.UpdateDataDic(&dataDicModel)
	if err != nil {
		response.FailWithMessage(message.OPER_ERR, err.Error(), c)
		return
	}

	response.OkWithDetailed(dataDicModel, message.OPER_OK, c)
}

func (api *DataDicApi) DeleteDataDicById(c *gin.Context) {
	var data comReq.GetById

	if !ctx.MustBindWithCtx(c, &data) {
		return
	}

	msgCode, err := dataDicService.DeleteDataDicByID(data.ID)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}

	response.OkWithMessage(msgCode, c)
}

func (api *DataDicApi) DeleteDataDicList(c *gin.Context) {
	var data comReq.GetByIds

	if !ctx.MustBindWithCtx(c, &data) {
		return
	}

	msgCode, err := dataDicService.DeleteDataDicList(data.Ids)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}

	response.OkWithMessage(msgCode, c)
}

func (api *DataDicApi) GetDataDicByID(c *gin.Context) {
	var data comReq.GetById

	if !ctx.MustBindWithCtx(c, &data) {
		return
	}

	msgCode, dataDic, err := dataDicService.GetDataDicByID(data.ID)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}

	msgCode, err = common.GetParentName(&system.SysDataDic{}, &dataDic)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}

	response.OkWithDetailed(dataDic, msgCode, c)
}

func (api *DataDicApi) GetDataDicListByParentIDs(c *gin.Context) {
	var data comReq.GetByParentIds

	if !ctx.MustBindWithCtx(c, &data) {
		return
	}

	msgCode, dataDics, err := dataDicService.GetDataDicByParentIds(data.ParentIds)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}

	//msgCode, err = common.GetParentName(dataDics)
	//if err != nil {
	//	response.FailWithMessage(msgCode, err.Error(), c)
	//	return
	//}

	response.OkWithDetailed(dataDics, msgCode, c)
}

func (api *DataDicApi) GetDataDicPagination(c *gin.Context) {
	var data request.GetDataDicPaginationReq
	var dicModel system.SysDataDic

	if !ctx.MustBindWithCopy(c, &data, &dicModel) {
		return
	}

	data.ParentIDs = append(data.ParentIDs, data.ParentID)

	msgCode, dataDics, count, err := dataDicService.GetDataDicPagination(data.PageInfo, &dicModel, data.ParentIDs)
	if err != nil {
		response.FailWithPage(data.PageInfo, msgCode, err.Error(), c)
		return
	}

	if len(dataDics) > 0 {

	}

	msgCode, err = common.GetParentName(&system.SysDataDic{}, &dataDics)
	if err != nil {
		response.FailWithPage(data.PageInfo, msgCode, err.Error(), c)
		return
	}

	response.OkWithDetailed(response.WrapPageData(data.PageInfo, count, dataDics), msgCode, c)
}

func (api *DataDicApi) GetDataDicListByCode(c *gin.Context) {
	var data comReq.GetByCode

	if !ctx.MustBindWithCtx(c, &data) {
		return
	}

	msgCode, dataDics, err := dataDicService.GetDataDicListByCode(data.Code)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
	}
	response.OkWithDetailed(dataDics, msgCode, c)
}
