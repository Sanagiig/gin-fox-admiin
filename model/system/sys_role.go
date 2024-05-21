package system

import "gin-one/model/common"

type SysRole struct {
	common.UuidDateModel
	common.StatusModel
	Name        string         `json:"name" gorm:"name;type:char(50);"`
	Code        string         `json:"code" gorm:"code;unique_index;char(50);"`
	Description string         `json:"description" gorm:"type:char(200);"`
	Users       []SysUser      `json:"-" gorm:"many2many:sys_user_role;"`
	Authorities []SysAuthority `json:"-" gorm:"many2many:sys_role_authority"`
}

func (SysRole) TableName() string {
	return "sys_role"
}
