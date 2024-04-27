package system

import (
	"gin-one/model/common"
)

type SysAuthority struct {
	common.UUIDModel
	AuthorityName string    `json:"authorityName" gorm:"comment:权限名"` // 角色名
	AuthorityType string    `json:"authorityType" gorm:"comment:权限类型"`
	AuthorityVal  string    `json:"authorityVal" gorm:"comment:权限 Value"`
	Description   string    `json:"description" gorm:"comment:权限 描述"`
	Roles         []SysRole `json:"-" gorm:"many2many:sys_role_authority"`
}

func (SysAuthority) TableName() string {
	return "sys_authority"
}
