package system

import (
	"gin-one/model/common"
)

type SysAuthority struct {
	common.DateModel
	common.StatusModel
	common.TreeModel
	common.IndexModel
	AuthorityName string    `json:"authorityName" gorm:"comment:权限名"` // 角色名
	AuthorityType string    `json:"authorityType" gorm:"comment:权限类型"`
	AuthorityVal  string    `json:"authorityVal" gorm:"unique_index;comment:权限 Value"`
	Description   string    `json:"description" gorm:"type:char(100);comment:权限 描述"`
	Roles         []SysRole `json:"-" gorm:"many2many:sys_role_authority"`
}

func (SysAuthority) TableName() string {
	return "sys_authority"
}
