package system

type RouterGroup struct {
	UserRouter
	RoleRouter
	AuthorityRouter
	AuthorizeRouter
	DataDicRouter
	BaseRouter
	FTPRouter
}
