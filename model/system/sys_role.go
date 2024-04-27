package system

import "gin-one/model/common"

type SysRole struct {
	common.UUIDModel
	RoleName    string         `json:"roleName" gorm:"role_name"`
	Description string         `json:"description"`
	Users       []SysUser      `json:"-" gorm:"many2many:sys_user_role;"`
	Authorities []SysAuthority `json:"-" gorm:"many2many:sys_role_authority"`
}

func (SysRole) TableName() string {
	return "sys_role"
}
