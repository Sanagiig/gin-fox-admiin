package request

import "gin-one/model/common/request"

// Register User register structure
type Register struct {
	Username     string `json:"username" example:"用户名"`
	Password     string `json:"password" example:"密码"`
	NickName     string `json:"nickName" example:"昵称"`
	HeaderImg    string `json:"headerImg" example:"头像链接"`
	AuthorityId  uint   `json:"authorityId" swaggertype:"string" example:"int 角色id"`
	Enable       int    `json:"enable" swaggertype:"string" example:"int 是否启用"`
	AuthorityIds []uint `json:"authorityIds" swaggertype:"string" example:"[]uint 角色id"`
	Phone        string `json:"phone" example:"电话号码"`
	Email        string `json:"email" example:"电子邮箱"`
}

type CreateUser struct {
	Username  string `json:"username" binding:"required" example:"用户名"`
	NickName  string `json:"nickName" binding:"required" example:"昵称"`
	HeaderImg string `json:"headerImg" example:"头像链接"`
	Password  string `json:"password" binding:"required" example:"密码"`
	Enable    int    `json:"enable" binding:"required" swaggertype:"string" example:"int 是否启用"`
	Phone     string `json:"phone"  example:"电话号码"`
	Email     string `json:"email"  example:"电子邮箱"`
}

type GetUserReq struct {
	Username string   `json:"username" form:"username"   example:"用户名"`
	NickName string   `json:"nickName" form:"nickName"  example:"昵称"`
	Enable   int      `json:"enable" form:"enable"  swaggertype:"string" example:"int 是否启用"`
	Phone    string   `json:"phone" form:"phone"   example:"电话号码"`
	Email    string   `json:"email" form:"email"   example:"电子邮箱"`
	RoleIds  []string `json:"roleIds" form:"roleIds"`
}

type GetUserPaginationReq struct {
	request.PageInfo
	GetUserReq
}

type UpdateUserReq struct {
	ID string `json:"id" example:"头像链接"`
	GetUserReq
}

// User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}
