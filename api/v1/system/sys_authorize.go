package system

import (
	"gin-one/model/common/response"
	"gin-one/model/system/request"
	"gin-one/utils/ctx"
	"github.com/gin-gonic/gin"
)

type AuthorizeApi struct{}

func (api *AuthorizeApi) UpdateUserRoles(c *gin.Context) {
	var data request.UpdateUserRolesReq

	if !ctx.MustBindWithCtx(c, &data) {
		return
	}

	msgCode, err := AuthorizeService.UpdateUserRolesByIds(data)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}

	response.OkWithMessage(msgCode, c)
}

func (api *AuthorizeApi) UpdateRoleAuthorities(c *gin.Context) {
	var data request.UpdateRoleAuthoritiesReq

	if !ctx.MustBindWithCtx(c, &data) {
		return
	}

	msgCode, err := AuthorizeService.UpdateRoleAuthoritiesByIds(data)
	if err != nil {
		response.FailWithMessage(msgCode, err.Error(), c)
		return
	}

	response.OkWithMessage(msgCode, c)
}
