package system

import (
	"fmt"
	"gin-one/model/common"
	"gorm.io/gorm"
)

type SysUserRole struct {
	common.UUIDModel
	UserID  string  `gorm:"column:user_id;type:varchar(36);"`
	RoleID  string  `gorm:"column:role_id;type:varchar(36);"`
	SysUser SysUser `gorm:"foreignKey:user_id;"`
	SysRole SysRole `gorm:"foreignKey:role_id;"`
}

func (s *SysUserRole) TableName() string {
	return "sys_user_role"
}

func (s *SysUserRole) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("sys before")
	return nil
}
