package system

import (
	"fmt"
	"gin-one/model/common"
	"gin-one/model/system"
)

type InitDBService struct{}

var InitDBServiceApp = new(InitDBService)

func (service *InitDBService) InitData() {
	roles := []system.SysRole{system.SysRole{
		UuidDateModel: common.UuidDateModel{
			UUIDModel: common.UUIDModel{ID: "0"},
		},
		Name: "管理员",
		Code: "admin",
	}, system.SysRole{
		UuidDateModel: common.UuidDateModel{
			UUIDModel: common.UUIDModel{ID: "1"},
		},
		StatusModel: common.StatusModel{},
		Name:        "普通用户",
		Code:        "user",
	}}

	for _, r := range roles {
		err := RoleServiceApp.CreateRole(&r)
		if err != nil {
			panic(err)
		}
		fmt.Println("add role", r)
	}
}
