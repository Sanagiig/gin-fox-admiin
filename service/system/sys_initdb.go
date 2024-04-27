package system

import (
	"fmt"
	"gin-one/model/common"
	"gin-one/model/system"
)

type InitDBService struct{}

var InitDBServiceApp = new(InitDBService)

func (service *InitDBService) InitData() {
	roles := []system.SysRole{
		system.SysRole{UUIDModel: common.UUIDModel{ID: "0"}, RoleName: "管理员"},
		system.SysRole{UUIDModel: common.UUIDModel{ID: "1"}, RoleName: "普通用户"},
	}

	for _, r := range roles {
		err := RoleServiceApp.CreateRole(&r)
		if err != nil {
			panic(err)
		}
		fmt.Println("add role", r)
	}
}
