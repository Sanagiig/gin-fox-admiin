package request

import "gin-one/model/common/request"

type CreateAuthorityReq struct {
	ParentID      string `json:"parentId"`
	AuthorityName string `json:"authorityName" binding:"required"`
	AuthorityType string `json:"authorityType" binding:"required"`
	AuthorityVal  string `json:"authorityVal" binding:"required"`
	Status        string `json:"status"`
	Description   string `json:"description"`
}

type UpdateAuthorityReq struct {
	ID string `json:"id" binding:"required"`
	CreateAuthorityReq
}

type GetAuthorityTreeReq struct {
	ParentID      string `json:"parentId" form:"parentId"`
	AuthorityName string `json:"authorityName" form:"authorityName"`
	AuthorityType string `json:"authorityType"  form:"authorityType"`
	AuthorityVal  string `json:"authorityVal" form:"authorityVal"`
	Status        string `json:"status" form:"status"`
}

type GetAuthorityReq struct {
	AuthorityName string   `json:"authorityName" form:"authorityName"`
	AuthorityType string   `json:"authorityType"  form:"authorityType"`
	AuthorityVal  string   `json:"authorityVal" form:"authorityVal"`
	Description   string   `json:"description" form:"description"`
	Status        string   `json:"status" form:"status"`
	RoleIds       []string `json:"roleIds" form:"roleIds"`
}

type GetAuthorityPaginationReq struct {
	request.PageInfo
	GetAuthorityReq
}
