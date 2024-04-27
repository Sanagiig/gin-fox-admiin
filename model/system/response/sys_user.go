package response

import "gin-one/model/system"

type LoginResponse struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

type UserResponse struct {
	User system.SysUser `json:"user"`
}
