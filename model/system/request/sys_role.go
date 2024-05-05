package request

import (
	"gin-one/model/common"
	"gin-one/model/common/request"
)

type CreateRoleReq struct {
	common.StatusModel
	RoleName    string   `json:"roleName" binding:"required"`
	RoleCode    string   `json:"roleCode"  binding:"required"`
	Description string   `json:"description" binding:"required"`
	Authorities []string `json:"authorities"`
}

type UpdateRoleReq struct {
	ID string `json:"id" binding:"required"`
	CreateRoleReq
}

type AuthorityInfo struct {
	common.UUIDModel
	AuthorityType string `json:"authorityType" binding:"required"`
}

type GetRoleReq struct {
	RoleName     string   `json:"roleName" form:"authorityName"`
	Description  string   `json:"description" form:"description"`
	AuthorityIds []string `json:"authorityIds" form:"authorityIds"`
}

type GetRolePaginationReq struct {
	request.PageInfo
	GetRoleReq
}
