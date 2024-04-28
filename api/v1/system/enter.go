package system

import (
	"gin-one/service"
)

type ApiGroup struct {
	UserApi      UserApi
	RoleApi      RoleApi
	AuthorityApi AuthorityApi
	SysBaseApi   SysBaseApi
	SystemApi    SystemApi
	FTPApi       FTPApi
}

var (
	userService      = service.ServiceGroupApp.SystemServiceGroup.UserService
	roleService      = service.ServiceGroupApp.SystemServiceGroup.RoleService
	authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	ftpService       = service.ServiceGroupApp.SystemServiceGroup.FTPService
	initDbService    = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	baseService      = service.ServiceGroupApp.SystemServiceGroup.BaseService
)
