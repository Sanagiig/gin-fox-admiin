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

type AuthorityApi struct{}

func (api *AuthorityApi) CreateAuthority(c *gin.Context) {
	var data request.CreateAuthorityReq
	var authorityModel system.SysAuthority

	if !ctx.MustBindWithCopy(c, &data, &authorityModel) {
		return
	}

	msgCode, err := authorityService.CreateAuthority(&authorityModel)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}

	response.OkWithDetailed(authorityModel, msgCode, c)
}

func (api *AuthorityApi) UpdateAuthority(c *gin.Context) {
	var data request.UpdateAuthorityReq
	var authorityModel system.SysAuthority

	if !ctx.MustBindWithCopy(c, &data, &authorityModel) {
		return
	}

	if err := authorityService.UpdateAuthority(&authorityModel); err != nil {
		response.FailWithMessage(message.OPER_ERR, err.Error(), c)
		return
	}

	response.OkWithDetailed(authorityModel, message.OPER_OK, c)
}

func (api *AuthorityApi) DeleteAuthorityById(c *gin.Context) {
	var data comReq.GetById

	if !ctx.MustBindWithCtx(c, &data) {
		return
	}

	msgCode, err := authorityService.DeleteAuthorityByID(data.ID)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}

	response.OkWithMessage(msgCode, c)
}

func (api *AuthorityApi) DeleteAuthorityList(c *gin.Context) {
	var data comReq.GetByIds

	if !ctx.MustBindWithCtx(c, &data) {
		return
	}

	msgCode, err := authorityService.DeleteAuthorityList(data.Ids)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}

	response.OkWithMessage(msgCode, c)
}

func (api *AuthorityApi) GetAuthorityByID(c *gin.Context) {
	var data comReq.GetById

	if !ctx.MustBindWithCtx(c, &data) {
		return
	}

	msgCode, auth, err := authorityService.GetAuthorityByID(data.ID)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}

	response.OkWithDetailed(auth, msgCode, c)
}

func (api *AuthorityApi) GetAuthority(c *gin.Context) {
	var data request.GetAuthorityReq
	var authorityModel system.SysAuthority

	if !ctx.MustBindWithCopy(c, &data, &authorityModel) {
		return
	}

}

func (api *AuthorityApi) GetAuthorityPagination(c *gin.Context) {
	var data request.GetAuthorityPaginationReq
	var authorityModel system.SysAuthority

	if !ctx.MustBindWithCopy(c, &data, &authorityModel) {
		return
	}

	msgCode, authorities, count, err := authorityService.GetAuthorityPagination(data.PageInfo, &authorityModel, data.RoleIds)
	if err != nil {
		response.FailWithDetailed(response.WrapPageData(data.PageInfo, count, authorities), msgCode, err.Error(), c)
		return
	}
	response.OkWithDetailed(response.WrapPageData(data.PageInfo, count, authorities), msgCode, c)
}

func (api *AuthorityApi) GetAuthorityChildren(c *gin.Context) {
	var data request.GetAuthorityTreeReq
	var authorityModel system.SysAuthority

	if !ctx.MustBindWithCopy(c, &data, &authorityModel) {
		response.FailWithMessage(message.DATA_STRUCT_ERR, "", c)
		return
	}

	parentIds := []string{data.ParentID}
	msgCode, auths, err := authorityService.GetAuthorityChildrenByParentIds(parentIds, &authorityModel)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}

	response.OkWithDetailed(auths, msgCode, c)
}

func (api *AuthorityApi) GetAuthorityTree(c *gin.Context) {
	var data request.GetAuthorityTreeReq
	var authorityModel system.SysAuthority

	if !ctx.MustBindWithCopy(c, &data, &authorityModel) {
		return
	}

	msgCode, authorities, err := authorityService.GetAuthorityTree(&authorityModel, nil)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}
	response.OkWithDetailed(authorities, msgCode, c)
}
