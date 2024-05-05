package system

type ServiceGroup struct {
	*BaseService
	*UserService
	*RoleService
	*AuthorityService
	*AuthorizeService
	*DataDicService
	*FTPService
	*CasbinService
	*InitDBService
}

var ServiceGroupApp = ServiceGroup{
	UserService:      UserServiceApp,
	RoleService:      RoleServiceApp,
	AuthorityService: AuthorityServiceApp,
	AuthorizeService: AuthorizeServiceApp,
	DataDicService:   DataDicServiceApp,
	BaseService:      BaseServiceApp,
	FTPService:       FTPServiceApp,
	CasbinService:    CasbinServiceApp,
	InitDBService:    InitDBServiceApp,
}
