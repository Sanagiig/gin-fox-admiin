package request

import (
	"gin-one/model/common"
	"gin-one/model/common/request"
)

type CreateDataDicReq struct {
	common.StatusModel
	ParentID    string `json:"parentID"`
	Name        string `json:"name" binding:"required"`
	Code        string `json:"code"  binding:"required"`
	Description string `json:"description"`
}

type UpdateDataDicReq struct {
	ID string `json:"id" binding:"required"`
	CreateDataDicReq
}

type GetDataDicReq struct {
	UpdateDataDicReq
}

type GetDataDicPaginationReq struct {
	request.PageInfo
	common.StatusModel
	ParentID    string   `json:"parentId" form:"parentId"`
	ParentIDs   []string `json:"parentIds" form:"parentIds"`
	Name        string   `json:"name" form:"name"`
	Code        string   `json:"code" form:"code"`
	Description string   `json:"description" form:"description"`
}
