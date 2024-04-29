package system

import (
	"gin-one/model/common"
)

type SysUser struct {
	common.UUIDModel
	Username    string    `json:"username" gorm:"index;comment:用户登录名"`                                                  // 用户登录名
	Password    string    `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	NickName    string    `json:"nickname" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	SideMode    string    `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                          // 用户侧边主题
	HeaderImg   string    `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	BaseColor   string    `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                           // 基础颜色
	ActiveColor string    `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                      // 活跃颜色
	Roles       []SysRole `json:"roles" gorm:"many2many:sys_user_role;joinForeignKey:user_id;joinReferences:role_id"`
	Phone       string    `json:"phone"  gorm:"comment:用户手机号"`                     // 用户手机号
	Email       string    `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Description string    `json:"description"  gorm:"comment:描述"`                  // 用户邮箱
	Enable      int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	HomePath    string    `json:"homePath" gorm:"default:/dashboard/analysis"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
