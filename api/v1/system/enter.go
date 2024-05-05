package system

import (
	"gin-one/service"
)

type ApiGroup struct {
	SysBaseApi   SysBaseApi
	UserApi      UserApi
	RoleApi      RoleApi
	AuthorityApi AuthorityApi
	AuthorizeApi AuthorizeApi
	DataDicApi   DataDicApi
	SystemApi    SystemApi
	FTPApi       FTPApi
}

var (
	baseService      = service.ServiceGroupApp.SystemServiceGroup.BaseService
	userService      = service.ServiceGroupApp.SystemServiceGroup.UserService
	roleService      = service.ServiceGroupApp.SystemServiceGroup.RoleService
	authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	AuthorizeService = service.ServiceGroupApp.SystemServiceGroup.AuthorizeService
	dataDicService   = service.ServiceGroupApp.SystemServiceGroup.DataDicService
	ftpService       = service.ServiceGroupApp.SystemServiceGroup.FTPService
	initDbService    = service.ServiceGroupApp.SystemServiceGroup.InitDBService
)
