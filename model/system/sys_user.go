package system

import (
	"gin-one/model/common"
)

type SysUser struct {
	common.UuidDateModel
	common.StatusModel
	Username    string    `json:"username" gorm:"index;comment:用户登录名;type:char(50);"`                                                  // 用户登录名
	Password    string    `json:"-"  gorm:"comment:用户登录密码;type:char(100);"`                                                            // 用户登录密码
	NickName    string    `json:"nickname" gorm:"index;type:char(50);default:系统用户;comment:用户昵称"`                                       // 用户昵称
	SideMode    string    `json:"sideMode" gorm:"type:char(50);default:dark;comment:用户侧边主题"`                                           // 用户侧边主题
	HeaderImg   string    `json:"headerImg" gorm:"type:char(300);default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	BaseColor   string    `json:"baseColor" gorm:"type:char(50);default:#fff;comment:基础颜色"`                                            // 基础颜色
	ActiveColor string    `json:"activeColor" gorm:"type:char(50);default:#1890ff;comment:活跃颜色"`                                       // 活跃颜色
	Roles       []SysRole `json:"roles" gorm:"many2many:sys_user_role;joinForeignKey:user_id;joinReferences:role_id"`
	Phone       string    `json:"phone"  gorm:"comment:用户手机号"`                   // 用户手机号
	Email       string    `json:"email"  gorm:"type:char(150);comment:用户邮箱"`     // 用户邮箱
	Description string    `json:"description"  gorm:"type:char(200);comment:描述"` // 用户邮箱
	HomePath    string    `json:"homePath" gorm:"type:char(150);default:/dashboard/analysis"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
