package core

import (
	"fmt"
	"gin-one/global"
	"gin-one/initialize"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	address := fmt.Sprintf(":%d", global.Config.System.Addr)
	Router := initialize.Routers()
	s := initServer(address, Router)

	global.Log.Info("server run success on ", zap.String("address", address))

	global.Log.Error(s.ListenAndServe().Error())
}
