package system

import "gin-one/model/common"

type SysDataDic struct {
	common.UuidDateModel
	common.StatusModel
	common.ParentModel
	common.IndexModel
	Name        string `json:"name" gorm:"name;type:char(100);"`
	Code        string `json:"code" gorm:"type:char(20);code;comment:编码"`
	Description string `json:"description" gorm:"description;type:char(200);"`
}

func (SysDataDic) TableName() string {
	return "sys_data_dic"
}
