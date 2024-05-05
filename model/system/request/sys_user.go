package request

import (
	"gin-one/model/common"
	"gin-one/model/common/request"
)

// Register User register structure
type Register struct {
	Username     string `json:"username"  binding:"required" example:"用户名"`
	Password     string `json:"password"  binding:"required" example:"密码"`
	NickName     string `json:"nickname"  binding:"required" example:"昵称"`
	HeaderImg    string `json:"headerImg" example:"头像链接"`
	AuthorityId  uint   `json:"authorityId" swaggertype:"string" example:"int 角色id"`
	Enable       int    `json:"enable" swaggertype:"string" example:"int 是否启用"`
	AuthorityIds []uint `json:"authorityIds" swaggertype:"string" example:"[]uint 角色id"`
	Phone        string `json:"phone"  binding:"required" example:"电话号码"`
	Email        string `json:"email" example:"电子邮箱"`
	CaptchaID    string `json:"CaptchaId" binding:"required"`
	CaptchaCode  string `json:"captchaCode" binding:"required"`
}

type CreateUser struct {
	Username  string `json:"username" binding:"required" example:"用户名"`
	NickName  string `json:"nickname" binding:"required" example:"昵称"`
	HeaderImg string `json:"headerImg" example:"头像链接"`
	Password  string `json:"password" binding:"required" example:"密码"`
	Enable    int    `json:"enable" binding:"required" swaggertype:"string" example:"int 是否启用"`
	Phone     string `json:"phone"  example:"电话号码"`
	Email     string `json:"email"  example:"电子邮箱"`
}

type GetUserReq struct {
	common.StatusModel
	Username string   `json:"username" form:"username"   example:"用户名"`
	NickName string   `json:"nickname" form:"nickname"  example:"昵称"`
	Enable   int      `json:"enable" form:"enable"  swaggertype:"string" example:"int 是否启用"`
	Phone    string   `json:"phone" form:"phone"   example:"电话号码"`
	Email    string   `json:"email" form:"email"   example:"电子邮箱"`
	RoleIds  []string `json:"roleIds" form:"roleIds"`
}

type GetUserPaginationReq struct {
	request.PageInfo
	GetUserReq
}

// User login structure
type Login struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	CaptchaId   string `json:"captchaId"`   // 验证码ID
	CaptchaCode string `json:"captchaCode"` // 验证码
}

type UpdateUserReq struct {
	common.StatusModel
	ID        string `json:"id" binding:"required" example:"id"`
	HeaderImg string `json:"headerImg" example:"头像链接"`
	Username  string `json:"username" form:"username"   example:"用户名"`
	NickName  string `json:"nickname" form:"nickname"  example:"昵称"`
	Enable    int    `json:"enable" form:"enable"  swaggertype:"string" example:"int 是否启用"`
	Phone     string `json:"phone" form:"phone"   example:"电话号码"`
	Email     string `json:"email" form:"email"   example:"电子邮箱"`
}
