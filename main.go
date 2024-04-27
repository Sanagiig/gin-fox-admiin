package main

import (
	"fmt"
	"gin-one/core"
	"gin-one/global"
	"gin-one/initialize"
)

func main() {
	global.Viper = core.Viper()
	global.Log = core.Zap()
	initialize.InitializeGlobal()

	fmt.Println(global.DB)
	if global.DB != nil {
		initialize.RegisterTables()
		db, _ := global.DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
