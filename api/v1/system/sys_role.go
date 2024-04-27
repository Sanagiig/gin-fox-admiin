package system

import (
	"gin-one/message"
	comReq "gin-one/model/common/request"
	"gin-one/model/common/response"
	"gin-one/model/system"
	"gin-one/model/system/request"
	"gin-one/utils/ctx"
	"github.com/gin-gonic/gin"
)

type RoleApi struct{}

func (api *RoleApi) CreateRole(c *gin.Context) {
	var data request.CreateRoleReq
	var roleModel system.SysRole
	var err error

	if !ctx.MustBindWithCopy(c, &data, &roleModel) {
		return
	}

	err = roleService.CreateRole(&roleModel)
	if err != nil {
		response.FailWithMessage(message.OPER_DB_ERR, err.Error(), c)
		return
	}

	response.OkWithDetailed(roleModel, message.OPER_OK, c)
}

func (api *RoleApi) UpdateRole(c *gin.Context) {
	var data request.UpdateRoleReq
	var roleModel system.SysRole
	var err error

	if !ctx.MustBindWithCopy(c, &data, &roleModel) {
		return
	}

	err = roleService.UpdateRole(&roleModel)
	if err != nil {
		response.FailWithMessage(message.OPER_ERR, err.Error(), c)
		return
	}

	response.OkWithDetailed(roleModel, message.OPER_OK, c)
}

func (api *RoleApi) DeleteRoleById(c *gin.Context) {
	var data comReq.GetById

	if !ctx.MustBindWithCtx(c, &data) {
		return
	}

	msgCode, err := roleService.DeleteRoleByID(data.ID)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}

	response.OkWithMessage(msgCode, c)
}

func (api *RoleApi) DeleteRoleList(c *gin.Context) {
	var data comReq.GetByIds

	if !ctx.MustBindWithCtx(c, &data) {
		return
	}

	msgCode, err := roleService.DeleteRoleList(data.Ids)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}

	response.OkWithMessage(msgCode, c)
}

func (api *RoleApi) GetRoleByID(c *gin.Context) {
	var data comReq.GetById

	if !ctx.MustBindWithCtx(c, &data) {
		return
	}

	msgCode, role, err := roleService.GetRoleByID(data.ID)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}

	response.OkWithDetailed(role, msgCode, c)
}

func (api *RoleApi) GetRolePagination(c *gin.Context) {
	var data request.GetRolePaginationReq
	var userModel system.SysRole

	if !ctx.MustBindWithCopy(c, &data, &userModel) {
		return
	}

	msgCode, roles, count, err := roleService.GetRolePagination(data.PageInfo, &userModel, data.AuthorityIds)
	if err != nil {
		response.FailWithDetailed(response.WrapPageData(data.PageInfo, count, roles), msgCode, err.Error(), c)
		return
	}

	response.OkWithDetailed(response.WrapPageData(data.PageInfo, count, roles), msgCode, c)
}
