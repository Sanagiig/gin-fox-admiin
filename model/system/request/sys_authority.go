package request

import "gin-one/model/common/request"

type CreateAuthorityReq struct {
	AuthorityName string `json:"authorityName" binding:"required"`
	AuthorityType string `json:"authorityType" binding:"required"`
	AuthorityVal  string `json:"authorityVal" binding:"required"`
	Description   string `json:"description" binding:"required"`
}

type UpdateAuthorityReq struct {
	ID string `json:"id" binding:"required"`
	CreateAuthorityReq
}

type GetAuthorityReq struct {
	AuthorityName string   `json:"authorityName" form:"authorityName"`
	AuthorityType string   `json:"authorityType"  form:"authorityType"`
	AuthorityVal  string   `json:"authorityVal" form:"authorityVal"`
	Description   string   `json:"description" form:"description"`
	RoleIds       []string `json:"roleIds" form:"roleIds"`
}

type GetAuthorityPaginationReq struct {
	request.PageInfo
	GetAuthorityReq
}
