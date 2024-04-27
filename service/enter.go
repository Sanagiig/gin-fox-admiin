package service

import "gin-one/service/system"

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupApp = ServiceGroup{
	SystemServiceGroup: system.ServiceGroupApp,
}
