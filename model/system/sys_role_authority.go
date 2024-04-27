package system

import "gin-one/model/common"

type SysRoleAuthority struct {
	common.CommonModel
	RoleID      string       `gorm:"role_id;type:varchar(36);"`
	AuthorityID string       `gorm:"authority_id;type:varchar(36);"`
	Role        SysRole      `gorm:"foreignKey:RoleID;"`
	Authority   SysAuthority `gorm:"foreignKey:AuthorityID;"`
}

func (SysRoleAuthority) TableName() string {
	return "sys_role_authority"
}
