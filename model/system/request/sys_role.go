package request

import "gin-one/model/common/request"

type CreateRoleReq struct {
	RoleName    string   `json:"authorityName" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Authorities []string `json:"authorities"`
}

type UpdateRoleReq struct {
	ID string `json:"id" binding:"required"`
	CreateRoleReq
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
