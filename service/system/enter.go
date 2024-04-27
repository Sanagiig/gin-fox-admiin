package system

type ServiceGroup struct {
	*UserService
	*RoleService
	*AuthorityService
	*BaseService
	*FTPService
	*CasbinService
	*InitDBService
}

var ServiceGroupApp = ServiceGroup{
	UserService:      UserServiceApp,
	RoleService:      RoleServiceApp,
	AuthorityService: AuthorityServiceApp,
	BaseService:      BaseServiceApp,
	FTPService:       FTPServiceApp,
	CasbinService:    CasbinServiceApp,
	InitDBService:    InitDBServiceApp,
}
